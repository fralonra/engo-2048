package scenes

const (
	sceneSplash   = "splash"
	sceneGame     = "game"

	messageChangeLabel = "changeLabel"
)

const (
	fontSizeL = 32
	fontSizeM = 24
	fontSizeS = 18
)

var (
	Splash   *SplashScene
	Game     *GameScene
)

func init() {
	// Scene init
	Splash = &SplashScene{}
	Splash.Name = sceneSplash

	Game = &GameScene{}
	Game.Name = sceneGame
}
