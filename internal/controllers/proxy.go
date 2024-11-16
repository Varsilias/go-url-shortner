package controllers

import (
	"database/sql"
	"fmt"
	"github.com/Varsilias/go-url-shortner/internal/db"
	"net/http"
)

func Redirect(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]
		fmt.Println("shortUrl", shortUrl)
		if shortUrl == "" {
			http.Error(w, "URL not provided", http.StatusBadRequest)
			return
		}

		originalUrl, err := db.GetOriginalURL(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		fmt.Println("originalUrl", originalUrl)
		http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
	}
}
