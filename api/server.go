package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mekuanint12/simple_api/db/sqlc"
)

type Server struct {
	users  *db.CreateUser
	router *gin.Engine
}

func NewServer(users *db.CreateUser) *Server {
	server := &Server{users: users}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	authRouter := router.Group("/user")
	authRouter.POST("/register", server.createAccount)
	router.GET("/login", server.getAccount)
	server.router = router
	return server
}
