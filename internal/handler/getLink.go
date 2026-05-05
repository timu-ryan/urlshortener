package handler

import "net/http"

func GetLink(w http.ResponseWriter, r *http.Request) {
	shortName := r.PathValue("shortname")
	originalLink, ok := ShortLinks[shortName]
	if !ok {
		http.Error(w, "Not Found!", http.StatusNotFound)
		return
	}

	// responseHeaders := w.Header()
	// responseHeaders.Set("Location", originalLink)
	http.Redirect(w, r, originalLink, http.StatusTemporaryRedirect)
}
