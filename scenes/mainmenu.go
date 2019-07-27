package scenes

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/fralonra/engo-utils"
)

type MainMenuScene struct {
	utils.Scene
}

func (*MainMenuScene) Preload() {}

func (*MainMenuScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&MainMenuSystem{})
	w.AddSystem(&utils.ClickableSystem{})

	entities := []utils.MenuEntity{
		utils.MenuEntity{
			Text: "New",
			OnClick: func() {
				engo.SetSceneByName(sceneGame, false)
			},
		},
		utils.MenuEntity{
			Text: "Quit",
			OnClick: func() {
				engo.Exit()
			},
		},
	}
	menu := utils.Menu{
		World:   w,
		Font:    mdFont,
		Texture: "ui/button.png",
		Gap:     30,
		Position: engo.Point{
			X: 30,
			Y: 30,
		},
		Entities: entities,
	}
	menu.Init()
}

type MainMenuSystem struct {
	utils.System
}

func (s *MainMenuSystem) Update(dt float32) {}
