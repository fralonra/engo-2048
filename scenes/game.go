package scenes

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/fralonra/engo-utils"
	"github.com/fralonra/go-2048/core"
)

const (
	buttonUp    = "up"
	buttonDown  = "down"
	buttonLeft  = "left"
	buttonRight = "right"

	buttonReset = "reset"
)

type GameScene struct {
	utils.Scene
}

func (*GameScene) Preload() {}

func (*GameScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	// add systems
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&GameSystem{
		game: core.NewGame(),
	})
	w.AddSystem(&utils.ClickableSystem{})

	// render
	label1 := utils.Label{
		World: w,
		Font:  lgFont,
		Text:  "Game",
		Position: engo.Point{
			X: 300,
			Y: 240,
		},
	}
	label1.Init()

	button := utils.Button{
		World: w,
		Font:  smFont,
		Text:  "Go back",
		Position: engo.Point{
			X: 100,
			Y: 500,
		},
	}
	button.Init()
	button.OnClick(func() {
		engo.SetSceneByName(sceneMainMenu, false)
	})

	// register buttons
	engo.Input.RegisterButton(buttonUp, engo.KeyArrowUp)
	engo.Input.RegisterButton(buttonDown, engo.KeyArrowDown)
	engo.Input.RegisterButton(buttonLeft, engo.KeyArrowLeft)
	engo.Input.RegisterButton(buttonRight, engo.KeyArrowRight)
	engo.Input.RegisterButton(buttonReset, engo.KeyR)
}

type GameSystem struct {
	utils.System

	game *core.Game
}

func (g *GameSystem) Update(dt float32) {
	// handle keys
	if engo.Input.Button(buttonUp).JustPressed() {
		g.game.ToTop()
	}
	if engo.Input.Button(buttonDown).JustPressed() {
		g.game.ToBottom()
	}
	if engo.Input.Button(buttonLeft).JustPressed() {
		g.game.ToLeft()
	}
	if engo.Input.Button(buttonRight).JustPressed() {
		g.game.ToRight()
	}
	if engo.Input.Button(buttonReset).JustPressed() {
		g.game = core.NewGame()
	}

	// render
	for idx := 0; idx < core.Size; idx++ {
		row := a.game.GetRow(idx)
		displayRow := []string{}
		for _, item := range row {
			var text string
			if item > 0 {
				text = strconv.Itoa(item)
			} else {
				text = ""
			}
			displayRow = append(displayRow, text)
		}
		a.table.Rows = append(a.table.Rows, displayRow)
	}
}
