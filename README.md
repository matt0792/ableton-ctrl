## Packages 

- als: A client for Ableton Live
- alsex: Extension methods for als 
- oscclient: Wrapper around [go-osc](github.com/hypebeast/go-osc)

## Prerequisites 

- Ableton Live 11 or above
- [AbletonOSC](https://github.com/ideoforms/AbletonOSC) installed and running

## Notes

The client uses a fire-and-forget model for commands that don't return values. For queries, the client waits for a response with a timeout. If a response isn't received, default values are returned (0, empty string, false, etc.).

The underlying OSC client handles concurrent sends safely. However, you should manage your own synchronization if you're accessing the client from multiple goroutines.