package store

import (
	"odkt/server/db/entity"
	"odkt/server/db/repo"
)

func GetUserByUUID(uuid string) *entity.User {
	var user *entity.User = nil
	if value, ok := UserIDStore.Load(uuid); ok {
		user = value.(*entity.User)
	} else {
		user = repo.GetUserByUUID(uuid)
		UserStore.Store(user.Username, user)
		UserIDStore.Store(uuid, user)
	}
	return user
}
func GetUserByUsername(username string) *entity.User {
	var user *entity.User = nil
	if value, ok := UserStore.Load(username); ok {
		user = value.(*entity.User)
	} else {
		user = repo.GetUserByUsername(username)
		UserStore.Store(username, user)
		UserIDStore.Store(user.UUID, user)
	}
	return user
}
