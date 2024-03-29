package editor

import (
	"image/color"
)

type SeekerConfig struct {
	// X, Y denote the left middle position of the seeker
	X, Y int

	LineWidth             int
	KnobWidth, KnobHeight int

	LineColor color.Color
	KnobColor color.Color

	Editor *Editor
}

type Seeker struct {
	config SeekerConfig
}

func NewSeeker(config SeekerConfig) *Seeker {
	return &Seeker{
		config: config,
	}
}

func (s Seeker) Draw() {
	elapsed := s.config.Editor.GetCurrentTime()
	_ = float64(elapsed*100.0) / float64(s.config.Editor.GetSongLength())

	// draw seeker at the bottom
	// ctx.Color = pixel.RGB(0.2, 0.2, 0.2)
	// ctx.Push(pixel.V(0, 0))
	// ctx.Push(pixel.V(1366, 48))
	// ctx.Rectangle(0)

	// // seeker line
	// ctx.Color = pixel.RGB(0.5, 0.5, 0.5)
	// ctx.Push(pixel.V(180, 23))
	// ctx.Push(pixel.V(180+1000, 25))
	// ctx.Rectangle(0)

	// // seeker handle
	// ctx.Color = pixel.RGB(1.0, 1.0, 1.0)
	// x := percent * 10
	// ctx.Push(pixel.V(180-4+x, 4))
	// ctx.Push(pixel.V(180+4+x, 44))
	// ctx.Rectangle(0)
}
