package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"odkt/server/db/entity"
	"odkt/server/db/repo"
	"odkt/server/util"
	"strings"
)

var LoginTokens = map[string]*entity.User{}

func Login() gin.HandlerFunc {
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
		user := repo.GetUserByUsername(requestData.Username)
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad credentials"})
			return
		}
		if !util.CheckPassword(user.Password, requestData.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad credentials"})
			return
		}
		id := uuid.New()
		LoginTokens[id.String()] = user
		token, err := util.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong. Please try again later."})
			return
		}
		c.JSON(http.StatusOK, gin.H{"jwt": token})
	}
}
