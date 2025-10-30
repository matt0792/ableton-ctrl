package oscclient

import (
	"errors"
	"sync"

	"github.com/hypebeast/go-osc/osc"
)

var (
	ErrDuplicateId = errors.New("duplicate id")
	ErrNotFound    = errors.New("not found")
)

type Queue struct {
	// Map from address -> list of pending response channels
	pending map[string][]chan *osc.Message
	mu      sync.RWMutex
}

// NewQueue creates a new queue for managing request/response correlation
func NewQueue() *Queue {
	return &Queue{
		pending: make(map[string][]chan *osc.Message),
	}
}

// Register creates and registers a response channel for an address
// Returns the channel that will receive the response
func (q *Queue) Register(addr string) chan *osc.Message {
	ch := make(chan *osc.Message, 1)
	q.mu.Lock()
	defer q.mu.Unlock()
	q.pending[addr] = append(q.pending[addr], ch)
	return ch
}

// Deliver sends a response to the oldest waiting channel for the address
// Uses FIFO ordering to match responses with requests
func (q *Queue) Deliver(msg *osc.Message) {
	q.mu.Lock()
	defer q.mu.Unlock()

	channels := q.pending[msg.Address]
	if len(channels) == 0 {
		// No one waiting for this response
		return
	}

	// Send to first (oldest) waiting channel
	channels[0] <- msg
	close(channels[0])

	// Remove from pending queue
	q.pending[msg.Address] = channels[1:]
}

// Cancel removes a pending channel from the queue (for cleanup/timeout)
func (q *Queue) Cancel(addr string, ch chan *osc.Message) {
	q.mu.Lock()
	defer q.mu.Unlock()

	channels := q.pending[addr]
	for i, pending := range channels {
		if pending == ch {
			// Remove this channel from the list
			q.pending[addr] = append(channels[:i], channels[i+1:]...)
			close(ch)
			return
		}
	}
}
