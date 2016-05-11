package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/core"
	"engo.io/engo/demos/demoutils"
)

type DefaultScene struct{}

var (
	edgeMargin  float64 = 20
	scrollSpeed float32 = 700

	worldWidth  int = 800
	worldHeight int = 800
)

func (*DefaultScene) Preload() {}

// Setup is called before the main loop is started
func (*DefaultScene) Setup(w *ecs.World) {
	core.SetBackground(color.White)
	w.AddSystem(&core.RenderSystem{})

	// The most important line in this whole demo:
	w.AddSystem(&core.EdgeScroller{scrollSpeed, edgeMargin})

	// Create the background; this way we'll see when we actually scroll
	demoutils.NewBackground(w, worldWidth, worldHeight, color.RGBA{102, 153, 0, 255}, color.RGBA{102, 173, 0, 255})
}

func (*DefaultScene) Type() string { return "Game" }

func main() {
	opts := engo.RunOptions{
		Title:  "EdgeScroller Demo",
		Width:  worldWidth,
		Height: worldHeight,
	}
	engo.Run(opts, &DefaultScene{})
}
