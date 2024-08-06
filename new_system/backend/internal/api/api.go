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
	a.app.Use(middleware.Cors())
	a.app.Use(middleware.HttpRecover())
	a.app.Use(gin.Logger())

	a.Router()

	return a.app.Run(fmt.Sprintf("0.0.0.0:%s", a.conf.ServiceConfiguration.Port))
}
