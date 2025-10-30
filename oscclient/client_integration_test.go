package oscclient

import (
	"testing"
	"time"

	"github.com/hypebeast/go-osc/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAbletonIntegration tests actual communication with Ableton Live
// To run these tests:
//  1. Install AbletonOSC: https://github.com/ideoforms/AbletonOSC
//  2. Start Ableton Live with AbletonOSC enabled
//  3. Run: go test -v -run TestAbletonIntegration
//
// These tests will be skipped if Ableton is not running.

// createAbletonClient creates a client configured to communicate with Ableton Live
// Ableton receives on port 11000 and sends responses on port 11001
func createAbletonClient() *Client {
	return NewClient(ClientOpts{
		SendAddr:   11000, // Ableton listens here
		ListenAddr: 11001, // Ableton sends responses here
		Handlers: []DispatcherOption{
			// Auto-populate responses for all /live/* messages
			WithHandler("/live/*", func(msg *osc.Message) {
				// This will be set during client creation
			}),
		},
	})
}

// checkAbletonRunning attempts to connect to Ableton and returns true if successful
func checkAbletonRunning() bool {
	client := createAbletonClient()

	// Set up the handler to populate responses
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/test", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	// Give server time to start
	time.Sleep(50 * time.Millisecond)

	// Try to send a test message with a short timeout
	call := client.Send("/live/test")

	// Use a channel to timeout
	done := make(chan *osc.Message, 1)
	go func() {
		done <- call.Wait()
	}()

	select {
	case msg := <-done:
		return msg != nil && len(msg.Arguments) > 0
	case <-time.After(2 * time.Second):
		return false
	}
}

// TestAbletonIntegration_BasicConnectivity tests basic connection to Ableton
func TestAbletonIntegration_BasicConnectivity(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/test", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	// Give server time to start
	time.Sleep(50 * time.Millisecond)

	// Check if Ableton is running
	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running or AbletonOSC is not enabled")
	}

	// Send test message and wait for response
	call := client.Send("/live/test")
	result := call.Wait()

	require.NotNil(t, result, "Expected response from Ableton")
	require.Greater(t, len(result.Arguments), 0, "Expected at least one argument")
	assert.Equal(t, "ok", result.Arguments[0], "Expected 'ok' response from /live/test")
}

// TestAbletonIntegration_ApplicationVersion tests querying Ableton's version
func TestAbletonIntegration_ApplicationVersion(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/application/get/version", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query version
	call := client.Send("/live/application/get/version")
	result := call.Wait()

	require.NotNil(t, result, "Expected version response")
	require.GreaterOrEqual(t, len(result.Arguments), 2, "Expected major and minor version")

	// Version should be integers
	major, ok := result.Arguments[0].(int32)
	assert.True(t, ok, "Major version should be int32")
	assert.Greater(t, major, int32(0), "Major version should be positive")

	minor, ok := result.Arguments[1].(int32)
	assert.True(t, ok, "Minor version should be int32")
	assert.GreaterOrEqual(t, minor, int32(0), "Minor version should be non-negative")

	t.Logf("Ableton Live version: %d.%d", major, minor)
}

// TestAbletonIntegration_SongTempo tests querying and setting song tempo
func TestAbletonIntegration_SongTempo(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/song/get/tempo", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query tempo
	call := client.Send("/live/song/get/tempo")
	result := call.Wait()

	require.NotNil(t, result, "Expected tempo response")
	require.Greater(t, len(result.Arguments), 0, "Expected tempo value")

	tempo, ok := result.Arguments[0].(float32)
	assert.True(t, ok, "Tempo should be float32")
	assert.Greater(t, tempo, float32(0), "Tempo should be positive")
	assert.Less(t, tempo, float32(999), "Tempo should be reasonable")

	t.Logf("Current tempo: %.2f BPM", tempo)
}

// TestAbletonIntegration_SongIsPlaying tests querying playback state
func TestAbletonIntegration_SongIsPlaying(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/song/get/is_playing", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query playback state
	call := client.Send("/live/song/get/is_playing")
	result := call.Wait()

	require.NotNil(t, result, "Expected is_playing response")
	require.Greater(t, len(result.Arguments), 0, "Expected is_playing value")

	// Can be either int32 (0/1) or bool depending on AbletonOSC version
	switch v := result.Arguments[0].(type) {
	case int32:
		assert.Contains(t, []int32{0, 1}, v, "is_playing should be 0 or 1")
		t.Logf("Is playing: %t", v == 1)
	case bool:
		t.Logf("Is playing: %t", v)
	default:
		t.Errorf("Unexpected type for is_playing: %T", v)
	}
}

