package router

import (
	"github.com/gin-gonic/gin"
	"github.com/musishere/chat-app/internal/user"
	"github.com/musishere/chat-app/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginUser)
	r.GET("/logout", userHandler.LogoutUser)

	r.POST("/ws/createroom", wsHandler.CreateRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
