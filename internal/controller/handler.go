package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type HandlerFunc func(conn Connection, params json.RawMessage) error

type GameController struct {
	handlers map[string]HandlerFunc
	log      *slog.Logger
}

func NewGameController(log *slog.Logger) *GameController {
	gc := &GameController{
		handlers: make(map[string]HandlerFunc),
		log:      log,
	}

	// Register handlers
	gc.handlers["move"] = gc.handleMove
	//gc.handlers["chat"] = gc.handleChat
	//gc.handlers["joinRoom"] = gc.handleJoinRoom

	return gc
}

func (gc *GameController) HandleMessage(conn Connection, msg *Message) error {
	handler, exists := gc.handlers[msg.Action]
	if !exists {
		return fmt.Errorf("unknown action: %s", msg.Action)
	}

	return handler(conn, msg.Params)
}

func (gc *GameController) handleMove(conn Connection, params json.RawMessage) error {
	//var moveParams MoveParams
	//if err := json.Unmarshal(params, &moveParams); err != nil {
	//	return fmt.Errorf("invalid move params: %w", err)
	//}

	// Handle movement logic
	// ...
	gc.log.Info("handle move action")
	return nil
}

func (gc *GameController) OnConnect(conn Connection) error {
	gc.log.Info("player connected", slog.String("conn_id", conn.ID()))
	// TODO: Add player to game state, send initial state, etc.
	return nil
}

func (gc *GameController) OnDisconnect(conn Connection) error {
	gc.log.Info("player disconnected", slog.String("conn_id", conn.ID()))
	// TODO: Remove player from game state, notify other players, etc.
	return nil
}
