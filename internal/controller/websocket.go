package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log/slog"
	"sync"
)

type WSConnection struct {
	conn   *websocket.Conn
	id     string
	send   chan []byte
	closed bool
	mu     sync.RWMutex
	log    *slog.Logger
}

func (c *WSConnection) Send(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return errors.New("connection closed")
	}
	c.mu.RUnlock()

	select {
	case c.send <- data:
		return nil
	default:
		return errors.New("send buffer full")
	}
}

func (c *WSConnection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	c.closed = true

	close(c.send)
	return c.conn.Close()

}

func (c *WSConnection) ID() string {
	return c.id
}

func (c *WSConnection) readPump(controller Controller) {
	defer func() {
		controller.OnDisconnect(c)
		c.Close()
	}()

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.log.Error("websocket error", slog.String("error", err.Error()), slog.String("conn_id", c.ID()))
			}
			break
		}

		if err := controller.HandleMessage(c, &msg); err != nil {
			c.log.Info("handle message error", slog.String("error", err.Error()), slog.String("conn_id", c.ID()))
			err := c.Send(map[string]string{
				"error": err.Error(),
			})
			if err != nil {
				c.log.Error("send error", slog.String("error", err.Error()), slog.String("conn_id", c.ID()))
			}
		}
	}
}

func (c *WSConnection) writePump() {
	defer c.Close()

	for data := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			c.log.Error("write error", slog.String("error", err.Error()), slog.String("conn_id", c.ID()))
			return
		}
	}
}
