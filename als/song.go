package als

import "github.com/matt0792/ableton-ctrl/oscclient"

// SongAPI provides methods for interacting with Ableton Live's Song API
type SongAPI struct {
	client *Client
}

// --- Methods ---

// CaptureMIDI triggers MIDI capture
func (s *SongAPI) CaptureMIDI() {
	s.client.send("/live/song/capture_midi")
}

// ContinuePlaying continues playing from the current position
func (s *SongAPI) ContinuePlaying() {
	s.client.send("/live/song/continue_playing")
}

// CreateAudioTrack creates a new audio track at the specified index
func (s *SongAPI) CreateAudioTrack(index int32) {
	s.client.send("/live/song/create_audio_track", index)
}

// CreateMIDITrack creates a new MIDI track at the specified index
func (s *SongAPI) CreateMIDITrack(index int32) {
	s.client.send("/live/song/create_midi_track", index)
}

// CreateReturnTrack creates a new return track
func (s *SongAPI) CreateReturnTrack() {
	s.client.send("/live/song/create_return_track")
}

// CreateScene creates a new scene at the specified index
func (s *SongAPI) CreateScene(index int32) {
	s.client.send("/live/song/create_scene", index)
}

// JumpToCuePoint jumps to the specified cue point
func (s *SongAPI) JumpToCuePoint(cuePoint int32) {
	s.client.send("/live/song/cue_point/jump", cuePoint)
}

// DeleteScene deletes the scene at the specified index
func (s *SongAPI) DeleteScene(sceneIndex int32) {
	s.client.send("/live/song/delete_scene", sceneIndex)
}

// DeleteReturnTrack deletes the return track at the specified index
func (s *SongAPI) DeleteReturnTrack(trackIndex int32) {
	s.client.send("/live/song/delete_return_track", trackIndex)
}

// DeleteTrack deletes the track at the specified index
func (s *SongAPI) DeleteTrack(trackIndex int32) {
	s.client.send("/live/song/delete_track", trackIndex)
}

// DuplicateScene duplicates the scene at the specified index
func (s *SongAPI) DuplicateScene(sceneIndex int32) {
	s.client.send("/live/song/duplicate_scene", sceneIndex)
}

// DuplicateTrack duplicates the track at the specified index
func (s *SongAPI) DuplicateTrack(trackIndex int32) {
	s.client.send("/live/song/duplicate_track", trackIndex)
}

// JumpBy jumps by the specified time
func (s *SongAPI) JumpBy(time float32) {
	s.client.send("/live/song/jump_by", time)
}

// JumpToNextCue jumps to the next cue point
func (s *SongAPI) JumpToNextCue() {
	s.client.send("/live/song/jump_to_next_cue")
}

// JumpToPrevCue jumps to the previous cue point
func (s *SongAPI) JumpToPrevCue() {
	s.client.send("/live/song/jump_to_prev_cue")
}

// Redo performs a redo operation
func (s *SongAPI) Redo() {
	s.client.send("/live/song/redo")
}

// StartPlaying starts playback
func (s *SongAPI) StartPlaying() {
	s.client.send("/live/song/start_playing")
}

// StopPlaying stops playback
func (s *SongAPI) StopPlaying() {
	s.client.send("/live/song/stop_playing")
}

// StopAllClips stops all playing clips
func (s *SongAPI) StopAllClips() {
	s.client.send("/live/song/stop_all_clips")
}

// TapTempo taps tempo
func (s *SongAPI) TapTempo() {
	s.client.send("/live/song/tap_tempo")
}

// TriggerSessionRecord triggers session recording
func (s *SongAPI) TriggerSessionRecord() {
	s.client.send("/live/song/trigger_session_record")
}

// Undo performs an undo operation
func (s *SongAPI) Undo() {
	s.client.send("/live/song/undo")
}

// --- Property Getters ---

