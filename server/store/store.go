package store

import "github.com/puzpuzpuz/xsync/v3"

var (
	// RoomStore stores all the rooms by its id
	RoomStore *xsync.Map
	// UserStore stores all the user by their usernames
	UserStore *xsync.Map
	// UserIDStore stores all the user by their uuid
	UserIDStore *xsync.Map
)

func InitStores() {
	RoomStore = xsync.NewMap()
	UserStore = xsync.NewMap()
	UserIDStore = xsync.NewMap()
}
