package internal

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
}

func (s *Server) Run() error {
	s.app = gin.Default()

	s.app.POST("/tradingview", s.tradingview)

	return s.app.Run("127.0.0.1:8374")
}

func (s *Server) tradingview(ctx *gin.Context) {
	all, err := io.ReadAll(ctx.Request.Body)
	if err == nil {
		fmt.Println(string(all))
		os.WriteFile("tradingview.log", all, 0644)
	}

	ctx.JSON(200, gin.H{"success": true})
}
