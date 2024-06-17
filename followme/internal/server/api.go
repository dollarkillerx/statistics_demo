package server

import (
	"followme/internal/conf"
	"followme/internal/models"
	"github.com/gin-gonic/gin"

	"log"
)

func (s *Server) router() {
	s.app.POST("/release", auth(conf.Config.Header), s.release)
	s.app.GET("/subscription", s.subscription)
}

func (s *Server) release(ctx *gin.Context) {
	type ReportRequest struct {
		Orders []models.Order `json:"orders"`
	}

	var reportRequest ReportRequest

	if err := ctx.ShouldBindJSON(&reportRequest); err != nil {
		log.Println("error: ", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	s.storage.SetOrder(reportRequest.Orders)
	ctx.JSON(200, gin.H{"message": "success"})
}

func (s *Server) subscription(ctx *gin.Context) {
	orders := s.storage.GetOrders()
	ctx.JSON(200, gin.H{"orders": orders})
}