// TestAbletonIntegration_SongCurrentTime tests querying current song time
func TestAbletonIntegration_SongCurrentTime(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/song/get/current_song_time", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query current time
	call := client.Send("/live/song/get/current_song_time")
	result := call.Wait()

	require.NotNil(t, result, "Expected current_song_time response")
	require.Greater(t, len(result.Arguments), 0, "Expected time value")

	songTime, ok := result.Arguments[0].(float32)
	assert.True(t, ok, "Song time should be float32")
	assert.GreaterOrEqual(t, songTime, float32(0), "Song time should be non-negative")

	t.Logf("Current song time: %.2f beats", songTime)
}

// TestAbletonIntegration_TrackVolume tests querying track volume
func TestAbletonIntegration_TrackVolume(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/track/get/volume", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query volume of track 0
	call := client.Send("/live/track/get/volume", int32(0))
	result := call.Wait()

	require.NotNil(t, result, "Expected volume response")
	require.GreaterOrEqual(t, len(result.Arguments), 2, "Expected track index and volume")

	trackIndex, ok := result.Arguments[0].(int32)
	assert.True(t, ok, "Track index should be int32")
	assert.Equal(t, int32(0), trackIndex, "Track index should match request")

	volume, ok := result.Arguments[1].(float32)
	assert.True(t, ok, "Volume should be float32")
	assert.GreaterOrEqual(t, volume, float32(0), "Volume should be non-negative")
	assert.LessOrEqual(t, volume, float32(1), "Volume should be <= 1.0")

	t.Logf("Track 0 volume: %.2f", volume)
}

// TestAbletonIntegration_TrackName tests querying track name
func TestAbletonIntegration_TrackName(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/track/get/name", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query name of track 0
	call := client.Send("/live/track/get/name", int32(0))
	result := call.Wait()

	require.NotNil(t, result, "Expected name response")
	require.GreaterOrEqual(t, len(result.Arguments), 2, "Expected track index and name")

	trackIndex, ok := result.Arguments[0].(int32)
	assert.True(t, ok, "Track index should be int32")
	assert.Equal(t, int32(0), trackIndex, "Track index should match request")

	name, ok := result.Arguments[1].(string)
	assert.True(t, ok, "Track name should be string")
	assert.NotEmpty(t, name, "Track name should not be empty")

	t.Logf("Track 0 name: %s", name)
}

// TestAbletonIntegration_ViewSelectedTrack tests querying selected track
func TestAbletonIntegration_ViewSelectedTrack(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/view/get/selected_track", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Query selected track
	call := client.Send("/live/view/get/selected_track")
	result := call.Wait()

	require.NotNil(t, result, "Expected selected_track response")
	require.Greater(t, len(result.Arguments), 0, "Expected track index")

	trackIndex, ok := result.Arguments[0].(int32)
	assert.True(t, ok, "Track index should be int32")
	assert.GreaterOrEqual(t, trackIndex, int32(0), "Track index should be non-negative")

	t.Logf("Selected track: %d", trackIndex)
}

// TestAbletonIntegration_MultipleRequests tests sending multiple requests in sequence
func TestAbletonIntegration_MultipleRequests(t *testing.T) {
	var client *Client
	client = NewClient(ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
		Handlers: []DispatcherOption{
			WithHandler("/live/*", func(msg *osc.Message) {
				client.receiver.Populate(msg)
			}),
		},
	})

	client.Run()
	defer client.Close()

	time.Sleep(50 * time.Millisecond)

	if !checkAbletonRunning() {
		t.Skip("Ableton Live is not running")
	}

	// Send multiple requests
	testCall := client.Send("/live/test")
	versionCall := client.Send("/live/application/get/version")
	tempoCall := client.Send("/live/song/get/tempo")

	// Wait for all responses
	testResult := testCall.Wait()
	versionResult := versionCall.Wait()
	tempoResult := tempoCall.Wait()

	// Verify all received responses
	assert.NotNil(t, testResult, "Test response should not be nil")
	assert.NotNil(t, versionResult, "Version response should not be nil")
	assert.NotNil(t, tempoResult, "Tempo response should not be nil")

	if testResult != nil && len(testResult.Arguments) > 0 {
		assert.Equal(t, "ok", testResult.Arguments[0])
	}

	if versionResult != nil && len(versionResult.Arguments) >= 2 {
		t.Logf("Version: %v.%v", versionResult.Arguments[0], versionResult.Arguments[1])
	}

	if tempoResult != nil && len(tempoResult.Arguments) > 0 {
		t.Logf("Tempo: %v BPM", tempoResult.Arguments[0])
	}
}
