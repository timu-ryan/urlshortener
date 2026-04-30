package handler

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var ShortLinks = make(map[string]string)

func PostLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}

	testKey := getRandomString()
	
	ShortLinks[testKey] = string(body)
	w.Write([]byte("Успех! Ссылка доступна по " + testKey + "."))
}


var alphabet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func getRandomString() string {
	var builder strings.Builder

	for range 8 {
		builder.WriteString(string(alphabet[rand.Intn(len(alphabet))]))	
	}
	return builder.String()
}
