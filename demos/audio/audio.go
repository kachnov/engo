package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/core"
	"log"
)

type DefaultScene struct{}

type Whoop struct {
	ecs.BasicEntity
	core.AudioComponent
}

func (*DefaultScene) Preload() {
	core.AudioSystemPreload()

	err := engo.Files.Load("326488.wav")
	if err != nil {
		log.Fatal(err)
	}
}

func (*DefaultScene) Setup(w *ecs.World) {
	core.SetBackground(color.White)

	w.AddSystem(&core.RenderSystem{})
	w.AddSystem(&core.AudioSystem{})
	w.AddSystem(&WhoopSystem{})

	whoop := Whoop{BasicEntity: ecs.NewBasic()}
	whoop.AudioComponent = core.AudioComponent{File: "326488.wav", Repeat: true, Background: true, RawVolume: 1}

	// Let's add our whoop to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *core.AudioSystem:
			// Note we are giving a `nil` reference to the `SpeedComponent`. This is because the documentation of the
			// AudioSystem says the `SpeedComponent` is only required when `AudioComponent.Background` is `false`.
			// In our case, it is `true` (it's a background noise, i.e. not tied to a location in the game world),
			// so we can omit it.
			sys.Add(&whoop.BasicEntity, &whoop.AudioComponent, nil)
		}
	}
}

func (*DefaultScene) Type() string { return "Game" }

type WhoopSystem struct {
	goingUp bool
}

// Remove is empty, because this system doesn't do anything with entities (note there's no `Add` method either)
func (w *WhoopSystem) Remove(basic ecs.BasicEntity) {}

func (w *WhoopSystem) Update(dt float32) {
	d := float64(dt * 0.1)
	if w.goingUp {
		core.MasterVolume += d
	} else {
		core.MasterVolume -= d
	}

	if core.MasterVolume < 0 {
		core.MasterVolume = 0
		w.goingUp = true
	} else if core.MasterVolume > 1 {
		core.MasterVolume = 1
		w.goingUp = false
	}
}

func main() {
	opts := engo.RunOptions{
		Title:  "Audio Demo",
		Width:  1024,
		Height: 640,
	}
	engo.Run(opts, &DefaultScene{})
}
