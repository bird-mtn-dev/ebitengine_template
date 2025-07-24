package scenes

import (
	"github.com/bird-mtn-dev/ebitengine_template/gamestate"
)


type BaseScene struct {
	Width  int
	Height int
}

func (s *BaseScene) Layout(w, h int) (int, int) {
	s.Width = w
	s.Height = h
	return w, h
}

func (s *BaseScene) Update() error {
	gamestate.STATE.InputSystem.Update()

	if gamestate.STATE.AudioPlayer != nil {
		if gamestate.STATE.InputHandler.ActionIsJustPressed(gamestate.ActionPauseMusic) {
			if gamestate.STATE.AudioPlayer.IsPlaying() {
				gamestate.STATE.AudioPlayer.Pause()
			} else {
				gamestate.STATE.AudioPlayer.Play()
			}
		}
	}
	return nil
}
