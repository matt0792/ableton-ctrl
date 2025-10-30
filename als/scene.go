package als

// SceneAPI provides methods for interacting with Ableton Live's Scene API
type SceneAPI struct {
	client *Client
}

// --- Methods ---

// Fire fires the specified scene
func (s *SceneAPI) Fire(sceneID int32) {
	s.client.send("/live/scene/fire", sceneID)
}

// FireAsSelected fires the specified scene as selected
func (s *SceneAPI) FireAsSelected(sceneID int32) {
	s.client.send("/live/scene/fire_as_selected", sceneID)
}

// FireSelected fires the currently selected scene
func (s *SceneAPI) FireSelected() {
	s.client.send("/live/scene/fire_selected")
}

// --- Property Getters ---

// GetColor returns the scene color
func (s *SceneAPI) GetColor(sceneID int32) int32 {
	msg := s.client.send("/live/scene/get/color", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

// GetColorIndex returns the scene color index
func (s *SceneAPI) GetColorIndex(sceneID int32) int32 {
	msg := s.client.send("/live/scene/get/color_index", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

// GetIsEmpty returns whether the scene is empty
func (s *SceneAPI) GetIsEmpty(sceneID int32) bool {
	msg := s.client.send("/live/scene/get/is_empty", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetIsTriggered returns whether the scene is triggered
func (s *SceneAPI) GetIsTriggered(sceneID int32) bool {
	msg := s.client.send("/live/scene/get/is_triggered", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetName returns the scene name
func (s *SceneAPI) GetName(sceneID int32) string {
	msg := s.client.send("/live/scene/get/name", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(string); ok {
			return val
		}
	}
	return ""
}

// GetTempo returns the scene tempo
func (s *SceneAPI) GetTempo(sceneID int32) float32 {
	msg := s.client.send("/live/scene/get/tempo", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(float32); ok {
			return val
		}
	}
	return 0
}

// GetTempoEnabled returns whether the scene tempo is enabled
func (s *SceneAPI) GetTempoEnabled(sceneID int32) bool {
	msg := s.client.send("/live/scene/get/tempo_enabled", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

// GetTimeSignatureNumerator returns the scene time signature numerator
func (s *SceneAPI) GetTimeSignatureNumerator(sceneID int32) int32 {
	msg := s.client.send("/live/scene/get/time_signature_numerator", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

// GetTimeSignatureDenominator returns the scene time signature denominator
func (s *SceneAPI) GetTimeSignatureDenominator(sceneID int32) int32 {
	msg := s.client.send("/live/scene/get/time_signature_denominator", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val
		}
	}
	return 0
}

// GetTimeSignatureEnabled returns whether the scene time signature is enabled
func (s *SceneAPI) GetTimeSignatureEnabled(sceneID int32) bool {
	msg := s.client.send("/live/scene/get/time_signature_enabled", sceneID).Wait()
	if len(msg.Arguments) >= 2 {
		if val, ok := msg.Arguments[1].(int32); ok {
			return val != 0
		}
	}
	return false
}

// --- Property Setters ---

// SetName sets the scene name
func (s *SceneAPI) SetName(sceneID int32, name string) {
	s.client.send("/live/scene/set/name", sceneID, name)
}

// SetColor sets the scene color
func (s *SceneAPI) SetColor(sceneID, color int32) {
	s.client.send("/live/scene/set/color", sceneID, color)
}

// SetColorIndex sets the scene color index
func (s *SceneAPI) SetColorIndex(sceneID, colorIndex int32) {
	s.client.send("/live/scene/set/color_index", sceneID, colorIndex)
}

// SetTempo sets the scene tempo
func (s *SceneAPI) SetTempo(sceneID int32, tempo float32) {
	s.client.send("/live/scene/set/tempo", sceneID, tempo)
}

// SetTempoEnabled sets whether the scene tempo is enabled
func (s *SceneAPI) SetTempoEnabled(sceneID int32, enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/scene/set/tempo_enabled", sceneID, val)
}

// SetTimeSignatureNumerator sets the scene time signature numerator
func (s *SceneAPI) SetTimeSignatureNumerator(sceneID, numerator int32) {
	s.client.send("/live/scene/set/time_signature_numerator", sceneID, numerator)
}

// SetTimeSignatureDenominator sets the scene time signature denominator
func (s *SceneAPI) SetTimeSignatureDenominator(sceneID, denominator int32) {
	s.client.send("/live/scene/set/time_signature_denominator", sceneID, denominator)
}

// SetTimeSignatureEnabled sets whether the scene time signature is enabled
func (s *SceneAPI) SetTimeSignatureEnabled(sceneID int32, enabled bool) {
	val := int32(0)
	if enabled {
		val = 1
	}
	s.client.send("/live/scene/set/time_signature_enabled", sceneID, val)
}

// --- Listening Methods ---

// StartListenProperty starts listening for changes to the specified property
func (s *SceneAPI) StartListenProperty(sceneIndex int32, property string) {
	s.client.send("/live/scene/start_listen/"+property, sceneIndex)
}

// StopListenProperty stops listening for changes to the specified property
func (s *SceneAPI) StopListenProperty(sceneIndex int32, property string) {
	s.client.send("/live/scene/stop_listen/"+property, sceneIndex)
}
