package main

import (
	"editor"
	"fmt"
	"log"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	ed, err := editor.NewEditor()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(ed)

	// start the editor
	ed.Start()
}

func main() {
	pixelgl.Run(run)
}
