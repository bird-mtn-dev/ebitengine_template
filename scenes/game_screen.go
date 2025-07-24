package scenes

import (
	"fmt"
	"log/slog"

	"github.com/bird-mtn-dev/ebitengine_template/gamestate"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/setanarut/kamera/v2"
	"golang.org/x/image/math/f64"
)

type GameScene struct {
	BaseScene
	MainCamera  *kamera.Camera
	worldScreen *ebiten.Image
	target      f64.Vec2
}

func (s *GameScene) Update() error {
	_ = s.BaseScene.Update()

	s.updateCamera()

	return nil
}

// This function is used to update the camera. It will move the camera based on the arrow keys or if the
// cursor is near the edge of the screen.
func (s *GameScene) updateCamera() {

	x, y := ebiten.CursorPosition()
	if gamestate.STATE.InputHandler.ActionIsPressed(gamestate.ActionDown) || (y > s.Height-20 && y <= s.Height) {
		s.target[1] += 2
	}

	if gamestate.STATE.InputHandler.ActionIsPressed(gamestate.ActionUp) || (y < 20 && y >= 0) {
		s.target[1] -= 2
	}

	if gamestate.STATE.InputHandler.ActionIsPressed(gamestate.ActionRight) || (x > s.Width-20 && x <= s.Width) {
		s.target[0] += 2
	}

	if gamestate.STATE.InputHandler.ActionIsPressed(gamestate.ActionLeft) || (x < 20 && x >= 0) {
		s.target[0] -= 2
	}

	s.MainCamera.LookAt(s.target[0], s.target[1])
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	//Draw your world to worldScreen
	ebitenutil.DebugPrintAt(s.worldScreen, gamestate.STATE.T("draw_here"), 200, 200)
	s.MainCamera.Draw(s.worldScreen, &ebiten.DrawImageOptions{}, screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d, %d", s.Width, s.Height))
}

func (s *GameScene) OnStart() {
	slog.Info("Setting up World Screen")
	s.worldScreen = ebiten.NewImage(400, 400)

	slog.Info("Setting up Camera")
	s.MainCamera = kamera.NewCamera(0, 0, 0, 0)
	s.MainCamera.ZoomFactor = 1
	s.target[0] = 0
	s.target[1] = 0
	slog.Info("Setting up UI")
}
