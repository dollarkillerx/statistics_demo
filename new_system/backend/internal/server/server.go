package server

import (
	"github.com/dollarkillerx/backend/internal/api"
	"github.com/dollarkillerx/backend/internal/conf"
	"github.com/dollarkillerx/backend/internal/storage"
)

type Server struct {
	storage   *storage.Storage
	apiServer *api.ApiServer
	conf      conf.Config
}

func NewServer(storage *storage.Storage, conf conf.Config) *Server {
	return &Server{storage: storage, apiServer: api.NewApiServer(storage, conf), conf: conf}
}

func (s *Server) Run() error {
	return s.apiServer.Run()
}
