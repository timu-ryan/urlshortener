package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLink(t *testing.T) {
	const (
		existingKey = "abc12345"
		targetURL   = "https://practicum.yandex.ru"
	)

	tests := []struct {
		name         string
		shortname    string
		setupStorage map[string]string
		wantStatus   int
		wantLocation string
	}{
		{
			name:         "existing short link redirects",
			shortname:    existingKey,
			setupStorage: map[string]string{existingKey: targetURL},
			wantStatus:   http.StatusTemporaryRedirect,
			wantLocation: targetURL,
		},
		{
			name:         "unknown short link returns 404",
			shortname:    "unknown1",
			setupStorage: map[string]string{},
			wantStatus:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShortLinks = tt.setupStorage

			r := httptest.NewRequest(http.MethodGet, "/"+tt.shortname, nil)
			r.SetPathValue("shortname", tt.shortname)
			w := httptest.NewRecorder()

			GetLink(w, r)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.wantStatus, res.StatusCode)

			if tt.wantLocation != "" {
				assert.Equal(t, tt.wantLocation, res.Header.Get("Location"))
			}
		})
	}
}
