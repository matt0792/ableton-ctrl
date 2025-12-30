package track

import (
	"context"
	"math"
	"time"

	"github.com/matt0792/ableton-ctrl/als"
)

type Track struct {
	api           *als.TrackAPI
	trackID       int32
	autoVolume    *AutoVolume
	volumeMonitor *VolumeMonitor
}

func New(api *als.TrackAPI, trackID int32) *Track {
	t := &Track{
		api:     api,
		trackID: trackID,
	}
	t.autoVolume = &AutoVolume{Track: t}
	t.volumeMonitor = NewVolumeMonitor(t)
	return t
}

type Volume struct {
	*Track
}

func (t *Track) Volume() *Volume {
	return &Volume{Track: t}
}

func (v *Volume) Auto(isEnabled bool) {
	if isEnabled {
		v.autoVolume.Start()
	} else {
		v.autoVolume.Stop()
	}
}

func (v *Volume) Set(value float32) {
	v.autoVolume.Stop()
	v.api.SetVolume(v.trackID, value)
}

func (v *Volume) Get() float32 {
	return v.api.GetVolume(v.trackID)
}

// AutoVolume handles auto volume adjustment, similar to a compressor
type AutoVolume struct {
	*Track
	ctx         context.Context
	cancel      context.CancelFunc
	targetRMS   float32
	maxAdjust   float32
	attackTime  float32
	releaseTime float32
}

func (av *AutoVolume) Start() {
	if av.cancel != nil {
		return
	}

	av.ctx, av.cancel = context.WithCancel(context.Background())

	av.targetRMS = -15.0 // dB RMS
	av.maxAdjust = 0.5
	av.attackTime = 2.0 // seconds
	av.releaseTime = 4.0

	go av.monitor()
}

func (av *AutoVolume) Stop() {
	if av.cancel != nil {
		av.cancel()
		av.cancel = nil
	}
}

func (av *AutoVolume) monitor() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-av.ctx.Done():
			return
		case <-ticker.C:
			av.adjust()
		}
	}
}

func (av *AutoVolume) adjust() {
	level := av.volumeMonitor.GetCurrentLevel()

	if level < -60.0 {
		return
	}

	diff := av.targetRMS - level

	if math.Abs(float64(diff)) < 3.0 {
		return
	}

	adjustSpeed := av.attackTime
	if diff > 0 {
		adjustSpeed = av.releaseTime
	}

	adjustment := diff / adjustSpeed

	// clamp
	if adjustment > av.maxAdjust {
		adjustment = av.maxAdjust
	} else if adjustment < -av.maxAdjust {
		adjustment = -av.maxAdjust
	}

	// get current & adjust
	currentVol := av.api.GetVolume(av.trackID)
	newVol := dbToLinear(linearToDb(currentVol) + adjustment)

	// clamp
	if newVol < 0.01 { // about -40dB
		newVol = 0.01
	} else if newVol > 1.0 { // 0dB
		newVol = 1.0
	}

	av.api.SetVolume(av.trackID, newVol)
}

// VolumeMonitor tracks audio levels
type VolumeMonitor struct {
	track        *Track
	currentLevel float32
	rmsWindow    []float32
	windowSize   int
}

func NewVolumeMonitor(track *Track) *VolumeMonitor {
	return &VolumeMonitor{
		track:      track,
		rmsWindow:  make([]float32, 0, 20),
		windowSize: 20, // 2 second avg
	}
}

func (vm *VolumeMonitor) GetCurrentLevel() float32 {
	// get peak from track
	peak := vm.track.api.GetOutputMeterLevel(vm.track.trackID)

	// add to rms window
	vm.rmsWindow = append(vm.rmsWindow, peak)
	if len(vm.rmsWindow) > vm.windowSize {
		vm.rmsWindow = vm.rmsWindow[1:]
	}

	// calc RMS
	var sum float32
	for _, v := range vm.rmsWindow {
		sum += v * v
	}
	rms := float32(math.Sqrt(float64(sum / float32(len(vm.rmsWindow)))))

	// to dB
	vm.currentLevel = linearToDb(rms)
	return vm.currentLevel
}

// --- Helpers ---
func linearToDb(linear float32) float32 {
	if linear <= 0 {
		return -100.0
	}
	return float32(20.0 * math.Log10(float64(linear)))
}

func dbToLinear(db float32) float32 {
	return float32(math.Pow(10.0, float64(db)/20.0))
}
