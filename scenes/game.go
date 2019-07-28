package scenes

import (
	"log"
	"strconv"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/fralonra/engo-utils"
	"github.com/fralonra/go-2048/colors"
	"github.com/fralonra/go-2048/core"
)

const (
	buttonUp    = "up"
	buttonDown  = "down"
	buttonLeft  = "left"
	buttonRight = "right"
	buttonReset = "reset"

	cellSize   = 76
	cellMargin = 4
)

type Cell struct {
	Position engo.Point
	World    *ecs.World
	Width    float32
	Height   float32
	Number   int

	box   *utils.Box
	label *utils.Label
}

func (c *Cell) Init() {
	c.box = &utils.Box{
		Color:    colors.CellColor(c.Number),
		Width:    c.Width,
		Height:   c.Height,
		Position: c.Position,
		World:    c.World,
	}
	c.box.Init()

	text := ""
	if c.Number != 0 {
		text = strconv.Itoa(c.Number)
	}
	c.label = &utils.Label{
		Font:  mdFont,
		Text:  text,
		World: c.World,
	}
	width, height, _ := c.label.TextDimensions()
	c.label.Position = engo.Point{
		X: c.box.SpaceComponent.Position.X + (cellSize-float32(width))/2,
		Y: c.box.SpaceComponent.Position.Y + (cellSize-float32(height))/2,
	}
	c.label.Init()
}

func (c *Cell) Reset(newNumber int) {
	if newNumber == c.Number {
		return
	}

	c.Number = newNumber
	c.box.RenderComponent.Color = colors.CellColor(newNumber)

	text := ""
	if newNumber != 0 {
		text = strconv.Itoa(newNumber)
	}

	c.label.SetText(text)
	if newNumber != 0 {
		width, height, _ := c.label.TextDimensions()
		c.label.SpaceComponent = common.SpaceComponent{
			Width:  float32(width),
			Height: float32(height),
			Position: engo.Point{
				X: c.box.SpaceComponent.Position.X + (cellSize-float32(width))/2,
				Y: c.box.SpaceComponent.Position.Y + (cellSize-float32(height))/2,
			},
		}
		if newNumber > 10 && width == 17 {
			log.Printf("newNumber: %v, %#v", newNumber, c.label)
		}
	}
}

type GameScene struct {
	utils.Scene
}

func (*GameScene) Preload() {}

func (*GameScene) Setup(u engo.Updater) {
	game := core.NewGame()
	cellTable := &[core.Size][core.Size]*Cell{}

	w, _ := u.(*ecs.World)

	// add systems
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&HUDTextSystem{})
	w.AddSystem(&GameSystem{
		game:      game,
		cellTable: cellTable,
	})
	w.AddSystem(&utils.ClickableSystem{})

	// render()
	for i := 0; i < core.Size; i++ {
		row := game.GetRow(i)
		for j, item := range row {
			cell := &Cell{
				Position: engo.Point{
					X: float32(j*(cellSize+cellMargin) + cellMargin),
					Y: float32(i*(cellSize+cellMargin) + cellMargin),
				},
				World:  w,
				Width:  cellSize,
				Height: cellSize,
				Number: item,
			}
			cell.Init()
			cellTable[i][j] = cell
		}
	}

	// register buttons
	engo.Input.RegisterButton(buttonUp, engo.KeyArrowUp)
	engo.Input.RegisterButton(buttonDown, engo.KeyArrowDown)
	engo.Input.RegisterButton(buttonLeft, engo.KeyArrowLeft)
	engo.Input.RegisterButton(buttonRight, engo.KeyArrowRight)
	engo.Input.RegisterButton(buttonReset, engo.KeyR)
}

type GameSystem struct {
	utils.System

	game      *core.Game
	cellTable *[core.Size][core.Size]*Cell
}

func (g *GameSystem) Update(dt float32) {
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

	g.renderGame()
	g.checkState()
}

func (g *GameSystem) renderGame() {
	for i := 0; i < core.Size; i++ {
		row := g.game.GetRow(i)
		for j := 0; j < core.Size; j++ {
			cell := g.cellTable[i][j]
			cell.Reset(row[j])
		}
	}
}

func (g *GameSystem) checkState() {
	switch g.game.State {
	case core.StateNormal:
		{
			message := ChangeLabel{
				Text: "MAX: " + strconv.Itoa(g.game.MaxNumber),
			}
			message.Message.Name = messageChangeLabel
			engo.Mailbox.Dispatch(message)
		}
	case core.StateWin:
		{
			message := ChangeLabel{
				Text: "You win",
			}
			message.Message.Name = messageChangeLabel
			engo.Mailbox.Dispatch(message)
		}
	case core.StateLost:
		{
			message := ChangeLabel{
				Text: "You lost",
			}
			message.Message.Name = messageChangeLabel
			engo.Mailbox.Dispatch(message)
		}
	}
}

type HUDTextSystem struct {
	utils.System

	label utils.Label
}

func (h *HUDTextSystem) New(w *ecs.World) {
	engo.Mailbox.Listen(messageChangeLabel, func(m engo.Message) {
		msg, ok := m.(ChangeLabel)
		if !ok {
			return
		}
		h.label.SetText(msg.Text)
	})

	h.label = utils.Label{
		World: w,
		Font:  smFont,
		Position: engo.Point{
			X: 10,
			Y: 10,
		},
	}
	h.label.Init()
}

type ChangeLabel struct {
	utils.Message

	Text string
}


