package game

import (
	"fmt"
	"odkt/server/connection"
)

type RoomType int
type RoomState int

const (
	DKT RoomType = iota
)
const (
	AWAITING_OWNER RoomState = iota
	LOBBY
	IN_GAME
	FINISHED
)

type Room struct {
	ID         string
	Name       string
	State      RoomState
	Type       RoomType
	MaxPlayers uint
	Owner      string
	Players    []*connection.Connection
	GameRoom   interface{}
}

func (r *Room) JoinRoom(conn *connection.Connection) error {
	if r.State == AWAITING_OWNER {
		if conn.User.UUID != r.Owner {
			return fmt.Errorf("invalid owner %s", conn.User.UUID)
		}
		r.State = LOBBY
	}
	if uint(len(r.Players)) >= r.MaxPlayers-1 {
		return fmt.Errorf("max players reached")
	}
	r.Players = append(r.Players, conn)
	conn.User.JoinedRoom = r.ID
	return nil
}
