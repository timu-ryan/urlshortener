package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timu-ryan/urlshortener/internal/config"
	"github.com/timu-ryan/urlshortener/internal/handler"
)

func main() {
	cfg := config.NewConfig()

	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.POST("/", handler.PostLink(cfg.BaseURL))
	r.GET("/:shortname", handler.GetLink)
	r.Run(cfg.ServerAddress)
}
