package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	*sdl.Window
}

func NewWindow() Window {
	return Window{
		Window: nil,
	}
}
