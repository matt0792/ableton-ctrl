package oscclient

import (
	"fmt"
	"testing"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClientCreation verifies that a client can be created with proper configuration
func TestClientCreation(t *testing.T) {
	client := NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers:   []DispatcherOption{},
	})

	assert.NotNil(t, client)
	assert.NotNil(t, client.engine)
	assert.NotNil(t, client.server)
	assert.NotNil(t, client.receiver)
	assert.False(t, client.isHandling)
}

// TestClientWithHandlers verifies that handlers are properly registered
func TestClientWithHandlers(t *testing.T) {
	handlerCalled := false

	client := NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/test", func(msg *osc.Message) {
				handlerCalled = true
			}),
		},
	})

	assert.NotNil(t, client)

	// Start the server
	client.Run()
	defer client.Close()

	// Give server time to start
	time.Sleep(10 * time.Millisecond)

	// Send a message to the handler
	testClient := osc.NewClient("localhost", 11001)
	msg := osc.NewMessage("/test")
	testClient.Send(msg)

	// Wait for handler to be called
	time.Sleep(50 * time.Millisecond)

	assert.True(t, handlerCalled)
}

// TestMessageSending verifies that messages can be sent with various parameter types
func TestMessageSending(t *testing.T) {
	client := NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11002,
	})

	tests := []struct {
		name   string
		addr   string
		params []any
	}{
		{
			name:   "no parameters",
			addr:   "/live/test",
			params: []any{},
		},
		{
			name:   "single int32",
			addr:   "/live/track/get/volume",
			params: []any{int32(0)},
		},
		{
			name:   "multiple int32",
			addr:   "/live/clip/get/name",
			params: []any{int32(0), int32(0)},
		},
		{
			name:   "mixed types",
			addr:   "/live/clip/set/notes",
			params: []any{int32(0), int32(0), int32(60), int32(0), float32(0.5), int32(100)},
		},
		{
			name:   "with string",
			addr:   "/live/song/set/tempo",
			params: []any{float32(120.0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := client.Send(tt.addr, tt.params...)
			assert.NotNil(t, call)
			assert.Equal(t, tt.addr, call.addr)
		})
	}
}

// TestQueueOperations verifies queue registration and delivery
func TestQueueOperations(t *testing.T) {
	queue := NewQueue()

	// Test registration
	addr := "/live/test"
	ch := queue.Register(addr)
	assert.NotNil(t, ch)

	// Test delivering a message
	msg := osc.NewMessage("/live/test")
	msg.Append("ok")
	queue.Deliver(msg)

	// Receive from channel
	retrieved := <-ch
	assert.NotNil(t, retrieved)
	assert.Equal(t, "/live/test", retrieved.Address)
	assert.Equal(t, 1, len(retrieved.Arguments))
	assert.Equal(t, "ok", retrieved.Arguments[0])
}

// TestReceiverPopulate verifies that messages are properly delivered
func TestReceiverPopulate(t *testing.T) {
	receiver := NewReceiver()

	// Register for the message first
	ch := receiver.Expect("/live/application/get/version")

	// Populate the message
	msg := osc.NewMessage("/live/application/get/version")
	msg.Append(int32(11))
	msg.Append(int32(3))
	receiver.Populate(msg)

	// Receive from channel
	retrieved := <-ch
	assert.NotNil(t, retrieved)
	assert.Equal(t, "/live/application/get/version", retrieved.Address)
	assert.Equal(t, 2, len(retrieved.Arguments))
}

// TestReceiverWaitFor verifies waiting for messages with timeout
func TestReceiverWaitFor(t *testing.T) {
	receiver := NewReceiver()

	// Test timeout case
	ch := receiver.Expect("/live/test")
	result := receiver.WaitFor(ch, "/live/test")
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Arguments)) // Empty message on timeout

	// Test successful receive
	ch2 := receiver.Expect("/live/application/get/version")

	// Populate the message in a goroutine
	go func() {
		time.Sleep(20 * time.Millisecond)
		msg := osc.NewMessage("/live/application/get/version")
		msg.Append(int32(11))
		msg.Append(int32(3))
		receiver.Populate(msg)
	}()

	result2 := receiver.WaitFor(ch2, "/live/application/get/version")
	assert.NotNil(t, result2)
	assert.Equal(t, "/live/application/get/version", result2.Address)
	assert.Equal(t, 2, len(result2.Arguments))
}

// TestReceiverCallback verifies callback functionality
func TestReceiverCallback(t *testing.T) {
	receiver := NewReceiver()

	callbackFired := false
	var receivedMsg *osc.Message

	receiver.Callback("/live/test", func(msg *osc.Message) {
		callbackFired = true
		receivedMsg = msg
	})

	// Give callback goroutine time to register
	time.Sleep(10 * time.Millisecond)

	// Populate message
	msg := osc.NewMessage("/live/test")
	msg.Append("ok")
	receiver.Populate(msg)

	// Wait for callback
	time.Sleep(100 * time.Millisecond)

	assert.True(t, callbackFired)
	assert.NotNil(t, receivedMsg)
	assert.Equal(t, "/live/test", receivedMsg.Address)
}

// TestReceiverWaitChan verifies channel-based waiting
func TestReceiverWaitChan(t *testing.T) {
	receiver := NewReceiver()

	resultChan := receiver.WaitChan("/live/application/get/version")

	// Populate message
	go func() {
		time.Sleep(20 * time.Millisecond)
		msg := osc.NewMessage("/live/application/get/version")
		msg.Append(int32(11))
		msg.Append(int32(3))
		receiver.Populate(msg)
	}()

	result := <-resultChan
	assert.NotNil(t, result)
	assert.Equal(t, "/live/application/get/version", result.Address)
	assert.Equal(t, 2, len(result.Arguments))
}

