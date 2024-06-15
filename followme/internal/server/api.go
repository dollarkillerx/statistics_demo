package server

import (
	"followme/internal/models"
	"github.com/gin-gonic/gin"

	"log"
)

func (s *Server) router() {
	s.app.POST("/release", s.release)
	s.app.POST("/subscription", s.subscription)
}

func (s *Server) release(ctx *gin.Context) {
	type ReportRequest struct {
		Orders []models.Order `json:"orders"`
	}

	if err := ctx.ShouldBindJSON(&ReportRequest{}); err != nil {
		log.Println("error: ", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	s.storage.SetOrder(ReportRequest{}.Orders)
	ctx.JSON(200, gin.H{"message": "success"})
}

func (s *Server) subscription(ctx *gin.Context) {
	orders := s.storage.GetOrders()
	ctx.JSON(200, gin.H{"orders": orders})
}
