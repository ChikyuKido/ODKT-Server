package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"odkt/server/db/repo"
	"odkt/server/util"
	"strings"
)

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if strings.TrimSpace(requestData.Username) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username must not be empty"})
			return
		}
		if strings.TrimSpace(requestData.Password) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password must not be empty"})
			return
		}
		if repo.DoesUserByUsernameExist(requestData.Username) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exist"})
			return
		}
		hashedPassword, err := util.HashPassword(requestData.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			logrus.Errorf("Failed to hash password: %v", err)
			return
		}
		if !repo.InsertNewUser(requestData.Username, hashedPassword) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new user"})
			logrus.Errorf("Failed to create new user: %v", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "successful create an account"})
		return
	}
}
