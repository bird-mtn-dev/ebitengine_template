package main

import (
	"embed"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/bird-mtn-dev/ebitengine_template/gamestate"
	"github.com/bird-mtn-dev/ebitengine_template/scenes"
	resourcemanager "github.com/bird-mtn-dev/resource-manager"
	"github.com/eduardolat/goeasyi18n"
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

//go:embed assets/translations/en/*.json
var enFS embed.FS

func main() {
	setupLogger()

	ebiten.SetWindowTitle("Ebitengine Template")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	gamestate.STATE = newGameState()

	if err := ebiten.RunGame(gamestate.STATE.TransitionScene(gamestate.SceneTitle)); err != nil {
		log.Fatal(err)
	}
}

func setupLogger() {
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	mwriter := io.MultiWriter(os.Stderr, f)

	jsonHandler := slog.NewJSONHandler(mwriter, &slog.HandlerOptions{AddSource: true})
	slog.SetDefault(slog.New(jsonHandler))
}

func newGameState() *gamestate.GameState {
	gameState := &gamestate.GameState{ResourceMgr: resourcemanager.Create()}

	slog.Info("Setting up font")
	_ = gameState.ResourceMgr.Font.LoadStandardFonts()

	slog.Info("Setting up audio")
	//Example below on setting up a looping mp3
	//audio.NewContext(44100)
	//gameState.ResourceMgr.Audio.LoadPlayer("song", "assets/02_seaside_village.mp3")
	//gameState.Player, _ = gameState.ResourceMgr.Audio.GetPlayer("song", audiomanager.Looping())
	//gameState.Player.Play()

	slog.Info("Setting up keybindings")
	gameState.InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		gamestate.ActionPauseMusic: {input.KeyP, input.KeyPause},
		gamestate.ActionNext:       {input.KeyMouseRight},
		gamestate.ActionDown:       {input.KeyDown},
		gamestate.ActionUp:         {input.KeyUp},
		gamestate.ActionRight:      {input.KeyRight},
		gamestate.ActionLeft:       {input.KeyLeft},
	}
	gameState.InputHandler = gameState.InputSystem.NewHandler(0, keymap)

	slog.Info("Setting up I18N")
	gameState.I18n = goeasyi18n.NewI18n(goeasyi18n.Config{
		FallbackLanguageName:    "en",  // It's optional, the default value is "en"
		DisableConsistencyCheck: false, // It's optional, the default value is false
	})
	// Load english translations from JSON files
	enTranslations, err := goeasyi18n.LoadFromJsonFS(enFS, "assets/translations/en/*.json")
	if err != nil {
		panic(err)
	}
	gameState.I18n.AddLanguage("en", enTranslations)

	gameState.T = gameState.I18n.NewLangTranslateFunc("en")

	slog.Info("Setting up Scene Map")
	gameState.AddScene(gamestate.SceneGame, &scenes.GameScene{})
	gameState.AddScene(gamestate.SceneTitle, &scenes.TitleScene{})

	return gameState
}
