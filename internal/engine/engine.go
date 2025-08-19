package engine

import (
	"SoB/internal/depth_track"
	"SoB/internal/player"
	"SoB/internal/world"
)

type Engine struct {
	players    []*player.Player
	gameMap    *world.Map
	depthTrack *depth_track.DepthTracker
	state      State
}
