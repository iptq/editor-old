package editor

import (
	"os"
	"path/filepath"
	"time"

	"editor/osu"

	"github.com/veandco/go-sdl2/sdl"
)

// Editor contains the state of the editor.
type Editor struct {
	quit bool

	// current audio timestamp in milliseconds
	timestamp  int
	playing    bool
	audio      *AudioManager
	beatmap    *EditorBeatmap
	lastUpdate time.Time

	seeker *Seeker

	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewEditor() (*Editor, error) {
	// initialize window
	window, renderer, err := sdl.CreateWindowAndRenderer(1366, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	editor := &Editor{
		quit: false,

		timestamp: 0.0,
		playing:   false,
		audio:     &AudioManager{},

		window:   window,
		renderer: renderer,
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

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch kind := event.(type) {
		case *sdl.QuitEvent:
			editor.quit = true
		case *sdl.KeyboardEvent:
			switch kind.Keysym.Sym {
			case sdl.K_SPACE:
				editor.playing = !editor.playing
			}
		}
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
	_ = float64(editor.timestamp) * 100.0 / float64(length)

	// im := imdraw.New(nil)

	// editor.seeker.Draw(ctx)

	// // playfield
	// im.Color = pixel.RGB(0.1, 0.1, 0.1)
	// im.Push(pixel.V(155, 48))
	// im.Push(pixel.V(1211, 642))
	// im.Rectangle(0)

	// // timeline
	// im.Color = pixel.RGB(0.2, 0.2, 0.2)
	// im.Push(pixel.V(0, 642))
	// im.Push(pixel.V(1366, 720))
	// im.Rectangle(0)

	// // toolbar
	// im.Color = pixel.RGB(0.3, 0.3, 0.3)
	// im.Push(pixel.V(0, 720))
	// im.Push(pixel.V(1366, 768))
	// im.Rectangle(0)

	// im.Draw(editor.window)
	// ctx.Finish()

	// // draw audio timestamp
	// formatted := FormatTimestamp(editor.timestamp)
	// timestamp := text.New(pixel.V(20, 20), editor.atlas)
	// fmt.Fprintf(timestamp, "%s (%.02f%%)", formatted, percent)
	// timestamp.Draw(editor.window, pixel.IM)
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
	for !editor.quit {
		editor.update()
		editor.draw()

		editor.window.UpdateSurface()

		// TODO: vsync
		// this is ~50fps max
		time.Sleep(time.Millisecond * 20)
	}
}
