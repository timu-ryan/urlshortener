package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLink(c *gin.Context) {
	shortName := c.Param("shortname")
	originalLink, ok := ShortLinks[shortName]
	if !ok {
		c.String(http.StatusNotFound, "Not Found!")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, originalLink)
}
