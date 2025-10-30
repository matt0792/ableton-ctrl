package als

// ApplicationAPI provides methods for interacting with Ableton Live's application-level API
type ApplicationAPI struct {
	client *Client
}

// Test sends a test command to Ableton Live and returns "ok" if successful
func (a *ApplicationAPI) Test() string {
	msg := a.client.send("/live/test").Wait()
	if len(msg.Arguments) > 0 {
		if result, ok := msg.Arguments[0].(string); ok {
			return result
		}
	}
	return ""
}

// GetVersion returns the major and minor version numbers of Ableton Live
func (a *ApplicationAPI) GetVersion() (major, minor int32) {
	msg := a.client.send("/live/application/get/version").Wait()
	if len(msg.Arguments) >= 2 {
		if maj, ok := msg.Arguments[0].(int32); ok {
			major = maj
		}
		if min, ok := msg.Arguments[1].(int32); ok {
			minor = min
		}
	}
	return
}

// Reload initiates a live reload of the AbletonOSC server code
func (a *ApplicationAPI) Reload() {
	a.client.send("/live/api/reload")
}

// GetLogLevel retrieves the current logging granularity level
func (a *ApplicationAPI) GetLogLevel() string {
	msg := a.client.send("/live/api/get/log_level").Wait()
	if len(msg.Arguments) > 0 {
		if level, ok := msg.Arguments[0].(string); ok {
			return level
		}
	}
	return ""
}

// SetLogLevel sets the log level (debug, info, warning, error, critical)
func (a *ApplicationAPI) SetLogLevel(level string) {
	a.client.send("/live/api/set/log_level", level)
}
