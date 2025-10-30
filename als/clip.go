package als

import "github.com/matt0792/ableton-ctrl/oscclient"

// ClipAPI provides methods for interacting with Ableton Live's Clip API
type ClipAPI struct {
	client *Client
}

// Note represents a MIDI note in a clip
type Note struct {
	Pitch     int32
	StartTime float32
	Duration  float32
	Velocity  int32
	Mute      bool
}

// --- Methods ---

// Fire fires the specified clip
func (c *ClipAPI) Fire(trackID, clipID int32) {
	c.client.send("/live/clip/fire", trackID, clipID)
}

// Stop stops the specified clip
func (c *ClipAPI) Stop(trackID, clipID int32) {
	c.client.send("/live/clip/stop", trackID, clipID)
}

// DuplicateLoop duplicates the loop in the specified clip
func (c *ClipAPI) DuplicateLoop(trackID, clipID int32) {
	c.client.send("/live/clip/duplicate_loop", trackID, clipID)
}

// GetNotes returns notes from the clip within the specified range
// If no range is specified, returns all notes
func (c *ClipAPI) GetNotes(trackID, clipID int32, rangeParams ...int32) []Note {
	var msg *oscclient.Call
	if len(rangeParams) == 4 {
		// startPitch, pitchSpan, startTime, timeSpan
		msg = c.client.send("/live/clip/get/notes", trackID, clipID,
			rangeParams[0], rangeParams[1], rangeParams[2], rangeParams[3])
	} else {
		msg = c.client.send("/live/clip/get/notes", trackID, clipID)
	}

	result := msg.Wait()
	notes := make([]Note, 0)

	// Notes come in groups of 5: pitch, start_time, duration, velocity, mute
	for i := 0; i+4 < len(result.Arguments); i += 5 {
		note := Note{}
		if pitch, ok := result.Arguments[i].(int32); ok {
			note.Pitch = pitch
		}
		if startTime, ok := result.Arguments[i+1].(float32); ok {
			note.StartTime = startTime
		}
		if duration, ok := result.Arguments[i+2].(float32); ok {
			note.Duration = duration
		}
		if velocity, ok := result.Arguments[i+3].(int32); ok {
			note.Velocity = velocity
		}
		if mute, ok := result.Arguments[i+4].(int32); ok {
			note.Mute = mute != 0
		}
		notes = append(notes, note)
	}

	return notes
}

// AddNotes adds notes to the clip
func (c *ClipAPI) AddNotes(trackID, clipID int32, notes ...Note) {
	params := []any{trackID, clipID}
	for _, note := range notes {
		muteVal := int32(0)
		if note.Mute {
			muteVal = 1
		}
		params = append(params, note.Pitch, note.StartTime, note.Duration, note.Velocity, muteVal)
	}
	c.client.send("/live/clip/add/notes", params...)
}

// RemoveNotes removes notes from the clip within the specified range
func (c *ClipAPI) RemoveNotes(trackID, clipID, startPitch, pitchSpan int32, startTime, timeSpan float32) {
	c.client.send("/live/clip/remove/notes", trackID, clipID, startPitch, pitchSpan, startTime, timeSpan)
}

// --- Property Getters ---

// GetColor returns the clip color
func (c *ClipAPI) GetColor(trackID, clipID int32) int32 {
	msg := c.client.send("/live/clip/get/color", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val
		}
	}
	return 0
}

// GetName returns the clip name
func (c *ClipAPI) GetName(trackID, clipID int32) string {
	msg := c.client.send("/live/clip/get/name", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(string); ok {
			return val
		}
	}
	return ""
}

// GetGain returns the clip gain
func (c *ClipAPI) GetGain(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/gain", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetLength returns the clip length
func (c *ClipAPI) GetLength(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/length", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetPitchCoarse returns the clip pitch coarse adjustment in semitones
func (c *ClipAPI) GetPitchCoarse(trackID, clipID int32) int32 {
	msg := c.client.send("/live/clip/get/pitch_coarse", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val
		}
	}
	return 0
}

// GetPitchFine returns the clip pitch fine adjustment in cents
func (c *ClipAPI) GetPitchFine(trackID, clipID int32) int32 {
	msg := c.client.send("/live/clip/get/pitch_fine", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val
		}
	}
	return 0
}

// GetFilePath returns the clip file path
func (c *ClipAPI) GetFilePath(trackID, clipID int32) string {
	msg := c.client.send("/live/clip/get/file_path", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(string); ok {
			return val
		}
	}
	return ""
}

