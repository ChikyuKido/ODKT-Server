package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		fmt.Println(requestData)
		if requestData.Username == "" || requestData.Password == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "username and password are incorrect"})
			return
		}
		//Todo: login

		c.JSON(http.StatusOK, gin.H{"token": "toke"})
	}
}
