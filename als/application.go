package als

// ApplicationAPI provides methods for interacting with Ableton
type ApplicationAPI struct {
	client *Client
}

func (a *ApplicationAPI) Test() string {
	msg := a.client.send("/live/test").Wait()
	if len(msg.Arguments) > 0 {
		if result, ok := msg.Arguments[0].(string); ok {
			return result
		}
	}
	return ""
}

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

// Reload initiates a live reload of the AbletonOSC server code.
func (a *ApplicationAPI) Reload() {
	a.client.send("/live/api/reload")
}

func (a *ApplicationAPI) GetLogLevel() string {
	msg := a.client.send("/live/api/get/log_level").Wait()
	if len(msg.Arguments) > 0 {
		if level, ok := msg.Arguments[0].(string); ok {
			return level
		}
	}
	return ""
}

func (a *ApplicationAPI) SetLogLevel(level string) {
	a.client.send("/live/api/set/log_level", level)
}
