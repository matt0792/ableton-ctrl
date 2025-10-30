package als_test

import (
	"fmt"
	"log"

	"github.com/hypebeast/go-osc/osc"
	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/oscclient"
)

// Example demonstrates basic usage of the Ableton client
func Example_basicUsage() {
	// Create a new client
	// By default, AbletonOSC listens on port 11000 and sends replies on port 11001
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000, // Port to send messages to (AbletonOSC default)
		ListenAddr: 11001, // Port to listen for responses on (AbletonOSC default)
	})

	// Start the client to begin listening for responses
	client.Run()
	defer client.Close()

	// Test the connection
	result := client.Application.Test()
	fmt.Println("Connection test:", result)

	// Get Ableton Live version
	major, minor := client.Application.GetVersion()
	fmt.Printf("Ableton Live version: %d.%d\n", major, minor)

	// Get and set tempo
	tempo := client.Song.GetTempo()
	fmt.Printf("Current tempo: %.2f BPM\n", tempo)

	// Set tempo to 120 BPM
	client.Song.SetTempo(120.0)

	// Start playback
	client.Song.StartPlaying()

	// Get track information
	numTracks := client.Song.GetNumTracks()
	fmt.Printf("Number of tracks: %d\n", numTracks)

	// Get track names
	trackNames := client.Song.GetTrackNames()
	for i, name := range trackNames {
		fmt.Printf("Track %d: %s\n", i, name)
	}

	// Get track volume (track 0)
	volume := client.Track.GetVolume(0)
	fmt.Printf("Track 0 volume: %.2f\n", volume)

	// Set track volume
	client.Track.SetVolume(0, 0.85)

	// Mute a track
	client.Track.SetMute(1, true)

	// Fire a clip
	client.Clip.Fire(0, 0) // Track 0, Clip 0

	// Get clip information
	clipName := client.Clip.GetName(0, 0)
	clipLength := client.Clip.GetLength(0, 0)
	isPlaying := client.Clip.GetIsPlaying(0, 0)
	fmt.Printf("Clip: %s, Length: %.2f, Playing: %v\n", clipName, clipLength, isPlaying)

	// Fire a scene
	client.Scene.Fire(0)

	// Get selected track
	selectedTrack := client.View.GetSelectedTrack()
	fmt.Printf("Selected track: %d\n", selectedTrack)

	// Change selected track
	client.View.SetSelectedTrack(2)

	// Stop playback
	client.Song.StopPlaying()
}

// Example_workingWithClips demonstrates working with MIDI clips
func Example_workingWithClips() {
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
	})
	client.Run()
	defer client.Close()

	trackID := int32(0)
	clipID := int32(0)

	// Get notes from a clip
	notes := client.Clip.GetNotes(trackID, clipID)
	fmt.Printf("Found %d notes in clip\n", len(notes))

	// Add notes to a clip
	newNotes := []als.Note{
		{Pitch: 60, StartTime: 0.0, Duration: 1.0, Velocity: 100, Mute: false}, // C4
		{Pitch: 64, StartTime: 1.0, Duration: 1.0, Velocity: 100, Mute: false}, // E4
		{Pitch: 67, StartTime: 2.0, Duration: 1.0, Velocity: 100, Mute: false}, // G4
		{Pitch: 72, StartTime: 3.0, Duration: 1.0, Velocity: 100, Mute: false}, // C5
	}
	client.Clip.AddNotes(trackID, clipID, newNotes...)

	// Remove notes in a specific range
	client.Clip.RemoveNotes(trackID, clipID, 60, 12, 0.0, 4.0) // Remove notes C4-B4, from time 0-4

	// Set loop points
	client.Clip.SetLoopStart(trackID, clipID, 0.0)
	client.Clip.SetLoopEnd(trackID, clipID, 4.0)
}

// Example_deviceControl demonstrates controlling devices
func Example_deviceControl() {
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
	})
	client.Run()
	defer client.Close()

	trackID := int32(0)
	deviceID := int32(0)

	// Get device information
	deviceName := client.Device.GetName(trackID, deviceID)
	deviceType := client.Device.GetType(trackID, deviceID)
	fmt.Printf("Device: %s (%s)\n", deviceName, deviceType)

	// Get parameter names
	paramNames := client.Device.GetParametersName(trackID, deviceID)
	for i, name := range paramNames {
		fmt.Printf("Parameter %d: %s\n", i, name)
	}

	// Get parameter values
	paramValues := client.Device.GetParametersValue(trackID, deviceID)
	for i, value := range paramValues {
		fmt.Printf("Parameter %d value: %.2f\n", i, value)
	}

	// Set a specific parameter value
	client.Device.SetParameterValue(trackID, deviceID, 0, 0.75)
}

// Example_listening demonstrates listening for property changes
func Example_listening() {
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []oscclient.DispatcherOption{
			oscclient.WithHandler("/live/song/get/tempo", func(msg *osc.Message) {
				if len(msg.Arguments) > 0 {
					if tempo, ok := msg.Arguments[0].(float32); ok {
						log.Printf("Tempo changed to: %.2f BPM\n", tempo)
					}
				}
			}),
		},
	})
	client.Run()
	defer client.Close()

	// Start listening for tempo changes
	client.Song.StartListenProperty("tempo")

	// Do other work...

	// Stop listening for tempo changes
	client.Song.StopListenProperty("tempo")

	// Listen for beat changes
	client.Song.StartListenBeat()
	// ... work ...
	client.Song.StopListenBeat()

	// Listen for track property changes
	client.Track.StartListenProperty(0, "volume")
	// ... work ...
	client.Track.StopListenProperty(0, "volume")
}
