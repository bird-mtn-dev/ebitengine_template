package gamestate

import (
	resourcemanager "github.com/bird-mtn-dev/resource-manager"
	"github.com/eduardolat/goeasyi18n"
	"github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionPauseMusic input.Action = iota
	ActionNext
	ActionDown
	ActionUp
	ActionLeft
	ActionRight
)

type GameState struct {
	ResourceMgr  resourcemanager.ResourceManager
	AudioPlayer  *audio.Player
	InputSystem  input.System
	InputHandler *input.Handler
	I18n         *goeasyi18n.I18n

	T func(translateKey string, options ...goeasyi18n.Options) string
}
