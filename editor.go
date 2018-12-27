package editor

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"editor/osu"
	"editor/ui"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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
	timestamp int
	playing   bool
	audio     *AudioManager
	beatmap   *EditorBeatmap

	atlas  *text.Atlas
	window *ui.Window
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
	window, err := ui.NewWindow(wincfg)
	if err != nil {
		return nil, err
	}

	editor := &Editor{
		timestamp: 0.0,
		playing:   false,
		audio:     &AudioManager{},

		atlas:  atlas,
		window: window,
	}
	return editor, nil
}

func (editor *Editor) update() {
	// update audio position based on scroll
	editor.timestamp += int(editor.window.MouseScroll().Y * SCROLL_CONSTANT)
	if editor.timestamp < 0 {
		editor.timestamp = 0
	}
}

func (editor *Editor) draw() {
	if editor.beatmap == nil {
		return
	}

	im := imdraw.New(nil)

	// draw timeline at the bottom
	// (1366 x 34) @ (0, 0)
	im.Color = pixel.RGB(0.2, 0.2, 0.2)
	im.Push(pixel.V(0, 0))
	im.Push(pixel.V(1366, 48))
	im.Rectangle(0)

	// playfield
	// (1056 x 594) @ (310, 30)

	im.Draw(editor.window)

	// draw audio timestamp
	timestamp := text.New(pixel.V(20, 20), editor.atlas)
	length := editor.audio.GetSongLength()
	fmt.Fprintf(timestamp, "%s (%.02f%%)", FormatTimestamp(editor.timestamp), float64(editor.timestamp)*100.0/float64(length))
	timestamp.Draw(editor.window, pixel.IM)
}

func (editor *Editor) Open(filename string) error {
	dir := filepath.Dir(filename)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	// convert to utf-8
	beatmap, err := osu.ParseBeatmap(f)
	if err != nil {
		return err
	}

	// open audio
	err = editor.audio.Open(filepath.Join(dir, beatmap.AudioFilename))
	if err != nil {
		return err
	}

	editor.beatmap = EditorBeatmapFrom(beatmap)
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
