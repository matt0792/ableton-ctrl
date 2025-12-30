package clip

import (
	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/alsex/note"
)

type Clip struct {
	api     *als.ClipAPI
	trackID int32
	clipID  int32
	gain    *Gain
	name    *Name
	notes   *Notes
}

func New(client *als.Client, trackId, clipId int32) *Clip {
	c := &Clip{
		api:     client.Clip,
		trackID: trackId,
		clipID:  clipId,
	}
	c.gain = &Gain{c}
	c.name = &Name{c}
	c.notes = &Notes{c}
	return c
}

func (c *Clip) Fire() {
	c.api.Fire(c.trackID, c.clipID)
}

func (c *Clip) Stop() {
	c.api.Stop(c.trackID, c.clipID)
}

// Gain

type Gain struct {
	*Clip
}

func (c *Clip) Gain() *Gain {
	return c.gain
}

func (g *Gain) Get() float32 {
	return g.api.GetGain(g.trackID, g.clipID)
}

func (g *Gain) Set(value float32) {
	g.api.SetGain(g.trackID, g.clipID, value)
}

// Name

type Name struct {
	*Clip
}

func (c *Clip) Name() *Name {
	return c.name
}

func (n *Name) Get() string {
	return n.api.GetName(n.trackID, n.clipID)
}

func (n *Name) Set(value string) {
	n.api.SetName(n.trackID, n.clipID, value)
}

func (c *Clip) Length() float32 {
	return c.api.GetLength(c.trackID, c.clipID)
}

// Notes

type Notes struct {
	*Clip
}

func (c *Clip) Notes() *Notes {
	return c.notes
}

func (n *Notes) Get() []als.Note {
	return n.api.GetNotes(n.trackID, n.clipID)
}

type NoteBuilder struct {
	clip     *Clip
	notes    []als.Note
	duration float32
	velocity int32
}

// Names creates notes from note names
func (n *Notes) Names(notes ...string) *NoteBuilder {
	nb := &NoteBuilder{
		clip:     n.Clip,
		notes:    make([]als.Note, 0, len(notes)),
		duration: 0.25,
	}

	for i, noteName := range notes {
		midiPitch, err := note.ToMidi(noteName)
		if err != nil {
			continue
		}
		nb.notes = append(nb.notes, als.Note{
			Pitch:     midiPitch,
			StartTime: float32(i) * nb.duration,
			Duration:  nb.duration,
			Velocity:  nb.velocity,
			Mute:      false,
		})
	}

	return nb
}

// Duration sets the duration for all notes
func (nb *NoteBuilder) Duration(beats float32) *NoteBuilder {
	nb.duration = beats
	for i := range nb.notes {
		nb.notes[i].Duration = beats
		nb.notes[i].StartTime = float32(i) * beats
	}
	return nb
}

// Velocity sets velocity for all notes
func (nb *NoteBuilder) Velocity(vel int32) *NoteBuilder {
	nb.velocity = vel
	for i := range nb.notes {
		nb.notes[i].Velocity = vel
	}
	return nb
}

// Pattern adjusts note timing based on pattern type
func (nb *NoteBuilder) Pattern(patternType string, swing, humanize int) *NoteBuilder {
	switch patternType {
	case "straight":
		// Already straight, do nothing
	}

	if swing > 0 {
		nb.applySwing(swing)
	}

	if humanize > 0 {
		nb.applyHumanize(humanize)
	}

	return nb
}

func (nb *NoteBuilder) applySwing(amount int) {
	// delay every other note
	swingFactor := float32(amount) / 100.0 * 0.1
	for i := range nb.notes {
		if i%2 == 1 {
			nb.notes[i].StartTime += swingFactor
		}
	}
}

func (nb NoteBuilder) applyHumanize(amount int) {
	factor := float32(amount) / 100.0
	for i := range nb.notes {
		// vary start time
		variation := factor * 0.05
		nb.notes[i].StartTime += variation * (float32(i%3) - 1) / 2

		// vary velocity
		velVariation := int32(float32(amount) * 0.2)
		nb.notes[i].Velocity += velVariation * (int32(i%3) - 1)
	}
}

func (nb *NoteBuilder) Build() {
	nb.clip.api.AddNotes(nb.clip.trackID, nb.clip.clipID, nb.notes...)
}
