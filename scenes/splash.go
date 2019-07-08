package scenes

import (
	"bytes"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/fralonra/engo-utils"
	"golang.org/x/image/font/gofont/gosmallcaps"
	"image/color"
)

const (
	buttonNext = "next"
)

var (
	err    error
	lgFont *common.Font
	smFont *common.Font
	mdFont *common.Font
)

type SplashScene struct {
	utils.Scene
}

func (*SplashScene) Preload() {
	engo.Files.Load("ui/button.png")
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
}

func (*SplashScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&SplashSystem{})

	lgFont = &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: fontSizeL,
	}
	err = lgFont.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	mdFont = &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: fontSizeM,
	}
	err = mdFont.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	smFont = &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: fontSizeS,
	}
	err = smFont.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	label1 := utils.Label{
		World: w,
		Font:  lgFont,
		Text:  "Splash",
		Position: engo.Point{
			X: 300,
			Y: 240,
		},
	}
	label1.Init()

	engo.Input.RegisterButton(buttonNext, engo.KeySpace)
}

type SplashSystem struct {
	utils.System
}

func (*SplashSystem) Update(dt float32) {
	if engo.Input.Button(buttonNext).JustPressed() {
		engo.SetSceneByName(sceneMainMenu, false)
	}
}
