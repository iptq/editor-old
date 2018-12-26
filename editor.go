package editor

import (
	"fmt"
	"image/color"
	"io/ioutil"

	"editor/osu"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	// the amount by which scrolling changes timestamp
	// it's negative because scrolling down moves forward
	SCROLL_CONSTANT float64 = -15.0
)

// Editor contains the state of the editor.
type Editor struct {
	// current audio timestamp in milliseconds
	timestamp float64
	playing   bool
	beatmap   *osu.Beatmap

	atlas  *text.Atlas
	window *pixelgl.Window
}

func NewEditor() (*Editor, error) {
	// text atlas
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	// initialize window
	wincfg := pixelgl.WindowConfig{
		Title:  "editor",
		Bounds: pixel.R(0, 0, 1366, 768),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		return nil, err
	}

	editor := &Editor{
		timestamp: 0.0,
		playing:   false,

		atlas:  atlas,
		window: window,
	}
	return editor, nil
}

func (editor *Editor) update() {
	// update audio position based on scroll
	editor.timestamp += editor.window.MouseScroll().Y * SCROLL_CONSTANT
}

func (editor *Editor) draw() {
	// draw audio timestamp
	timestamp := text.New(pixel.V(100, 500), editor.atlas)
	fmt.Fprintf(timestamp, "%f", editor.timestamp)
	timestamp.Draw(editor.window, pixel.IM)
}

func (editor *Editor) Open(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// convert to utf-8
	contents := string(data)
	beatmap, err := osu.DeserializeBeatmap(contents)
	if err != nil {
		return err
	}

	editor.beatmap = beatmap
	return nil
}

// Start runs the editor
func (editor *Editor) Start() {
	// the main game loop
	for !editor.window.Closed() {
		editor.window.Clear(color.Black)
		editor.update()
		editor.draw()
		editor.window.Update()
	}
}
