package dkt

import (
	"github.com/google/uuid"
	"odkt/server/db/entity"
	"odkt/server/game"
)

type ODKRoom struct {
}

func CreateNewODKRoom(maxPlayers uint, owner *entity.User) *game.Room {
	if maxPlayers < 2 || maxPlayers > 4 {
		return nil
	}
	var odkRoom ODKRoom
	return &game.Room{
		ID:         uuid.New().String(),
		State:      game.LOBBY,
		Type:       game.DKT,
		MaxPlayers: maxPlayers,
		Owner:      owner,
		Players:    []*entity.User{},
		GameRoom:   odkRoom,
	}
}