// GetArrangementOverdub returns the arrangement overdub state
func (s *SongAPI) GetArrangementOverdub() bool {
	msg := s.client.send("/live/song/get/arrangement_overdub").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetBackToArranger returns the back to arranger state
func (s *SongAPI) GetBackToArranger() bool {
	msg := s.client.send("/live/song/get/back_to_arranger").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetCanRedo returns whether redo is available
func (s *SongAPI) GetCanRedo() bool {
	msg := s.client.send("/live/song/get/can_redo").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetCanUndo returns whether undo is available
func (s *SongAPI) GetCanUndo() bool {
	msg := s.client.send("/live/song/get/can_undo").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetClipTriggerQuantization returns the clip trigger quantization setting
func (s *SongAPI) GetClipTriggerQuantization() int32 {
	msg := s.client.send("/live/song/get/clip_trigger_quantization").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetCurrentSongTime returns the current song time
func (s *SongAPI) GetCurrentSongTime() float32 {
	msg := s.client.send("/live/song/get/current_song_time").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetGrooveAmount returns the groove amount
func (s *SongAPI) GetGrooveAmount() float32 {
	msg := s.client.send("/live/song/get/groove_amount").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetIsPlaying returns whether the song is playing
func (s *SongAPI) GetIsPlaying() bool {
	msg := s.client.send("/live/song/get/is_playing").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetLoop returns whether loop is enabled
func (s *SongAPI) GetLoop() bool {
	msg := s.client.send("/live/song/get/loop").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetLoopLength returns the loop length
func (s *SongAPI) GetLoopLength() float32 {
	msg := s.client.send("/live/song/get/loop_length").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetLoopStart returns the loop start position
func (s *SongAPI) GetLoopStart() float32 {
	msg := s.client.send("/live/song/get/loop_start").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetMetronome returns whether the metronome is on
func (s *SongAPI) GetMetronome() bool {
	msg := s.client.send("/live/song/get/metronome").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetMIDIRecordingQuantization returns the MIDI recording quantization setting
func (s *SongAPI) GetMIDIRecordingQuantization() int32 {
	msg := s.client.send("/live/song/get/midi_recording_quantization").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetNudgeDown returns the nudge down state
func (s *SongAPI) GetNudgeDown() bool {
	msg := s.client.send("/live/song/get/nudge_down").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetNudgeUp returns the nudge up state
func (s *SongAPI) GetNudgeUp() bool {
	msg := s.client.send("/live/song/get/nudge_up").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetPunchIn returns the punch in state
func (s *SongAPI) GetPunchIn() bool {
	msg := s.client.send("/live/song/get/punch_in").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetPunchOut returns the punch out state
func (s *SongAPI) GetPunchOut() bool {
	msg := s.client.send("/live/song/get/punch_out").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetRecordMode returns the record mode
func (s *SongAPI) GetRecordMode() bool {
	msg := s.client.send("/live/song/get/record_mode").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetSessionRecord returns the session record state
func (s *SongAPI) GetSessionRecord() bool {
	msg := s.client.send("/live/song/get/session_record").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetSessionRecordStatus returns the session record status
func (s *SongAPI) GetSessionRecordStatus() int32 {
	msg := s.client.send("/live/song/get/session_record_status").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetSignatureDenominator returns the time signature denominator
func (s *SongAPI) GetSignatureDenominator() int32 {
	msg := s.client.send("/live/song/get/signature_denominator").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetSignatureNumerator returns the time signature numerator
func (s *SongAPI) GetSignatureNumerator() int32 {
	msg := s.client.send("/live/song/get/signature_numerator").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetSongLength returns the song length
func (s *SongAPI) GetSongLength() float32 {
	msg := s.client.send("/live/song/get/song_length").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetTempo returns the current tempo in BPM
func (s *SongAPI) GetTempo() float32 {
	msg := s.client.send("/live/song/get/tempo").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

// GetNumScenes returns the number of scenes
func (s *SongAPI) GetNumScenes() int32 {
	msg := s.client.send("/live/song/get/num_scenes").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetNumTracks returns the number of tracks
func (s *SongAPI) GetNumTracks() int32 {
	msg := s.client.send("/live/song/get/num_tracks").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetTrackNames returns track names in the specified range
// If no range is specified, returns all track names
func (s *SongAPI) GetTrackNames(indexRange ...int32) []string {
	var msg *oscclient.Call
	if len(indexRange) == 2 {
		msg = s.client.send("/live/song/get/track_names", indexRange[0], indexRange[1])
	} else {
		msg = s.client.send("/live/song/get/track_names")
	}

	result := msg.Wait()
	names := make([]string, 0)
	for _, arg := range result.Arguments {
		if name, ok := arg.(string); ok {
			names = append(names, name)
		}
	}
	return names
}

// --- Property Setters ---

// SetArrangementOverdub sets the arrangement overdub state
func (s *SongAPI) SetArrangementOverdub(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/arrangement_overdub", val)
}

// SetBackToArranger sets the back to arranger state
func (s *SongAPI) SetBackToArranger(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/back_to_arranger", val)
}

// SetClipTriggerQuantization sets the clip trigger quantization
func (s *SongAPI) SetClipTriggerQuantization(quantization int32) {
	s.client.send("/live/song/set/clip_trigger_quantization", quantization)
}

// SetCurrentSongTime sets the current song time
func (s *SongAPI) SetCurrentSongTime(time float32) {
	s.client.send("/live/song/set/current_song_time", time)
}

// SetGrooveAmount sets the groove amount
func (s *SongAPI) SetGrooveAmount(amount float32) {
	s.client.send("/live/song/set/groove_amount", amount)
}

// SetLoop sets whether loop is enabled
func (s *SongAPI) SetLoop(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/loop", val)
}

// SetLoopLength sets the loop length
func (s *SongAPI) SetLoopLength(length float32) {
	s.client.send("/live/song/set/loop_length", length)
}

// SetLoopStart sets the loop start position
func (s *SongAPI) SetLoopStart(start float32) {
	s.client.send("/live/song/set/loop_start", start)
}

// SetMetronome sets whether the metronome is on
func (s *SongAPI) SetMetronome(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/metronome", val)
}

// SetMIDIRecordingQuantization sets the MIDI recording quantization
func (s *SongAPI) SetMIDIRecordingQuantization(quantization int32) {
	s.client.send("/live/song/set/midi_recording_quantization", quantization)
}

// SetNudgeDown sets the nudge down state
func (s *SongAPI) SetNudgeDown(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/nudge_down", val)
}

// SetNudgeUp sets the nudge up state
func (s *SongAPI) SetNudgeUp(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/nudge_up", val)
}

// SetPunchIn sets the punch in state
func (s *SongAPI) SetPunchIn(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/punch_in", val)
}

// SetPunchOut sets the punch out state
func (s *SongAPI) SetPunchOut(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/punch_out", val)
}

// SetRecordMode sets the record mode
func (s *SongAPI) SetRecordMode(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/record_mode", val)
}

// SetSessionRecord sets the session record state
func (s *SongAPI) SetSessionRecord(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/session_record", val)
}

// SetSignatureDenominator sets the time signature denominator
func (s *SongAPI) SetSignatureDenominator(denominator int32) {
	s.client.send("/live/song/set/signature_denominator", denominator)
}

// SetSignatureNumerator sets the time signature numerator
func (s *SongAPI) SetSignatureNumerator(numerator int32) {
	s.client.send("/live/song/set/signature_numerator", numerator)
}

// SetTempo sets the tempo in BPM
func (s *SongAPI) SetTempo(bpm float32) {
	s.client.send("/live/song/set/tempo", bpm)
}

// --- Listening Methods ---

// StartListenProperty starts listening for changes to the specified property
func (s *SongAPI) StartListenProperty(property string) {
	s.client.send("/live/song/start_listen/" + property)
}

// StopListenProperty stops listening for changes to the specified property
func (s *SongAPI) StopListenProperty(property string) {
	s.client.send("/live/song/stop_listen/" + property)
}

// StartListenBeat starts listening for beat changes
func (s *SongAPI) StartListenBeat() {
	s.client.send("/live/song/start_listen/beat")
}

// StopListenBeat stops listening for beat changes
func (s *SongAPI) StopListenBeat() {
	s.client.send("/live/song/stop_listen/beat")
}
