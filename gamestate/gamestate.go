package gamestate

import (
	resourcemanager "github.com/bird-mtn-dev/resource-manager"
	"github.com/eduardolat/goeasyi18n"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/noppikinatta/bamenn"
	input "github.com/quasilyte/ebitengine-input"
)

var STATE *GameState

const (
	ActionPauseMusic input.Action = iota
	ActionNext
	ActionDown
	ActionUp
	ActionLeft
	ActionRight
)

type SceneType uint32

const (
	SceneTitle SceneType = iota
	SceneGame
)

type GameState struct {
	ResourceMgr  resourcemanager.ResourceManager
	AudioPlayer  *audio.Player
	InputSystem  input.System
	InputHandler *input.Handler
	I18n         *goeasyi18n.I18n
	scenes       map[SceneType]ebiten.Game
	seq          *bamenn.Sequence

	T func(translateKey string, options ...goeasyi18n.Options) string
}

func (gs *GameState) TransitionScene(scene SceneType) *bamenn.Sequence {
	if gs.seq == nil {
		gs.seq = bamenn.NewSequence(gs.scenes[scene])
	} else {
		gs.seq.Switch(gs.scenes[scene])
	}
	return gs.seq
}

func (gs *GameState) AddScene(sceneType SceneType, scene ebiten.Game) {
	if gs.scenes == nil {
		gs.scenes = map[SceneType]ebiten.Game{}
	}
	gs.scenes[sceneType] = scene
}
