package game

import "odkt/server/db/entity"

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
