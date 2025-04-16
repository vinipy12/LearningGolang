package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	urlMap = make(map[string]string)
	mu     sync.Mutex
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code string `json:"code"`
}

func generateCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func isValidUrl(url string) bool {
	return len(url) > 0 && (url[:7] == "http://" || url[:7] == "https:/")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if !isValidUrl(req.URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	code := generateCode()
	urlMap[code] = req.URL

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{Code: code})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[:1]

	mu.Lock()
	defer mu.Unlock()
	url, ok := urlMap[code]

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/redirect", redirectHandler)

	fmt.Println("Server starting on :8080")

	http.ListenAndServe(":8080", nil)
}
