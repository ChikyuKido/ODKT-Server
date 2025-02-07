package room

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"odkt/server/game"
	"odkt/server/store"
	"odkt/server/util"
)

func JoinRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			RoomID string `json:"room_id"`
		}
		user := util.GetUserFromContext(c)
		if user == nil {
			return
		}
		if user.JoinedRoom != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you have already joined a room"})
			return
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
			return
		}

		roomValue, ok := store.RoomStore.Load(requestData.RoomID)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "room does not exist"})
			return
		}
		room := roomValue.(*game.Room)
		if er := room.JoinRoom(user); er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}
		//TODO: return websocket link
		c.JSON(http.StatusOK, gin.H{"room": room})
	}
}
