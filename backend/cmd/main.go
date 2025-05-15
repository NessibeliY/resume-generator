package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/NessibeliY/resume-generator/internal/config"
	"github.com/NessibeliY/resume-generator/internal/handler"
	"github.com/NessibeliY/resume-generator/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	srv := service.NewAIService(cfg.OpenAIKey)
	h := handler.NewHandler(srv)

	r.POST("/api/generate", h.GenerateResume)

	log.Printf("Starting server on port %s", cfg.Port)
	r.Run(cfg.Port)
}
