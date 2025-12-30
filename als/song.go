package als

import "github.com/matt0792/ableton-ctrl/oscclient"

// SongAPI provides methods for interacting with Ableton Live's Song API.
type SongAPI struct {
	client *Client
}

// --- Methods ---

func (s *SongAPI) CaptureMIDI() {
	s.client.send("/live/song/capture_midi")
}

func (s *SongAPI) ContinuePlaying() {
	s.client.send("/live/song/continue_playing")
}

func (s *SongAPI) CreateAudioTrack(index int32) {
	s.client.send("/live/song/create_audio_track", index)
}

func (s *SongAPI) CreateMIDITrack(index int32) {
	s.client.send("/live/song/create_midi_track", index)
}

func (s *SongAPI) CreateReturnTrack() {
	s.client.send("/live/song/create_return_track")
}

func (s *SongAPI) CreateScene(index int32) {
	s.client.send("/live/song/create_scene", index)
}

func (s *SongAPI) JumpToCuePoint(cuePoint int32) {
	s.client.send("/live/song/cue_point/jump", cuePoint)
}

func (s *SongAPI) DeleteScene(sceneIndex int32) {
	s.client.send("/live/song/delete_scene", sceneIndex)
}

func (s *SongAPI) DeleteReturnTrack(trackIndex int32) {
	s.client.send("/live/song/delete_return_track", trackIndex)
}

func (s *SongAPI) DeleteTrack(trackIndex int32) {
	s.client.send("/live/song/delete_track", trackIndex)
}

func (s *SongAPI) DuplicateScene(sceneIndex int32) {
	s.client.send("/live/song/duplicate_scene", sceneIndex)
}

func (s *SongAPI) DuplicateTrack(trackIndex int32) {
	s.client.send("/live/song/duplicate_track", trackIndex)
}

func (s *SongAPI) JumpBy(time float32) {
	s.client.send("/live/song/jump_by", time)
}

func (s *SongAPI) JumpToNextCue() {
	s.client.send("/live/song/jump_to_next_cue")
}

func (s *SongAPI) JumpToPrevCue() {
	s.client.send("/live/song/jump_to_prev_cue")
}

func (s *SongAPI) Redo() {
	s.client.send("/live/song/redo")
}

func (s *SongAPI) StartPlaying() {
	s.client.send("/live/song/start_playing")
}

func (s *SongAPI) StopPlaying() {
	s.client.send("/live/song/stop_playing")
}

func (s *SongAPI) StopAllClips() {
	s.client.send("/live/song/stop_all_clips")
}

func (s *SongAPI) TapTempo() {
	s.client.send("/live/song/tap_tempo")
}

func (s *SongAPI) TriggerSessionRecord() {
	s.client.send("/live/song/trigger_session_record")
}

func (s *SongAPI) Undo() {
	s.client.send("/live/song/undo")
}

// --- Property Getters ---

func (s *SongAPI) GetArrangementOverdub() bool {
	msg := s.client.send("/live/song/get/arrangement_overdub").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetBackToArranger() bool {
	msg := s.client.send("/live/song/get/back_to_arranger").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetCanRedo() bool {
	msg := s.client.send("/live/song/get/can_redo").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetCanUndo() bool {
	msg := s.client.send("/live/song/get/can_undo").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetClipTriggerQuantization() int32 {
	msg := s.client.send("/live/song/get/clip_trigger_quantization").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetCurrentSongTime() float32 {
	msg := s.client.send("/live/song/get/current_song_time").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetGrooveAmount() float32 {
	msg := s.client.send("/live/song/get/groove_amount").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetIsPlaying() bool {
	msg := s.client.send("/live/song/get/is_playing").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetLoop() bool {
	msg := s.client.send("/live/song/get/loop").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetLoopLength() float32 {
	msg := s.client.send("/live/song/get/loop_length").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetLoopStart() float32 {
	msg := s.client.send("/live/song/get/loop_start").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(float32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetMetronome() bool {
	msg := s.client.send("/live/song/get/metronome").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetMIDIRecordingQuantization() int32 {
	msg := s.client.send("/live/song/get/midi_recording_quantization").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetNudgeDown() bool {
	msg := s.client.send("/live/song/get/nudge_down").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetNudgeUp() bool {
	msg := s.client.send("/live/song/get/nudge_up").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetPunchIn() bool {
	msg := s.client.send("/live/song/get/punch_in").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetPunchOut() bool {
	msg := s.client.send("/live/song/get/punch_out").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetRecordMode() bool {
	msg := s.client.send("/live/song/get/record_mode").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetSessionRecord() bool {
	msg := s.client.send("/live/song/get/session_record").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (s *SongAPI) GetSessionRecordStatus() int32 {
	msg := s.client.send("/live/song/get/session_record_status").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetSignatureDenominator() int32 {
	msg := s.client.send("/live/song/get/signature_denominator").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

func (s *SongAPI) GetSignatureNumerator() int32 {
	msg := s.client.send("/live/song/get/signature_numerator").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

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

func (s *SongAPI) GetNumScenes() int32 {
	msg := s.client.send("/live/song/get/num_scenes").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

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

func (s *SongAPI) SetArrangementOverdub(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/arrangement_overdub", val)
}

func (s *SongAPI) SetBackToArranger(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/back_to_arranger", val)
}

func (s *SongAPI) SetClipTriggerQuantization(quantization int32) {
	s.client.send("/live/song/set/clip_trigger_quantization", quantization)
}

func (s *SongAPI) SetCurrentSongTime(time float32) {
	s.client.send("/live/song/set/current_song_time", time)
}

func (s *SongAPI) SetGrooveAmount(amount float32) {
	s.client.send("/live/song/set/groove_amount", amount)
}

func (s *SongAPI) SetLoop(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/loop", val)
}

func (s *SongAPI) SetLoopLength(length float32) {
	s.client.send("/live/song/set/loop_length", length)
}

func (s *SongAPI) SetLoopStart(start float32) {
	s.client.send("/live/song/set/loop_start", start)
}

func (s *SongAPI) SetMetronome(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/metronome", val)
}

func (s *SongAPI) SetMIDIRecordingQuantization(quantization int32) {
	s.client.send("/live/song/set/midi_recording_quantization", quantization)
}

func (s *SongAPI) SetNudgeDown(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/nudge_down", val)
}

func (s *SongAPI) SetNudgeUp(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/nudge_up", val)
}

func (s *SongAPI) SetPunchIn(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/punch_in", val)
}

func (s *SongAPI) SetPunchOut(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/punch_out", val)
}

func (s *SongAPI) SetRecordMode(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/record_mode", val)
}

func (s *SongAPI) SetSessionRecord(enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/song/set/session_record", val)
}

func (s *SongAPI) SetSignatureDenominator(denominator int32) {
	s.client.send("/live/song/set/signature_denominator", denominator)
}

func (s *SongAPI) SetSignatureNumerator(numerator int32) {
	s.client.send("/live/song/set/signature_numerator", numerator)
}

// SetTempo sets the tempo in BPM
func (s *SongAPI) SetTempo(bpm float32) {
	s.client.send("/live/song/set/tempo", bpm)
}

// --- Listening Methods ---

func (s *SongAPI) StartListenProperty(property string) {
	s.client.send("/live/song/start_listen/" + property)
}

func (s *SongAPI) StopListenProperty(property string) {
	s.client.send("/live/song/stop_listen/" + property)
}

func (s *SongAPI) StartListenBeat() {
	s.client.send("/live/song/start_listen/beat")
}

func (s *SongAPI) StopListenBeat() {
	s.client.send("/live/song/stop_listen/beat")
}
