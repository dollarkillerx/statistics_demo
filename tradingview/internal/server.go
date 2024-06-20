package internal

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	app *gin.Engine

	mu      sync.Mutex
	dataMap map[string]TVResp
}

type TVResp struct {
	Ticker    string `json:"ticker"`
	Action    string `json:"action"`
	Contracts string `json:"contracts"`
	Price     string `json:"price"`
}

func (t *TVResp) ToJSON() string {
	marshal, err := json.Marshal(t)
	if err == nil {
		return string(marshal)
	}

	return ""
}

func (s *Server) Run() error {
	s.app = gin.Default()
	s.dataMap = map[string]TVResp{}

	InitLog()

	s.app.POST("/tradingview", s.tradingview)
	s.app.GET("/all", s.all)
	s.app.GET("/by/:Ticker", s.byTicker)

	return s.app.Run("127.0.0.1:8374")
}

func (s *Server) tradingview(ctx *gin.Context) {
	var input TVResp
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Error().Msgf("input error: %s", err)
		ctx.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	log.Info().Msgf("input: %s", input.ToJSON())

	// storage
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dataMap[input.Ticker] = input

	ctx.JSON(200, gin.H{"success": true})
}

func (s *Server) all(ctx *gin.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx.JSON(200, s.dataMap)
}

func (s *Server) byTicker(ctx *gin.Context) {
	param := strings.TrimSpace(ctx.Param("Ticker"))
	s.mu.Lock()
	defer s.mu.Unlock()
	resp, ex := s.dataMap[param]
	if !ex {
		ctx.JSON(404, gin.H{
			"error": "not found",
		})
		log.Error().Msgf("獲取不存在: %s", param)
		return
	}

	ctx.JSON(200, resp)
}
