package controller

type Connection interface {
	Send(interface{}) error
	Close() error
	ID() string
}

// Controller interface for handling game logic
type Controller interface {
	HandleMessage(conn Connection, msg *Message) error
	OnConnect(conn Connection) error
	OnDisconnect(conn Connection) error
}
