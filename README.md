# Ableton-ctrl

A Go client library for controlling Ableton Live via OSC using [AbletonOSC](https://github.com/ideoforms/AbletonOSC).

## Prerequisites

- Ableton Live 11 or above
- [AbletonOSC](https://github.com/ideoforms/AbletonOSC) installed and running

## Installation

```bash
go get github.com/matt0792/ableton-ctrl/als
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/matt0792/ableton-ctrl/als"
    "github.com/matt0792/ableton-ctrl/internal/oscclient"
)

func main() {
    // Create client (default AbletonOSC ports: send 11000, listen 11001)
    client := als.NewClient(oscclient.ClientOpts{
        SendAddr:   11000,
        ListenAddr: 11001,
    })

    // Start the client
    client.Run()
    defer client.Close()

    // Test connection
    result := client.Application.Test()
    fmt.Println("Connection:", result) // Prints "ok"

    // Get version
    major, minor := client.Application.GetVersion()
    fmt.Printf("Ableton Live %d.%d\n", major, minor)

    // Control playback
    client.Song.StartPlaying()
    tempo := client.Song.GetTempo()
    fmt.Printf("Current tempo: %.2f BPM\n", tempo)
}
```

## API Structure

The client is organized into the following API namespaces:

### Application API
```go
client.Application.Test()                    
client.Application.GetVersion()              
client.Application.SetLogLevel("debug")      
```

### Song API
```go
// Playback control
client.Song.StartPlaying()
client.Song.StopPlaying()
client.Song.ContinuePlaying()

// Transport
client.Song.GetTempo()
client.Song.SetTempo(120.0)
client.Song.GetCurrentSongTime()
client.Song.JumpBy(4.0)

// Track/Scene management
client.Song.GetNumTracks()
client.Song.GetTrackNames()
client.Song.CreateMIDITrack(0)
client.Song.CreateAudioTrack(0)
client.Song.CreateScene(0)

// Recording
client.Song.TriggerSessionRecord()
client.Song.CaptureMIDI()
client.Song.SetMetronome(true)

// Arrangement
client.Song.GetLoop()
client.Song.SetLoop(true)
client.Song.GetLoopStart()
client.Song.SetLoopStart(0.0)
```

### Track API
```go
// Track properties
client.Track.GetName(0)
client.Track.SetName(0, "Lead Synth")
client.Track.GetVolume(0)
client.Track.SetVolume(0, 0.85)
client.Track.GetPanning(0)
client.Track.SetPanning(0, 0.0)

// Track state
client.Track.SetMute(0, true)
client.Track.SetSolo(0, true)
client.Track.SetArm(0, true)

// Sends
client.Track.GetSend(0, 0)
client.Track.SetSend(0, 0, 0.5)

// Clips
client.Track.GetClipsName(0)
client.Track.GetClipsLength(0)
client.Track.StopAllClips(0)

// Devices
client.Track.GetNumDevices(0)
client.Track.GetDevicesName(0)
```

### Clip API
```go
// Clip control
client.Clip.Fire(0, 0)          // Fire clip at track 0, slot 0
client.Clip.Stop(0, 0)

// Clip properties
client.Clip.GetName(0, 0)
client.Clip.SetName(0, 0, "Bass Line")
client.Clip.GetLength(0, 0)
client.Clip.GetIsPlaying(0, 0)

// Loop settings
client.Clip.GetLoopStart(0, 0)
client.Clip.SetLoopStart(0, 0, 0.0)
client.Clip.GetLoopEnd(0, 0)
client.Clip.SetLoopEnd(0, 0, 4.0)

// MIDI notes
notes := client.Clip.GetNotes(0, 0)
client.Clip.AddNotes(0, 0, als.Note{
    Pitch: 60,        // C4
    StartTime: 0.0,
    Duration: 1.0,
    Velocity: 100,
    Mute: false,
})
client.Clip.RemoveNotes(0, 0, 60, 12, 0.0, 4.0)
```

### Scene API
```go
client.Scene.Fire(0)
client.Scene.GetName(0)
client.Scene.SetName(0, "Intro")
client.Scene.GetTempo(0)
client.Scene.SetTempo(0, 128.0)
```

### Device API
```go
// Device info
client.Device.GetName(0, 0)
client.Device.GetType(0, 0)
client.Device.GetClassName(0, 0)

// Parameters
client.Device.GetNumParameters(0, 0)
client.Device.GetParametersName(0, 0)
client.Device.GetParametersValue(0, 0)
client.Device.GetParameterValue(0, 0, 0)
client.Device.SetParameterValue(0, 0, 0, 0.75)
```

### View API
```go
// Selection
client.View.GetSelectedTrack()
client.View.SetSelectedTrack(2)
client.View.GetSelectedScene()
client.View.SetSelectedScene(1)
trackIdx, sceneIdx := client.View.GetSelectedClip()
client.View.SetSelectedClip(0, 0)
```

### ClipSlot API
```go
client.ClipSlot.Fire(0, 0)
client.ClipSlot.CreateClip(0, 0, 4.0)  // Create 4-bar clip
client.ClipSlot.DeleteClip(0, 0)
client.ClipSlot.GetHasClip(0, 0)
client.ClipSlot.DuplicateClipTo(0, 0, 1, 0)  // Duplicate to another slot
```

## Listening for Changes

You can listen for property changes by setting up handlers:

```go
import "github.com/hypebeast/go-osc/osc"

client := als.NewClient(oscclient.ClientOpts{
    SendAddr:   11000,
    ListenAddr: 11001,
    Handlers: []oscclient.DispatcherOption{
        oscclient.WithHandler("/live/song/get/tempo", func(msg *osc.Message) {
            if tempo, ok := msg.Arguments[0].(float32); ok {
                fmt.Printf("Tempo: %.2f BPM\n", tempo)
            }
        }),
    },
})

// Start listening
client.Song.StartListenProperty("tempo")

// Stop listening
client.Song.StopListenProperty("tempo")
```

## Working with MIDI Notes

```go
// Get notes from a clip
notes := client.Clip.GetNotes(trackID, clipID)
for _, note := range notes {
    fmt.Printf("Note: pitch=%d, start=%.2f, duration=%.2f, velocity=%d\n",
        note.Pitch, note.StartTime, note.Duration, note.Velocity)
}

// Add a chord
chord := []als.Note{
    {Pitch: 60, StartTime: 0.0, Duration: 1.0, Velocity: 100},  // C
    {Pitch: 64, StartTime: 0.0, Duration: 1.0, Velocity: 100},  // E
    {Pitch: 67, StartTime: 0.0, Duration: 1.0, Velocity: 100},  // G
}
client.Clip.AddNotes(trackID, clipID, chord...)

// Get notes in a specific range
// startPitch, pitchSpan, startTime, timeSpan
notes = client.Clip.GetNotes(trackID, clipID, 60, 12, 0.0, 4.0)
```

## Error Handling

The client uses a fire-and-forget model for commands that don't return values. For queries, the client waits for a response with a timeout. If a response isn't received, default values are returned (0, empty string, false, etc.).

## Thread Safety

The underlying OSC client handles concurrent sends safely. However, you should manage your own synchronization if you're accessing the client from multiple goroutines.
