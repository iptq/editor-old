package editor

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"time"

	"editor/osu"
	"editor/ui"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

// Editor contains the state of the editor.
type Editor struct {
	// current audio timestamp in milliseconds
	timestamp  int
	playing    bool
	audio      *AudioManager
	beatmap    *EditorBeatmap
	lastUpdate time.Time

	seeker *Seeker

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

	seekerConfig := SeekerConfig{
		Editor: editor,
	}
	editor.seeker = NewSeeker(seekerConfig)

	return editor, nil
}

func (editor *Editor) GetCurrentTime() int {
	return editor.timestamp
}

func (editor *Editor) GetSongLength() int {
	return editor.audio.GetSongLength()
}

func (editor *Editor) update() {
	now := time.Now()
	elapsed := now.Sub(editor.lastUpdate)

	// update audio position based on scroll
	editor.timestamp += int(editor.window.MouseScroll().Y * -150.0)

	// change play/pause
	if editor.window.JustReleased(pixelgl.KeySpace) {
		editor.playing = !editor.playing
	}

	// if playing
	if editor.playing {
		editor.timestamp += int(elapsed.Nanoseconds() / 1000000)
	}

	editor.lastUpdate = now
}

func (editor *Editor) draw() {
	if editor.beatmap == nil {
		return
	}

	length := editor.audio.GetSongLength()
	percent := float64(editor.timestamp) * 100.0 / float64(length)

	ctx := ui.ContextFrom(editor.window)
	im := imdraw.New(nil)

	editor.seeker.Draw(ctx)

	// playfield
	im.Color = pixel.RGB(0.1, 0.1, 0.1)
	im.Push(pixel.V(155, 48))
	im.Push(pixel.V(1211, 642))
	im.Rectangle(0)

	// timeline
	im.Color = pixel.RGB(0.2, 0.2, 0.2)
	im.Push(pixel.V(0, 642))
	im.Push(pixel.V(1366, 720))
	im.Rectangle(0)

	// toolbar
	im.Color = pixel.RGB(0.3, 0.3, 0.3)
	im.Push(pixel.V(0, 720))
	im.Push(pixel.V(1366, 768))
	im.Rectangle(0)

	im.Draw(editor.window)
	ctx.Finish()

	// draw audio timestamp
	formatted := FormatTimestamp(editor.timestamp)
	timestamp := text.New(pixel.V(20, 20), editor.atlas)
	fmt.Fprintf(timestamp, "%s (%.02f%%)", formatted, percent)
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
	editor.lastUpdate = time.Now()

	// the main game loop
	for !editor.window.Closed() {
		editor.window.Clear(color.Black)

		editor.update()
		editor.draw()

		editor.window.Update()
	}
}
