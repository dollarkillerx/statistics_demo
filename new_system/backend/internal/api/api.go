package api

import (
	"fmt"

	"github.com/dollarkillerx/backend/internal/conf"
	"github.com/dollarkillerx/backend/internal/middleware"
	"github.com/dollarkillerx/backend/internal/storage"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	storage *storage.Storage
	conf    conf.Config
	app     *gin.Engine
}

func NewApiServer(storage *storage.Storage, conf conf.Config) *ApiServer {
	return &ApiServer{storage: storage, conf: conf}
}

func (a *ApiServer) Run() error {

	a.app = gin.New()
	a.app.Use(middleware.HttpRecover())
	//a.app.Use(middleware.RateLimiter())
	a.app.Use(gin.Logger())
	a.app.Use(middleware.Cors())

	a.Router()

	return a.app.Run(fmt.Sprintf("127.0.0.1:%s", a.conf.ServiceConfiguration.Port))
}
