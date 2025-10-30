package als

// ClipSlotAPI provides methods for interacting with Ableton Live's Clip Slot API
type ClipSlotAPI struct {
	client *Client
}

// --- Methods ---

// Fire fires the clip in the specified clip slot
func (c *ClipSlotAPI) Fire(trackIndex, clipIndex int32) {
	c.client.send("/live/clip_slot/fire", trackIndex, clipIndex)
}

// CreateClip creates a clip in the specified clip slot with the given length
func (c *ClipSlotAPI) CreateClip(trackIndex, clipIndex int32, length float32) {
	c.client.send("/live/clip_slot/create_clip", trackIndex, clipIndex, length)
}

// DeleteClip deletes the clip in the specified clip slot
func (c *ClipSlotAPI) DeleteClip(trackIndex, clipIndex int32) {
	c.client.send("/live/clip_slot/delete_clip", trackIndex, clipIndex)
}

// DuplicateClipTo duplicates a clip to another clip slot
func (c *ClipSlotAPI) DuplicateClipTo(trackIndex, clipIndex, targetTrackIndex, targetClipIndex int32) {
	c.client.send("/live/clip_slot/duplicate_clip_to", trackIndex, clipIndex, targetTrackIndex, targetClipIndex)
}

// --- Property Getters ---

// GetHasClip returns whether the clip slot has a clip
func (c *ClipSlotAPI) GetHasClip(trackIndex, clipIndex int32) bool {
	msg := c.client.send("/live/clip_slot/get/has_clip", trackIndex, clipIndex).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetHasStopButton returns whether the clip slot has a stop button
func (c *ClipSlotAPI) GetHasStopButton(trackIndex, clipIndex int32) bool {
	msg := c.client.send("/live/clip_slot/get/has_stop_button", trackIndex, clipIndex).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val != 0
		}
	}
	return false
}

// --- Property Setters ---

// SetHasStopButton sets whether the clip slot has a stop button
func (c *ClipSlotAPI) SetHasStopButton(trackIndex, clipIndex int32, hasStopButton bool) {
	val := int32(0)
	if hasStopButton {
		val = 1
	}
	c.client.send("/live/clip_slot/set/has_stop_button", trackIndex, clipIndex, val)
}
