package room

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"odkt/server/game"
	"odkt/server/store"
)

type RoomData struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	MaxPlayers uint   `json:"max_players"`
	Players    uint   `json:"players"`
	OwnerName  string `json:"owner_name"`
}

func ListRooms() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rooms []RoomData
		store.RoomStore.Range(func(key string, value interface{}) bool {
			if key == "" {
				return false
			}

			room := value.(*game.Room)
			owner := store.GetUserByUUID(room.Owner)
			if owner == nil {
				logrus.Errorf("failed to get owner from uuid in room list")
				return true
			}
			roomData := RoomData{
				UUID:       room.ID,
				Name:       room.Name,
				MaxPlayers: room.MaxPlayers,
				Players:    uint(len(room.Players)),
				OwnerName:  owner.Username,
			}
			rooms = append(rooms, roomData)
			return true
		})

		c.JSON(http.StatusOK, gin.H{"rooms": rooms})
	}
}
