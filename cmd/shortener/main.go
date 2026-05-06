package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timu-ryan/urlshortener/internal/handler"
)

func main() {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.POST("/", handler.PostLink)
	r.GET("/:shortname", handler.GetLink)
	r.Run(":8080")
}
