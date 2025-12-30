package project

import (
	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/oscclient"
)

type Project struct {
	api *als.Client
}

func NewProject() *Project {
	return &Project{
		api: als.NewClient(oscclient.ClientOpts{
			SendAddr:   11000,
			ListenAddr: 11001,
		}),
	}
}
