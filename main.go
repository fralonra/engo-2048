package main

import (
	"github.com/EngoEngine/engo"
	"github.com/fralonra/engo-2048/scenes"
)

const (
	Title  = "2048"
	Width  = 324
	Height = 324
)

func main() {
	engo.RegisterScene(scenes.Game)

	opts := engo.RunOptions{
		Title:          Title,
		Width:          Width,
		Height:         Height,
		StandardInputs: true,
	}

	engo.Run(opts, scenes.Splash)
}
