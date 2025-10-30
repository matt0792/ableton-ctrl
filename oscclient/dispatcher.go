package oscclient

import "github.com/hypebeast/go-osc/osc"

type DispatcherOption func(*osc.StandardDispatcher)

func WithHandler(addr string, handler func(msg *osc.Message)) DispatcherOption {
	return func(d *osc.StandardDispatcher) {
		d.AddMsgHandler(addr, handler)
	}
}
