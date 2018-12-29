package main

import (
	"editor"
	"log"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	ui.Init()
	defer ui.Close()

	ed, err := editor.NewEditor()
	if err != nil {
		log.Fatal(err)
	}

	err = ed.Open("/home/michael/osu/Songs/firststorm/Will Stetson - First Storm (Japanese ver.) (deadcode) [thanks for singing Brother].osu")
	if err != nil {
		log.Fatal(err)
	}

	// start the editor
	ed.Start()
}

func main() {
	pixelgl.Run(run)
}
