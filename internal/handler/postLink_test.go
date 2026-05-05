package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostLink(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		body         string
		wantStatus   int
		wantURLInMap bool
	}{
		{
			name:         "valid POST with URL",
			method:       http.MethodPost,
			body:         "https://practicum.yandex.ru",
			wantStatus:   http.StatusCreated,
			wantURLInMap: true,
		},
		{
			name:       "GET method not allowed",
			method:     http.MethodGet,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "PUT method not allowed",
			method:     http.MethodPut,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:       "DELETE method not allowed",
			method:     http.MethodDelete,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShortLinks = make(map[string]string)

			r := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			PostLink(w, r)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.wantStatus, res.StatusCode)

			if tt.wantURLInMap {
				resBody, err := io.ReadAll(res.Body)
				require.NoError(t, err)

				shortURL := string(resBody)
				assert.True(t, strings.HasPrefix(shortURL, "http://"), "ответ должен начинаться с http://")

				// извлекаем ключ из конца короткого URL и проверяем запись в хранилище
				parts := strings.Split(strings.TrimRight(shortURL, "\n"), "/")
				key := parts[len(parts)-1]
				require.NotEmpty(t, key)

				original, ok := ShortLinks[key]
				assert.True(t, ok, "ключ должен быть в ShortLinks")
				assert.Equal(t, tt.body, original)
			}
		})
	}
}
