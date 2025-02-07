package room

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"odkt/server/game"
	"odkt/server/store"
)

func ListRooms() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rooms []*game.Room
		store.RoomStore.Range(func(key string, value interface{}) bool {
			rooms = append(rooms, value.(*game.Room))
			return true
		})

		c.JSON(http.StatusOK, gin.H{"rooms": rooms})
	}
}
