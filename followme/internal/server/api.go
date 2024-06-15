package server

import (
	"followme/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) router() {
	s.app.POST("/report", s.Report)
}

func (s *Server) Report(ctx *gin.Context) {
	type ReportRequest struct {
		Orders []models.Order `json:"orders"`
	}

}
