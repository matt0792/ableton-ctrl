package als

// ViewAPI provides methods for interacting with Ableton Live's View API
type ViewAPI struct {
	client *Client
}

// --- Property Getters ---

// GetSelectedScene returns the index of the selected scene
func (v *ViewAPI) GetSelectedScene() int32 {
	msg := v.client.send("/live/view/get/selected_scene").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetSelectedTrack returns the index of the selected track
func (v *ViewAPI) GetSelectedTrack() int32 {
	msg := v.client.send("/live/view/get/selected_track").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

// GetSelectedClip returns the track and scene indices of the selected clip
func (v *ViewAPI) GetSelectedClip() (trackIndex, sceneIndex int32) {
	msg := v.client.send("/live/view/get/selected_clip").Wait()
	if len(msg.Arguments) >= 2 {
		if track, ok := msg.Arguments[0].(int32); ok {
			trackIndex = track
		}
		if scene, ok := msg.Arguments[1].(int32); ok {
			sceneIndex = scene
		}
	}
	return
}

// GetSelectedDevice returns the track and device indices of the selected device
func (v *ViewAPI) GetSelectedDevice() (trackIndex, deviceIndex int32) {
	msg := v.client.send("/live/view/get/selected_device").Wait()
	if len(msg.Arguments) >= 2 {
		if track, ok := msg.Arguments[0].(int32); ok {
			trackIndex = track
		}
		if device, ok := msg.Arguments[1].(int32); ok {
			deviceIndex = device
		}
	}
	return
}

// --- Property Setters ---

// SetSelectedScene sets the selected scene
func (v *ViewAPI) SetSelectedScene(sceneIndex int32) {
	v.client.send("/live/view/set/selected_scene", sceneIndex)
}

// SetSelectedTrack sets the selected track
func (v *ViewAPI) SetSelectedTrack(trackIndex int32) {
	v.client.send("/live/view/set/selected_track", trackIndex)
}

// SetSelectedClip sets the selected clip
func (v *ViewAPI) SetSelectedClip(trackIndex, sceneIndex int32) {
	v.client.send("/live/view/set/selected_clip", trackIndex, sceneIndex)
}

// SetSelectedDevice sets the selected device
func (v *ViewAPI) SetSelectedDevice(trackIndex, deviceIndex int32) {
	v.client.send("/live/view/set/selected_device", trackIndex, deviceIndex)
}

// --- Listening Methods ---

// StartListenSelectedScene starts listening for selected scene changes
func (v *ViewAPI) StartListenSelectedScene() {
	v.client.send("/live/view/start_listen/selected_scene")
}

// StopListenSelectedScene stops listening for selected scene changes
func (v *ViewAPI) StopListenSelectedScene() {
	v.client.send("/live/view/stop_listen/selected_scene")
}

// StartListenSelectedTrack starts listening for selected track changes
func (v *ViewAPI) StartListenSelectedTrack() {
	v.client.send("/live/view/start_listen/selected_track")
}

// StopListenSelectedTrack stops listening for selected track changes
func (v *ViewAPI) StopListenSelectedTrack() {
	v.client.send("/live/view/stop_listen/selected_track")
}
