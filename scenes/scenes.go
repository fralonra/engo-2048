package scenes

const (
	sceneSplash   = "splash"
	sceneMainMenu = "mainmenu"
	sceneGame     = "game"
)

const (
	fontSizeL = 64
	fontSizeM = 32
	fontSizeS = 24
)

var (
	Splash   *SplashScene
	MainMenu *MainMenuScene
	Game     *GameScene
)

func init() {
	// Scene init
	Splash = &SplashScene{}
	Splash.Name = sceneSplash

	MainMenu = &MainMenuScene{}
	MainMenu.Name = sceneMainMenu

	Game = &GameScene{}
	Game.Name = sceneGame
}
