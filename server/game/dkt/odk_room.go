package dkt

import (
	"github.com/google/uuid"
	"odkt/server/connection"
	"odkt/server/game"
)

type ODKRoom struct {
	Room game.Room
}

func CreateNewODKRoom(maxPlayers uint, name, ownerUUID string) *game.Room {
	if maxPlayers < 2 || maxPlayers > 4 {
		return nil
	}
	var odkRoom ODKRoom
	room := game.Room{
		ID:         uuid.New().String(),
		Name:       name,
		State:      game.AWAITING_OWNER,
		Type:       game.DKT,
		MaxPlayers: maxPlayers,
		Owner:      ownerUUID,
		Players:    []*connection.Connection{},
		GameRoom:   odkRoom,
	}
	return &room
}
