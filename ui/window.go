package ui

import (
	"github.com/faiface/pixel/pixelgl"
)

type Window struct {
	*pixelgl.Window
}

func NewWindow(config pixelgl.WindowConfig) (*Window, error) {
	inner, err := pixelgl.NewWindow(config)
	if err != nil {
		return nil, err
	}

	window := &Window{
		Window: inner,
	}
	return window, nil
}
