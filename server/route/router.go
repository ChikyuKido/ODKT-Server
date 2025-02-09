package route

import (
	"github.com/gin-gonic/gin"
	"odkt/server/route/auth"
	"odkt/server/route/middleware"
	"odkt/server/route/room"
	"odkt/server/route/websocket"
)

func InitRouter(r *gin.Engine) {
	initAuthRoutes(r)
	initRoomRoutes(r)
	initWebsocketRoutes(r)
}

func initWebsocketRoutes(r *gin.Engine) {
	group := r.Group("/ws")
	group.Use(middleware.AuthMiddleware())
	group.GET("joinRoom", websocket.JoinRoom())
}

func initAuthRoutes(r *gin.Engine) {
	group := r.Group("/api/v1/auth/")
	group.POST("register", auth.Register())
	group.POST("login", auth.Login())
}
func initRoomRoutes(r *gin.Engine) {
	group := r.Group("/api/v1/room/")
	group.Use(middleware.AuthMiddleware())
	group.POST("create", room.CreateRoom())
	group.GET("list", room.ListRooms())
}
