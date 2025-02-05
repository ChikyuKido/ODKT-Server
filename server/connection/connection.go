package connection

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"odkt/server/db/entity"
)

type ConnectionHandler interface {
	HandleMessage(msg Message) bool
	ConnectionClosed(conn *Connection) bool
}

type Connection struct {
	Conn               *websocket.Conn
	connectionHandlers []ConnectionHandler
	User               *entity.User
	send               chan []byte
}

type Message struct {
	Conn *Connection
	Msg  []byte
}

func NewConnection(conn *websocket.Conn) *Connection {
	con := Connection{
		Conn: conn,
		send: make(chan []byte),
	}
	go con.readMessage()
	go con.writeMessages()
	return &con
}
func (c *Connection) AddConnectionHandler(handler ConnectionHandler, first bool) {
	if first {
		c.connectionHandlers = append(c.connectionHandlers[:0], append([]ConnectionHandler{handler}, c.connectionHandlers[:0]...)...)
	} else {
		c.connectionHandlers = append(c.connectionHandlers, handler)
	}
}
func (c *Connection) RemoveConnectionHandler(handler ConnectionHandler) {
	for i, h := range c.connectionHandlers {
		if h == handler {
			c.connectionHandlers = append(c.connectionHandlers[:i], c.connectionHandlers[i+1:]...)
			break
		}
	}
}
func (c *Connection) readMessage() {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("Client %s unexpectedly closed the connection: %v", c.Conn.RemoteAddr().String(), err)
			}
			for _, h := range c.connectionHandlers {
				if h.ConnectionClosed(c) {
					break
				}
			}
			break
		}
		for _, h := range c.connectionHandlers {
			if h.HandleMessage(Message{
				Conn: c,
				Msg:  message,
			}) {
				break
			}
		}

	}
}

func (c *Connection) SendMessage(message []byte) {
	select {
	case c.send <- message:
	default:
		logrus.Warn("Send buffer full, dropping message")
	}
}

func (c *Connection) writeMessages() {
	for message := range c.send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			logrus.Errorf("Error writing message: %v", err)
			for _, h := range c.connectionHandlers {
				if h.ConnectionClosed(c) {
					break
				}
			}
			break
		}
	}
}
