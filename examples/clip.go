package examples

import (
	"fmt"

	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/oscclient"
)

func Basic() {
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
