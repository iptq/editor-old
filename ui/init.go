package ui

import "github.com/veandco/go-sdl2/sdl"

func Init() (err error) {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return
	}
}

func Close() {
	sdl.Quit()
}
