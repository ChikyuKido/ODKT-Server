package game

import (
	"fmt"
	"odkt/server/db/entity"
)

type RoomType int
type RoomState int

const (
	DKT RoomType = iota
)
const (
	LOBBY RoomState = iota
	IN_GAME
	FINISHED
)

type Room struct {
	ID         string
	State      RoomState
	Type       RoomType
	MaxPlayers uint
	Owner      *entity.User
	Players    []*entity.User
	GameRoom   interface{}
}

func (r *Room) JoinRoom(user *entity.User) error {
	if uint(len(r.Players)) >= r.MaxPlayers-1 {
		return fmt.Errorf("max players reached")
	}
	r.Players = append(r.Players, user)
	user.JoinedRoom = r.ID
	return nil
}
