package oscclient

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

type Client struct {
	engine      *osc.Client
	server      *osc.Server
	receiver    *Receiver
	isHandling  bool
	rateLimiter *rateLimiter
}

type ClientOpts struct {
	Handlers     []DispatcherOption
	SendAddr     int
	ListenAddr   int
	Timeout      time.Duration
	EnableLogger bool
	RateLimit    int
}

// implements token bucket rate limiting
type rateLimiter struct {
	enabled    bool
	tokens     int
	maxTokens  int
	refillRate time.Duration
	mu         sync.Mutex
	lastRefill time.Time
}

func newRateLimiter(requestsPerSec int) *rateLimiter {
	if requestsPerSec <= 0 {
		return &rateLimiter{enabled: false}
	}
	return &rateLimiter{
		enabled:    true,
		tokens:     requestsPerSec,
		maxTokens:  requestsPerSec,
		refillRate: time.Second / time.Duration(requestsPerSec),
		lastRefill: time.Now(),
	}
}

func (rl *rateLimiter) wait() {
	if !rl.enabled {
		return
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Refill tokens based on elapsed time
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	tokensToAdd := int(elapsed / rl.refillRate)

	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastRefill = now
	}

	// Wait if no tokens available
	for rl.tokens <= 0 {
		rl.mu.Unlock()
		time.Sleep(rl.refillRate)
		rl.mu.Lock()

		now = time.Now()
		elapsed = now.Sub(rl.lastRefill)
		tokensToAdd = int(elapsed / rl.refillRate)

		if tokensToAdd > 0 {
			rl.tokens += tokensToAdd
			if rl.tokens > rl.maxTokens {
				rl.tokens = rl.maxTokens
			}
			rl.lastRefill = now
		}
	}

	rl.tokens--
}

func NewClient(opts ClientOpts) *Client {
	d := osc.NewStandardDispatcher()

	c := &Client{
		engine: osc.NewClient("localhost", opts.SendAddr),
		server: &osc.Server{
			Addr:       fmt.Sprintf("127.0.0.1:%d", opts.ListenAddr),
			Dispatcher: d,
		},
		receiver:    NewReceiver(opts.Timeout, opts.EnableLogger),
		isHandling:  false,
		rateLimiter: newRateLimiter(opts.RateLimit),
	}

	// routes all messages to receiver
	d.AddMsgHandler("*", func(msg *osc.Message) {
		c.receiver.Populate(msg)
	})

	for _, opt := range opts.Handlers {
		opt(d)
	}

	return c
}

type SendOption func(*sendConfig)

type sendConfig struct {
	wait bool
}

func WithWait() SendOption {
	return func(c *sendConfig) {
		c.wait = true
	}
}

func (c *Client) Send(addr string, params ...any) *Call {
	return c.sendWith(addr, params...)
}

func (c *Client) sendWith(addr string, params ...any) *Call {
	c.rateLimiter.wait()

	msg := osc.NewMessage(addr)
	for _, v := range params {
		msg.Append(v)
	}

	call := &Call{
		receiver: c.receiver,
		addr:     addr,
	}

	c.engine.Send(msg)
	return call
}

type Call struct {
	receiver *Receiver
	addr     string
}

// Wait blocks until response is received
func (c *Call) Wait() *osc.Message {
	ch := c.receiver.Expect(c.addr)
	return c.receiver.WaitFor(ch, c.addr)
}

func (c *Client) Run() {
	go c.server.ListenAndServe()
	c.isHandling = true
}

func (c *Client) Close() {
	defer log.Println("server closed")

	if !c.isHandling {
		return
	}

	err := c.server.CloseConnection()
	if err != nil {
		log.Println(err)
	}
}
