package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timu-ryan/urlshortener/internal/config"
	"github.com/timu-ryan/urlshortener/internal/handler"
	"github.com/timu-ryan/urlshortener/internal/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()

	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}
	defer logger.Log.Sync()

	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(logger.RequestLogger())
	r.POST("/", handler.PostLink(cfg.BaseURL))
	r.GET("/:shortname", handler.GetLink)

	logger.Log.Info("Starting server", zap.String("addr", cfg.ServerAddress))
	r.Run(cfg.ServerAddress)
}
