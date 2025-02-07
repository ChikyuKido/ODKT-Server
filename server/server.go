package server

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"odkt/server/connection"
	"odkt/server/route"
	"odkt/server/route/auth"
	"odkt/server/route/middleware"
	"strings"
)

type Server struct {
	Messages    chan connection.Message
	Connections []*connection.Connection
	Upgrader    websocket.Upgrader
	Addr        *string
}

var (
	server = Server{
		Messages:    make(chan connection.Message),
		Connections: make([]*connection.Connection, 0),
		Upgrader:    websocket.Upgrader{},
		Addr:        flag.String("addr", "localhost:8080", "http service address"),
	}
)

func Start() {
	r := gin.Default()
	route.InitRouter(r)
	r.GET("/ws", func(c *gin.Context) {
		connectionHandler(c.Writer, c.Request)
	})
	r.GET("/test", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"test": "test"})
	})
	go readMessages()
	err := r.Run(*server.Addr)
	if err != nil {
		logrus.Fatal(err)
	}
}
func (s *Server) HandleMessage(msg connection.Message) bool {
	s.Messages <- msg
	return false
}
func (s *Server) ConnectionClosed(conn *connection.Connection) bool {
	conn.Conn.Close()
	logrus.Infof("Connection from %v closed", conn.Conn.RemoteAddr())
	for i, c := range s.Connections {
		if c == conn {
			s.Connections = append(s.Connections[:i], s.Connections[i+1:]...)
			break
		}
	}
	return false
}

func connectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := server.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Failed to upgrade connection: %v", err)
		return
	}
	authentication := r.Header.Get("Authentication")
	if authentication == "" || !strings.Contains(authentication, "Bearer") || auth.LoginTokens[strings.Split(authentication, " ")[1]] == nil {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Authentication authentication is invalid"))
		conn.Close()
		logrus.Infof("Connection from %v closed due to invalid authentication", conn.RemoteAddr())
		return
	}
	token := strings.Split(authentication, " ")[1]
	user := auth.LoginTokens[token]
	delete(auth.LoginTokens, token)
	logrus.Infof("New connection from %v", conn.RemoteAddr())
	c := connection.NewConnection(conn)
	c.AddConnectionHandler(&server, true)
	c.User = user
	server.Connections = append(server.Connections, c)
}

func readMessages() {
	for {
		message := <-server.Messages
		go func() {
			message.Conn.SendMessage([]byte(message.Msg))
		}()
	}
}