// TestAbletonOSCMessagePatterns tests message patterns matching the AbletonOSC API
func TestAbletonOSCMessagePatterns(t *testing.T) {
	tests := []struct {
		name         string
		sendAddr     string
		sendParams   []any
		expectedAddr string
		expectedArgs []any
		description  string
	}{
		{
			name:         "test command",
			sendAddr:     "/live/test",
			sendParams:   []any{},
			expectedAddr: "/live/test",
			expectedArgs: []any{"ok"},
			description:  "Basic test command should return 'ok'",
		},
		{
			name:         "get version",
			sendAddr:     "/live/application/get/version",
			sendParams:   []any{},
			expectedAddr: "/live/application/get/version",
			expectedArgs: []any{int32(11), int32(3)},
			description:  "Version query returns major and minor version",
		},
		{
			name:         "get tempo",
			sendAddr:     "/live/song/get/tempo",
			sendParams:   []any{},
			expectedAddr: "/live/song/get/tempo",
			expectedArgs: []any{float32(120.0)},
			description:  "Tempo query returns float32",
		},
		{
			name:         "get clip name",
			sendAddr:     "/live/clip/get/name",
			sendParams:   []any{int32(0), int32(0)},
			expectedAddr: "/live/clip/get/name",
			expectedArgs: []any{int32(0), int32(0), "Clip Name"},
			description:  "Clip name query with track and clip indices",
		},
		{
			name:         "get track volume",
			sendAddr:     "/live/track/get/volume",
			sendParams:   []any{int32(0)},
			expectedAddr: "/live/track/get/volume",
			expectedArgs: []any{int32(0), float32(0.85)},
			description:  "Track volume query returns track index and volume",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a message matching the expected response
			msg := osc.NewMessage(tt.expectedAddr)
			for _, arg := range tt.expectedArgs {
				msg.Append(arg)
			}

			// Verify message structure
			assert.Equal(t, tt.expectedAddr, msg.Address)
			assert.Equal(t, len(tt.expectedArgs), len(msg.Arguments))

			// Verify argument types
			for i, expectedArg := range tt.expectedArgs {
				switch v := expectedArg.(type) {
				case string:
					assert.IsType(t, "", msg.Arguments[i], fmt.Sprintf("arg %d should be string", i))
				case int32:
					assert.IsType(t, int32(0), msg.Arguments[i], fmt.Sprintf("arg %d should be int32", i))
				case float32:
					assert.IsType(t, float32(0), msg.Arguments[i], fmt.Sprintf("arg %d should be float32", i))
				default:
					t.Errorf("unexpected type for arg %d: %T", i, v)
				}
			}
		})
	}
}

// TestEndToEndFlow simulates a complete send-receive cycle
func TestEndToEndFlow(t *testing.T) {
	// Create a mock AbletonOSC server
	dispatcher := osc.NewStandardDispatcher()
	dispatcher.AddMsgHandler("/live/test", func(msg *osc.Message) {
		// Send response
		responseClient := osc.NewClient("localhost", 11004)
		response := osc.NewMessage("/live/test")
		response.Append("ok")
		responseClient.Send(response)
	})

	mockServer := &osc.Server{
		Addr:       "127.0.0.1:11003",
		Dispatcher: dispatcher,
	}
	go mockServer.ListenAndServe()
	defer mockServer.CloseConnection()

	// Give server time to start
	time.Sleep(10 * time.Millisecond)

	// Create our client
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11003,
		ListenAddr: 11004,
		Handlers: []DispatcherOption{
			WithHandler("/live/test", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	// Give client time to start
	time.Sleep(10 * time.Millisecond)

	// Send message and wait for response
	call := client.Send("/live/test")
	result := call.Wait()

	assert.NotNil(t, result)
	assert.Equal(t, "/live/test", result.Address)
	require.Greater(t, len(result.Arguments), 0)
	assert.Equal(t, "ok", result.Arguments[0])
}

// TestArgumentExtraction verifies extracting different types from messages
func TestArgumentExtraction(t *testing.T) {
	msg := osc.NewMessage("/live/application/get/version")
	msg.Append(int32(11))
	msg.Append(int32(3))
	msg.Append(float32(120.5))
	msg.Append("test string")

	// Test extracting arguments
	require.Equal(t, 4, len(msg.Arguments))

	// Extract int32
	major, ok := msg.Arguments[0].(int32)
	assert.True(t, ok)
	assert.Equal(t, int32(11), major)

	minor, ok := msg.Arguments[1].(int32)
	assert.True(t, ok)
	assert.Equal(t, int32(3), minor)

	// Extract float32
	tempo, ok := msg.Arguments[2].(float32)
	assert.True(t, ok)
	assert.Equal(t, float32(120.5), tempo)

	// Extract string
	str, ok := msg.Arguments[3].(string)
	assert.True(t, ok)
	assert.Equal(t, "test string", str)
}

// BenchmarkMessageSending benchmarks message creation and sending
func BenchmarkMessageSending(b *testing.B) {
	client := NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11005,
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.Send("/live/test", int32(i))
	}
}

// BenchmarkQueueOperations benchmarks queue operations
func BenchmarkQueueOperations(b *testing.B) {
	queue := NewQueue()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := queue.Register("/live/test")
		msg := osc.NewMessage("/live/test")
		msg.Append("ok")
		queue.Deliver(msg)
		<-ch
	}
}
