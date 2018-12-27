package ui

import (
	"github.com/faiface/pixel/imdraw"
)

type Context struct {
	*Window
	G *imdraw.IMDraw
}
