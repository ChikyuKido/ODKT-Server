package repo

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"odkt/server/db"
	"odkt/server/db/entity"
	"time"
)

func InsertNewUser(username, password string) bool {
	user := entity.User{
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().Unix(),
		UUID:      uuid.New().String(),
	}
	if err := db.DB().Create(&user).Error; err != nil {
		logrus.Errorf("failed to create use: %v", err)
		return false
	}
	return true
}
func GetAllUsers() []entity.User {
	var users []entity.User
	if err := db.DB().Find(&users).Error; err != nil {
		logrus.Errorf("failed to get all users: %v", err)
		return nil
	}
	return users
}
func GetUserByUsername(username string) *entity.User {
	var user entity.User
	if err := db.DB().First(&user, entity.User{Username: username}).Error; err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil
	}
	return &user
}
func GetUserByID(id uint) *entity.User {
	var user entity.User
	if err := db.DB().Where(entity.User{ID: id}).First(&user).Error; err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return nil
	}
	return &user
}
func DoesUserByIDExists(id uint) bool {
	var user entity.User
	if err := db.DB().First(&user, entity.User{ID: id}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return true
	}
	return true
}

func DoesUserByUsernameExist(username string) bool {
	var user entity.User
	err := db.DB().First(&user, entity.User{Username: username}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logrus.Errorf("failed to get user: %v", err)
		return true
	}
	return true
}

func DeleteUser(userID uint) bool {
	var user = entity.User{
		ID: userID,
	}
	if err := db.DB().Delete(&user).Error; err != nil {
		logrus.Errorf("failed to delete user: %v", err)
		return false
	}
	return true
}
