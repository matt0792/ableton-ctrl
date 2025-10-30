package als

// DeviceAPI provides methods for interacting with Ableton Live's Device API
type DeviceAPI struct {
	client *Client
}

// --- Property Getters ---

// GetName returns the device name
func (d *DeviceAPI) GetName(trackID, deviceID int32) string {
	msg := d.client.send("/live/device/get/name", trackID, deviceID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(string); ok {
			return val
		}
	}
	return ""
}

// GetClassName returns the device class name
func (d *DeviceAPI) GetClassName(trackID, deviceID int32) string {
	msg := d.client.send("/live/device/get/class_name", trackID, deviceID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(string); ok {
			return val
		}
	}
	return ""
}

// GetType returns the device type
func (d *DeviceAPI) GetType(trackID, deviceID int32) string {
	msg := d.client.send("/live/device/get/type", trackID, deviceID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(string); ok {
			return val
		}
	}
	return ""
}

// GetNumParameters returns the number of parameters
func (d *DeviceAPI) GetNumParameters(trackID, deviceID int32) int32 {
	msg := d.client.send("/live/device/get/num_parameters", trackID, deviceID).Wait()
	if len(msg.Arguments) >= 3 {
		if val, ok := msg.Arguments[2].(int32); ok {
			return val
		}
	}
	return 0
}

// GetParametersName returns all parameter names
func (d *DeviceAPI) GetParametersName(trackID, deviceID int32) []string {
	msg := d.client.send("/live/device/get/parameters/name", trackID, deviceID).Wait()
	names := make([]string, 0)
	// Skip first two arguments which are track ID and device ID
	for i := 2; i < len(msg.Arguments); i++ {
		if name, ok := msg.Arguments[i].(string); ok {
			names = append(names, name)
		}
	}
	return names
}

// GetParametersValue returns all parameter values
func (d *DeviceAPI) GetParametersValue(trackID, deviceID int32) []float32 {
	msg := d.client.send("/live/device/get/parameters/value", trackID, deviceID).Wait()
	values := make([]float32, 0)
	// Skip first two arguments which are track ID and device ID
	for i := 2; i < len(msg.Arguments); i++ {
		if value, ok := msg.Arguments[i].(float32); ok {
			values = append(values, value)
		}
	}
	return values
}

// GetParametersMin returns all parameter minimum values
func (d *DeviceAPI) GetParametersMin(trackID, deviceID int32) []float32 {
	msg := d.client.send("/live/device/get/parameters/min", trackID, deviceID).Wait()
	mins := make([]float32, 0)
	// Skip first two arguments which are track ID and device ID
	for i := 2; i < len(msg.Arguments); i++ {
		if min, ok := msg.Arguments[i].(float32); ok {
			mins = append(mins, min)
		}
	}
	return mins
}

// GetParametersMax returns all parameter maximum values
func (d *DeviceAPI) GetParametersMax(trackID, deviceID int32) []float32 {
	msg := d.client.send("/live/device/get/parameters/max", trackID, deviceID).Wait()
	maxs := make([]float32, 0)
	// Skip first two arguments which are track ID and device ID
	for i := 2; i < len(msg.Arguments); i++ {
		if max, ok := msg.Arguments[i].(float32); ok {
			maxs = append(maxs, max)
		}
	}
	return maxs
}

// GetParametersIsQuantized returns whether parameters are quantized
func (d *DeviceAPI) GetParametersIsQuantized(trackID, deviceID int32) []bool {
	msg := d.client.send("/live/device/get/parameters/is_quantized", trackID, deviceID).Wait()
	quantized := make([]bool, 0)
	// Skip first two arguments which are track ID and device ID
	for i := 2; i < len(msg.Arguments); i++ {
		if val, ok := msg.Arguments[i].(int32); ok {
			quantized = append(quantized, val != 0)
		}
	}
	return quantized
}

// GetParameterValue returns the value of a specific parameter
func (d *DeviceAPI) GetParameterValue(trackID, deviceID, parameterID int32) float32 {
	msg := d.client.send("/live/device/get/parameter/value", trackID, deviceID, parameterID).Wait()
	if len(msg.Arguments) >= 4 {
		if val, ok := msg.Arguments[3].(float32); ok {
			return val
		}
	}
	return 0
}

// GetParameterValueString returns the value of a specific parameter as a string
func (d *DeviceAPI) GetParameterValueString(trackID, deviceID, parameterID int32) string {
	msg := d.client.send("/live/device/get/parameter/value_string", trackID, deviceID, parameterID).Wait()
	if len(msg.Arguments) >= 4 {
		if val, ok := msg.Arguments[3].(string); ok {
			return val
		}
	}
	return ""
}

// --- Property Setters ---

// SetParametersValue sets multiple parameter values
func (d *DeviceAPI) SetParametersValue(trackID, deviceID int32, values ...float32) {
	params := []any{trackID, deviceID}
	for _, val := range values {
		params = append(params, val)
	}
	d.client.send("/live/device/set/parameters/value", params...)
}

// SetParameterValue sets a specific parameter value
func (d *DeviceAPI) SetParameterValue(trackID, deviceID, parameterID int32, value float32) {
	d.client.send("/live/device/set/parameter/value", trackID, deviceID, parameterID, value)
}
