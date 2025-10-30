package als

import "github.com/matt0792/ableton-ctrl/oscclient"

// Client provides a high-level interface to the Ableton Live OSC API
type Client struct {
	osc         *oscclient.Client
	Application *ApplicationAPI
	Song        *SongAPI
	Track       *TrackAPI
	Clip        *ClipAPI
	Scene       *SceneAPI
	Device      *DeviceAPI
	View        *ViewAPI
	ClipSlot    *ClipSlotAPI
}

// NewClient creates a new Ableton Live OSC client
// By default, AbletonOSC listens on port 11000 and sends replies on port 11001
func NewClient(opts oscclient.ClientOpts) *Client {
	oscClient := oscclient.NewClient(opts)

	c := &Client{
		osc: oscClient,
	}

	// Initialize API namespaces
	c.Application = &ApplicationAPI{client: c}
	c.Song = &SongAPI{client: c}
	c.Track = &TrackAPI{client: c}
	c.Clip = &ClipAPI{client: c}
	c.Scene = &SceneAPI{client: c}
	c.Device = &DeviceAPI{client: c}
	c.View = &ViewAPI{client: c}
	c.ClipSlot = &ClipSlotAPI{client: c}

	return c
}

// Run starts the OSC client server to listen for incoming messages
func (c *Client) Run() {
	c.osc.Run()
}

// Close stops the OSC client server
func (c *Client) Close() {
	c.osc.Close()
}

// send is a helper method for sending OSC messages
func (c *Client) send(addr string, params ...any) *oscclient.Call {
	return c.osc.Send(addr, params...)
}
