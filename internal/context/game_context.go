package context

import (
	"SoB/internal/engine"
	"SoB/internal/player"
)

type GameContext struct {
	Engine       *engine.Engine
	ActivePlayer *player.Player
	Players      []*player.Player
}
