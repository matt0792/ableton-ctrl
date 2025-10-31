package examples

import (
	"fmt"
	"time"

	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/oscclient"
)

func Basic() {
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
	})

	client.Run()
	defer client.Close()
	time.Sleep(100 * time.Millisecond)

	result := client.Application.Test()
	fmt.Println(result)

}
