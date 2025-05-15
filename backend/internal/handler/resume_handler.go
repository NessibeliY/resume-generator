package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/NessibeliY/resume-generator/internal/models"
	"github.com/NessibeliY/resume-generator/internal/service"
)

type Handler struct {
	svc *service.AIService
}

func NewHandler(svc *service.AIService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GenerateResume(c *gin.Context) {
	var req models.ResumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	text, err := h.svc.GenerateResume(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate resume"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, models.ResumeResponse{ResumeText: text})
}
