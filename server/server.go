package server

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"odkt/server/connection"
)

type Server struct {
	Messages chan connection.Message
}

var (
	upgrader    = websocket.Upgrader{}
	connections = make([]*connection.Connection, 0)
	addr        = flag.String("addr", "localhost:8080", "http service address")
	server      = Server{Messages: make(chan connection.Message)}
)

func Start() {

	http.HandleFunc("/ws", connectionHandler)
	readMessages()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logrus.Fatal(err)
	}
}
func (s *Server) HandleMessage(msg connection.Message) {
	s.Messages <- msg
}

func connectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Failed to upgrade connection: %v", err)
		return
	}
	logrus.Infof("New connection from %v", conn.RemoteAddr())
	connections = append(connections, connection.NewConnection(conn, &server))
}

func readMessages() {
	go func() {
		for {
			message := <-server.Messages
			logrus.Infof("Received message: %v", string(message.Msg))
			message.Conn.SendMessage([]byte("test"))
		}
	}()
}
