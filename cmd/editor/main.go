package main

import (
	"editor"
	"log"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	ed, err := editor.NewEditor()
	if err != nil {
		log.Panic(err)
	}

	// start the editor
	ed.Start()
}

func main() {
	pixelgl.Run(run)
}
