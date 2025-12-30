package als

// ViewAPI provides methods for interacting with Ableton Live's View API.
type ViewAPI struct {
	client *Client
}

// --- Property Getters ---

func (v *ViewAPI) GetSelectedScene() int32 {
	msg := v.client.send("/live/view/get/selected_scene").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

func (v *ViewAPI) GetSelectedTrack() int32 {
	msg := v.client.send("/live/view/get/selected_track").Wait()
	if len(msg.Arguments) > 0 {
		if val, ok := msg.Arguments[0].(int32); ok {
			return val
		}
	}
	return 0
}

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

func (v *ViewAPI) SetSelectedScene(sceneIndex int32) {
	v.client.send("/live/view/set/selected_scene", sceneIndex)
}

func (v *ViewAPI) SetSelectedTrack(trackIndex int32) {
	v.client.send("/live/view/set/selected_track", trackIndex)
}

func (v *ViewAPI) SetSelectedClip(trackIndex, sceneIndex int32) {
	v.client.send("/live/view/set/selected_clip", trackIndex, sceneIndex)
}

func (v *ViewAPI) SetSelectedDevice(trackIndex, deviceIndex int32) {
	v.client.send("/live/view/set/selected_device", trackIndex, deviceIndex)
}

// --- Listening Methods ---

func (v *ViewAPI) StartListenSelectedScene() {
	v.client.send("/live/view/start_listen/selected_scene")
}

func (v *ViewAPI) StopListenSelectedScene() {
	v.client.send("/live/view/stop_listen/selected_scene")
}

func (v *ViewAPI) StartListenSelectedTrack() {
	v.client.send("/live/view/start_listen/selected_track")
}

func (v *ViewAPI) StopListenSelectedTrack() {
	v.client.send("/live/view/stop_listen/selected_track")
}
