package server

import (
	"followme/internal/conf"
	"followme/internal/storage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app     *gin.Engine
	storage *storage.Storage
}

func New() *Server {
	return &Server{
		app:     gin.Default(),
		storage: storage.NewStorage(),
	}
}

func (s *Server) Run() error {
	s.app.Use(auth(conf.Config.Header))

	s.router()
	return s.app.Run(conf.Config.Address)
}

func auth(token string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header != token {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
