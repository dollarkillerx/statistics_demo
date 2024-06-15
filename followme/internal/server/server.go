package server

import (
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
	s.router()
	return s.app.Run()
}
