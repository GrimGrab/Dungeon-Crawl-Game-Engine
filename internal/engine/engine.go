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

func New() *Engine {
	return &Engine{
		players:    make([]*player.Player, 0),
		gameMap:    world.NewMap(),
		depthTrack: nil, // Initialize with nil, can be set later when adventure is chosen
		state:      State{Round: 0, Phase: "initial", Turn: nil},
	}
}
