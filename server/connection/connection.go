package connection

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type MessageHandler interface {
	HandleMessage(msg Message)
}

type Connection struct {
	conn       *websocket.Conn
	msgHandler MessageHandler
	send       chan []byte
}

type Message struct {
	Conn *Connection
	Msg  []byte
}

func NewConnection(conn *websocket.Conn, handler MessageHandler) *Connection {
	con := Connection{
		conn:       conn,
		msgHandler: handler,
		send:       make(chan []byte),
	}
	go con.readMessage()
	go con.writeMessages()
	return &con
}

func (c *Connection) readMessage() {
	defer func() {
		err := c.conn.Close()
		if err != nil {
			logrus.Errorf("Error closing connection: %v", err)
		}
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("Client %s unexpectedly closed the connection: %v", c.conn.RemoteAddr().String(), err)
			}
			break
		}
		c.msgHandler.HandleMessage(Message{
			Conn: c,
			Msg:  message,
		})
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
	defer c.conn.Close()

	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			logrus.Errorf("Error writing message: %v", err)
			break
		}
	}
}
