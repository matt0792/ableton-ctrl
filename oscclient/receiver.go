package oscclient

import (
	"time"

	"github.com/hypebeast/go-osc/osc"
)

type Receiver struct {
	queue   *Queue
	timeout time.Duration
}

func NewReceiver() *Receiver {
	return &Receiver{
		queue:   NewQueue(),
		timeout: 1000 * time.Millisecond,
	}
}

// Expect registers a pending request and returns the response channel
func (r *Receiver) Expect(addr string) chan *osc.Message {
	return r.queue.Register(addr)
}

// WaitFor waits for response on the given channel with timeout
func (r *Receiver) WaitFor(ch chan *osc.Message, addr string) *osc.Message {
	timeout := time.After(r.timeout)

	select {
	case msg, ok := <-ch:
		if !ok || msg == nil {
			return &osc.Message{}
		}
		return msg
	case <-timeout:
		r.queue.Cancel(addr, ch)
		return &osc.Message{}
	}
}

// Callback is non blocking, and calls a func when result arrives
func (r *Receiver) Callback(addr string, callback func(*osc.Message)) {
	ch := r.Expect(addr)
	go func() {
		result := r.WaitFor(ch, addr)
		callback(result)
	}()
}

// WaitChan returns a channel that will receive the result
func (r *Receiver) WaitChan(addr string) <-chan *osc.Message {
	result := make(chan *osc.Message, 1)
	ch := r.Expect(addr)

	go func() {
		msg := r.WaitFor(ch, addr)
		result <- msg
		close(result)
	}()

	return result
}

// Populate delivers an incoming message to the appropriate waiting channel
func (r *Receiver) Populate(msg *osc.Message) {
	r.queue.Deliver(msg)
}
