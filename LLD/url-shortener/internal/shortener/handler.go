package shortener

import (
	"encoding/json"
	"net/http"
)

// Request body for POST /shorten
type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

// Response body for GET /
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.LongURL == "" {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}

	// generate short code
	shortCode := generateRandomShortURL()

	// save mapping
	saveShortURL(shortCode, req.LongURL)

	resp := ShortenResponse{
		ShortURL: "http://localhost:8080/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// get longURL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:] // remove "/"

	longURL, exists := getLongURL(shortCode)
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
