package main

import (
	"editor"
	"log"
)

func main() {
	ed, err := editor.NewEditor()
	if err != nil {
		log.Fatal(err)
	}

	err = ed.Open("/home/michael/osu/Songs/firststorm/Will Stetson - First Storm (Japanese Cover) (deadcode) [thanks for singing Brother].osu")
	if err != nil {
		log.Fatal(err)
	}

	// start the editor
	ed.Start()
}
