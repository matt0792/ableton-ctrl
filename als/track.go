package als

// TrackAPI provides methods for interacting with Ableton Live's Track API.
type TrackAPI struct {
	client *Client
}

// --- Methods ---

func (t *TrackAPI) StopAllClips(trackID int32) {
	t.client.send("/live/track/stop_all_clips", trackID)
}

// --- Property Getters ---

func (t *TrackAPI) GetArm(trackID int32) bool {
	msg := t.client.send("/live/track/get/arm", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetAvailableInputRoutingChannels(trackID int32) []string {
	msg := t.client.send("/live/track/get/available_input_routing_channels", trackID).Wait()
	channels := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if channel, ok := msg.Arguments[i].(string); ok {
			channels = append(channels, channel)
		}
	}
	return channels
}

func (t *TrackAPI) GetAvailableInputRoutingTypes(trackID int32) []string {
	msg := t.client.send("/live/track/get/available_input_routing_types", trackID).Wait()
	types := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if typ, ok := msg.Arguments[i].(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

func (t *TrackAPI) GetAvailableOutputRoutingChannels(trackID int32) []string {
	msg := t.client.send("/live/track/get/available_output_routing_channels", trackID).Wait()
	channels := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if channel, ok := msg.Arguments[i].(string); ok {
			channels = append(channels, channel)
		}
	}
	return channels
}

func (t *TrackAPI) GetAvailableOutputRoutingTypes(trackID int32) []string {
	msg := t.client.send("/live/track/get/available_output_routing_types", trackID).Wait()
	types := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if typ, ok := msg.Arguments[i].(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

func (t *TrackAPI) GetCanBeArmed(trackID int32) bool {
	msg := t.client.send("/live/track/get/can_be_armed", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetColor(trackID int32) int32 {
	msg := t.client.send("/live/track/get/color", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetColorIndex(trackID int32) int32 {
	msg := t.client.send("/live/track/get/color_index", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetCurrentMonitoringState(trackID int32) int32 {
	msg := t.client.send("/live/track/get/current_monitoring_state", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetFiredSlotIndex(trackID int32) int32 {
	msg := t.client.send("/live/track/get/fired_slot_index", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return -1
}

func (t *TrackAPI) GetFoldState(trackID int32) bool {
	msg := t.client.send("/live/track/get/fold_state", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetHasAudioInput(trackID int32) bool {
	msg := t.client.send("/live/track/get/has_audio_input", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetHasAudioOutput(trackID int32) bool {
	msg := t.client.send("/live/track/get/has_audio_output", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetHasMIDIInput(trackID int32) bool {
	msg := t.client.send("/live/track/get/has_midi_input", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetHasMIDIOutput(trackID int32) bool {
	msg := t.client.send("/live/track/get/has_midi_output", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetInputRoutingChannel(trackID int32) string {
	msg := t.client.send("/live/track/get/input_routing_channel", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

func (t *TrackAPI) GetInputRoutingType(trackID int32) string {
	msg := t.client.send("/live/track/get/input_routing_type", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

func (t *TrackAPI) GetOutputRoutingChannel(trackID int32) string {
	msg := t.client.send("/live/track/get/output_routing_channel", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

func (t *TrackAPI) GetOutputMeterLeft(trackID int32) float32 {
	msg := t.client.send("/live/track/get/output_meter_left", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetOutputMeterLevel(trackID int32) float32 {
	msg := t.client.send("/live/track/get/output_meter_level", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetOutputMeterRight(trackID int32) float32 {
	msg := t.client.send("/live/track/get/output_meter_right", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetOutputRoutingType(trackID int32) string {
	msg := t.client.send("/live/track/get/output_routing_type", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

func (t *TrackAPI) GetIsFoldable(trackID int32) bool {
	msg := t.client.send("/live/track/get/is_foldable", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetIsGrouped(trackID int32) bool {
	msg := t.client.send("/live/track/get/is_grouped", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetIsVisible(trackID int32) bool {
	msg := t.client.send("/live/track/get/is_visible", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetMute(trackID int32) bool {
	msg := t.client.send("/live/track/get/mute", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetName(trackID int32) string {
	msg := t.client.send("/live/track/get/name", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

func (t *TrackAPI) GetPanning(trackID int32) float32 {
	msg := t.client.send("/live/track/get/panning", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetPlayingSlotIndex(trackID int32) int32 {
	msg := t.client.send("/live/track/get/playing_slot_index", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return -1
}

func (t *TrackAPI) GetSend(trackID, sendID int32) float32 {
	msg := t.client.send("/live/track/get/send", trackID, sendID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(float32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetSolo(trackID int32) bool {
	msg := t.client.send("/live/track/get/solo", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

func (t *TrackAPI) GetVolume(trackID int32) float32 {
	msg := t.client.send("/live/track/get/volume", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

// --- Property Setters ---

func (t *TrackAPI) SetArm(trackID int32, armed bool) {
	val := int32(0)
	if armed {
		val = 1
	}
	t.client.send("/live/track/set/arm", trackID, val)
}

func (t *TrackAPI) SetColor(trackID, color int32) {
	t.client.send("/live/track/set/color", trackID, color)
}

func (t *TrackAPI) SetColorIndex(trackID, colorIndex int32) {
	t.client.send("/live/track/set/color_index", trackID, colorIndex)
}

func (t *TrackAPI) SetCurrentMonitoringState(trackID, state int32) {
	t.client.send("/live/track/set/current_monitoring_state", trackID, state)
}

func (t *TrackAPI) SetFoldState(trackID int32, folded bool) {
	val := int32(0)
	if folded {
		val = 1
	}
	t.client.send("/live/track/set/fold_state", trackID, val)
}

func (t *TrackAPI) SetInputRoutingChannel(trackID int32, channel string) {
	t.client.send("/live/track/set/input_routing_channel", trackID, channel)
}

func (t *TrackAPI) SetInputRoutingType(trackID int32, routingType string) {
	t.client.send("/live/track/set/input_routing_type", trackID, routingType)
}

func (t *TrackAPI) SetMute(trackID int32, muted bool) {
	val := int32(0)
	if muted {
		val = 1
	}
	t.client.send("/live/track/set/mute", trackID, val)
}

func (t *TrackAPI) SetName(trackID int32, name string) {
	t.client.send("/live/track/set/name", trackID, name)
}

func (t *TrackAPI) SetOutputRoutingChannel(trackID int32, channel string) {
	t.client.send("/live/track/set/output_routing_channel", trackID, channel)
}

func (t *TrackAPI) SetOutputRoutingType(trackID int32, routingType string) {
	t.client.send("/live/track/set/output_routing_type", trackID, routingType)
}

func (t *TrackAPI) SetPanning(trackID int32, panning float32) {
	t.client.send("/live/track/set/panning", trackID, panning)
}

func (t *TrackAPI) SetSend(trackID, sendID int32, value float32) {
	t.client.send("/live/track/set/send", trackID, sendID, value)
}

func (t *TrackAPI) SetSolo(trackID int32, soloed bool) {
	val := int32(0)
	if soloed {
		val = 1
	}
	t.client.send("/live/track/set/solo", trackID, val)
}

func (t *TrackAPI) SetVolume(trackID int32, volume float32) {
	t.client.send("/live/track/set/volume", trackID, volume)
}

// --- Clip Properties ---

func (t *TrackAPI) GetClipsName(trackID int32) []string {
	msg := t.client.send("/live/track/get/clips/name", trackID).Wait()
	names := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if name, ok := msg.Arguments[i].(string); ok {
			names = append(names, name)
		}
	}
	return names
}

func (t *TrackAPI) GetClipsLength(trackID int32) []float32 {
	msg := t.client.send("/live/track/get/clips/length", trackID).Wait()
	lengths := make([]float32, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if length, ok := msg.Arguments[i].(float32); ok {
			lengths = append(lengths, length)
		}
	}
	return lengths
}

func (t *TrackAPI) GetClipsColor(trackID int32) []int32 {
	msg := t.client.send("/live/track/get/clips/color", trackID).Wait()
	colors := make([]int32, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if color, ok := msg.Arguments[i].(int32); ok {
			colors = append(colors, color)
		}
	}
	return colors
}

func (t *TrackAPI) GetArrangementClipsName(trackID int32) []string {
	msg := t.client.send("/live/track/get/arrangement_clips/name", trackID).Wait()
	names := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if name, ok := msg.Arguments[i].(string); ok {
			names = append(names, name)
		}
	}
	return names
}

func (t *TrackAPI) GetArrangementClipsLength(trackID int32) []float32 {
	msg := t.client.send("/live/track/get/arrangement_clips/length", trackID).Wait()
	lengths := make([]float32, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if length, ok := msg.Arguments[i].(float32); ok {
			lengths = append(lengths, length)
		}
	}
	return lengths
}

func (t *TrackAPI) GetArrangementClipsStartTime(trackID int32) []float32 {
	msg := t.client.send("/live/track/get/arrangement_clips/start_time", trackID).Wait()
	startTimes := make([]float32, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if startTime, ok := msg.Arguments[i].(float32); ok {
			startTimes = append(startTimes, startTime)
		}
	}
	return startTimes
}

// --- Device Properties ---

func (t *TrackAPI) GetNumDevices(trackID int32) int32 {
	msg := t.client.send("/live/track/get/num_devices", trackID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

func (t *TrackAPI) GetDevicesName(trackID int32) []string {
	msg := t.client.send("/live/track/get/devices/name", trackID).Wait()
	names := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if name, ok := msg.Arguments[i].(string); ok {
			names = append(names, name)
		}
	}
	return names
}

func (t *TrackAPI) GetDevicesType(trackID int32) []string {
	msg := t.client.send("/live/track/get/devices/type", trackID).Wait()
	types := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if typ, ok := msg.Arguments[i].(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

func (t *TrackAPI) GetDevicesClassName(trackID int32) []string {
	msg := t.client.send("/live/track/get/devices/class_name", trackID).Wait()
	classNames := make([]string, 0)
	// Skip first argument which is track ID
	for i := 1; i < len(msg.Arguments); i++ {
		if className, ok := msg.Arguments[i].(string); ok {
			classNames = append(classNames, className)
		}
	}
	return classNames
}

// --- Listening Methods ---

func (t *TrackAPI) StartListenProperty(trackIndex int32, property string) {
	t.client.send("/live/track/start_listen/"+property, trackIndex)
}

func (t *TrackAPI) StopListenProperty(trackIndex int32, property string) {
	t.client.send("/live/track/stop_listen/"+property, trackIndex)
}
