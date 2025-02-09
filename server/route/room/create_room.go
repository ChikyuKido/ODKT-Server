package room

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"odkt/server/game"
	"odkt/server/game/dkt"
	"odkt/server/store"
	"odkt/server/util"
)

func CreateRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			Name         string                 `json:"name"`
			RoomType     game.RoomType          `json:"room_type"`
			MaxUsers     uint                   `json:"max_users"`
			GameSettings map[string]interface{} `json:"game_settings"`
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
		if requestData.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
			return
		}
		var room *game.Room = nil
		if requestData.RoomType == game.DKT {
			room = dkt.CreateNewODKRoom(requestData.MaxUsers, requestData.Name, user.UUID)
			if room == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new ODK room"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the room type is not supported"})
			return
		}

		store.RoomStore.Store(room.ID, room)
		user.JoinedRoom = "awaitingJoin"
		c.JSON(http.StatusOK, gin.H{"message": "room successfully created", "roomID": room.ID})
	}
}
