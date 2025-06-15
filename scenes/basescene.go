package scenes

import (
	"github.com/bird-mtn-dev/ebitengine_template/gamestate"
	"github.com/joelschutz/stagehand"
)

type BaseScene struct {
	// your scene fields
	sm     *stagehand.SceneManager[*gamestate.GameState]
	state  *gamestate.GameState
	Width  int
	Height int
}

func (s *BaseScene) Load(state *gamestate.GameState, manager stagehand.SceneController[*gamestate.GameState]) {
	// your load code
	s.sm = manager.(*stagehand.SceneManager[*gamestate.GameState]) // This type assertion is important
	s.state = state
}

func (s *BaseScene) Unload() *gamestate.GameState {
	// your unload code
	return s.state
}

func (s *BaseScene) Layout(w, h int) (int, int) {
	s.Width = w
	s.Height = h
	return w, h
}

func (s *BaseScene) Update() error {
	s.state.InputSystem.Update()

	if s.state.AudioPlayer != nil {
		if s.state.InputHandler.ActionIsJustPressed(gamestate.ActionPauseMusic) {
			if s.state.AudioPlayer.IsPlaying() {
				s.state.AudioPlayer.Pause()
			} else {
				s.state.AudioPlayer.Play()
			}
		}
	}
	return nil
}
