package handler

import (
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var ShortLinks = make(map[string]string)

func PostLink(baseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, "Ошибка чтения тела запроса")
			return
		}

		key := getRandomString()
		ShortLinks[key] = string(body)

		c.String(http.StatusCreated, baseURL+"/"+key)
	}
}


var alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func getRandomString() string {
	var builder strings.Builder

	for range 8 {
		builder.WriteString(string(alphabet[rand.Intn(len(alphabet))]))
	}
	return builder.String()
}
