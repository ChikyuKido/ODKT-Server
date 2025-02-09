package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"odkt/server/connection"
	"odkt/server/game"
	"odkt/server/store"
	"odkt/server/util"
)

var upgrader = websocket.Upgrader{}

func JoinRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			RoomID string `form:"roomID"`
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logrus.Warnf("Failed to upgrade connection: %v", err)
			return
		}
		user := util.GetUserFromContextWithoutError(c)
		if user == nil {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(3000, "Authentication is invalid"))
			conn.Close()
			logrus.Infof("Connection from %v closed due to invalid authentication", conn.RemoteAddr())
			return
		}
		if user.JoinedRoom != "" && user.JoinedRoom != "awaitingJoin" {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "User already joined in a room"))
			conn.Close()
			return
		}
		if err := c.ShouldBindQuery(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
			return
		}

		roomValue, ok := store.RoomStore.Load(requestData.RoomID)
		if !ok {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Room with that id does not exist"))
			conn.Close()
			return
		}
		userConn := connection.Connection{
			Conn: conn,
			User: user,
		}
		room := roomValue.(*game.Room)
		if er := room.JoinRoom(&userConn); er != nil {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Could not join room: "+err.Error()))
			conn.Close()
			return
		}
	}
}