// GetIsAudioClip returns whether the clip is an audio clip
func (c *ClipAPI) GetIsAudioClip(trackID, clipID int32) bool {
	msg := c.client.send("/live/clip/get/is_audio_clip", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetIsMIDIClip returns whether the clip is a MIDI clip
func (c *ClipAPI) GetIsMIDIClip(trackID, clipID int32) bool {
	msg := c.client.send("/live/clip/get/is_midi_clip", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetIsPlaying returns whether the clip is playing
func (c *ClipAPI) GetIsPlaying(trackID, clipID int32) bool {
	msg := c.client.send("/live/clip/get/is_playing", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetIsRecording returns whether the clip is recording
func (c *ClipAPI) GetIsRecording(trackID, clipID int32) bool {
	msg := c.client.send("/live/clip/get/is_recording", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetPlayingPosition returns the clip playing position
func (c *ClipAPI) GetPlayingPosition(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/playing_position", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetLoopStart returns the clip loop start position
func (c *ClipAPI) GetLoopStart(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/loop_start", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetLoopEnd returns the clip loop end position
func (c *ClipAPI) GetLoopEnd(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/loop_end", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetWarping returns whether warping is enabled
func (c *ClipAPI) GetWarping(trackID, clipID int32) bool {
	msg := c.client.send("/live/clip/get/warping", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetStartMarker returns the clip start marker position
func (c *ClipAPI) GetStartMarker(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/start_marker", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// GetEndMarker returns the clip end marker position
func (c *ClipAPI) GetEndMarker(trackID, clipID int32) float32 {
	msg := c.client.send("/live/clip/get/end_marker", trackID, clipID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

// --- Property Setters ---

// SetColor sets the clip color
func (c *ClipAPI) SetColor(trackID, clipID, color int32) {
	c.client.send("/live/clip/set/color", trackID, clipID, color)
}

// SetName sets the clip name
func (c *ClipAPI) SetName(trackID, clipID int32, name string) {
	c.client.send("/live/clip/set/name", trackID, clipID, name)
}

// SetGain sets the clip gain
func (c *ClipAPI) SetGain(trackID, clipID int32, gain float32) {
	c.client.send("/live/clip/set/gain", trackID, clipID, gain)
}

// SetPitchCoarse sets the clip pitch coarse adjustment in semitones
func (c *ClipAPI) SetPitchCoarse(trackID, clipID, semitones int32) {
	c.client.send("/live/clip/set/pitch_coarse", trackID, clipID, semitones)
}

// SetPitchFine sets the clip pitch fine adjustment in cents
func (c *ClipAPI) SetPitchFine(trackID, clipID, cents int32) {
	c.client.send("/live/clip/set/pitch_fine", trackID, clipID, cents)
}

// SetLoopStart sets the clip loop start position
func (c *ClipAPI) SetLoopStart(trackID, clipID int32, loopStart float32) {
	c.client.send("/live/clip/set/loop_start", trackID, clipID, loopStart)
}

// SetLoopEnd sets the clip loop end position
func (c *ClipAPI) SetLoopEnd(trackID, clipID int32, loopEnd float32) {
	c.client.send("/live/clip/set/loop_end", trackID, clipID, loopEnd)
}

// SetWarping sets whether warping is enabled
func (c *ClipAPI) SetWarping(trackID, clipID int32, enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	c.client.send("/live/clip/set/warping", trackID, clipID, val)
}

// SetStartMarker sets the clip start marker position
func (c *ClipAPI) SetStartMarker(trackID, clipID int32, startMarker float32) {
	c.client.send("/live/clip/set/start_marker", trackID, clipID, startMarker)
}

// SetEndMarker sets the clip end marker position
func (c *ClipAPI) SetEndMarker(trackID, clipID int32, endMarker float32) {
	c.client.send("/live/clip/set/end_marker", trackID, clipID, endMarker)
}

// --- Listening Methods ---

// StartListenPlayingPosition starts listening for playing position changes
func (c *ClipAPI) StartListenPlayingPosition(trackID, clipID int32) {
	c.client.send("/live/clip/start_listen/playing_position", trackID, clipID)
}

// StopListenPlayingPosition stops listening for playing position changes
func (c *ClipAPI) StopListenPlayingPosition(trackID, clipID int32) {
	c.client.send("/live/clip/stop_listen/playing_position", trackID, clipID)
}
