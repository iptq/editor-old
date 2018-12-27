package ui

import (
	"github.com/faiface/pixel/imdraw"
)

type Context struct {
	window *Window
	*imdraw.IMDraw
}

func ContextFrom(window *Window) *Context {
	return &Context{
		IMDraw: imdraw.New(nil),
		window: window,
	}
}

func (ctx *Context) Finish() {
	ctx.IMDraw.Draw(ctx.window)
}
