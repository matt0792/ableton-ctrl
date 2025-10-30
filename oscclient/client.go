package oscclient

import (
	"fmt"
	"log"

	"github.com/hypebeast/go-osc/osc"
)

type Client struct {
	engine     *osc.Client
	server     *osc.Server
	receiver   *Receiver
	isHandling bool
}

type ClientOpts struct {
	Handlers   []DispatcherOption
	SendAddr   int
	ListenAddr int
}

func NewClient(opts ClientOpts) *Client {
	d := osc.NewStandardDispatcher()

	c := &Client{
		engine: osc.NewClient("localhost", opts.SendAddr),
		server: &osc.Server{
			Addr:       fmt.Sprintf("127.0.0.1:%d", opts.ListenAddr),
			Dispatcher: d,
		},
		receiver:   NewReceiver(),
		isHandling: false,
	}

	// Add default handler to route all messages to receiver
	d.AddMsgHandler("*", func(msg *osc.Message) {
		c.receiver.Populate(msg)
	})

	// Apply custom handlers (these will override the default for specific addresses)
	for _, opt := range opts.Handlers {
		opt(d)
	}

	return c
}

type SendOption func(*sendConfig)

type sendConfig struct {
	wait bool
}

// WithWait makes the send wait for a response
func WithWait() SendOption {
	return func(c *sendConfig) {
		c.wait = true
	}
}

func (c *Client) Send(addr string, params ...any) *Call {
	return c.sendWith(addr, params...)
}

func (c *Client) sendWith(addr string, params ...any) *Call {
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

// Run starts the server, and listens for messages
func (c *Client) Run() {
	go c.server.ListenAndServe()
	c.isHandling = true
}

// Close stops the server
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
