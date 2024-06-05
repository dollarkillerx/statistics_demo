package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"seahorse/internal/conf"
	"seahorse/internal/storage"
)

type ApiServer struct {
	conf    *conf.Conf
	storage *storage.Storage
	gin     *gin.Engine
}

func NewApiServer(conf *conf.Conf, storage *storage.Storage) *ApiServer {
	return &ApiServer{
		conf:    conf,
		storage: storage,
	}
}

func (a *ApiServer) Start() {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.Use(gin.Logger())

	a.gin = engine
	a.RegisterRoutes()
	log.Println("Starting API server on", a.conf.Address)
	err := engine.Run(a.conf.Address)
	if err != nil {
		panic(err)
	}
}
