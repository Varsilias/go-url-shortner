package controllers

import (
	"database/sql"
	"fmt"
	"github.com/Varsilias/go-url-shortner/internal/db"
	"github.com/Varsilias/go-url-shortner/internal/url"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		originalUrl := r.FormValue("url")
		if originalUrl == "" {
			http.Error(w, "URL to shorten not provided", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://") {
			originalUrl = "https://" + originalUrl
		}

		urlExists, err := db.GetOriginalURLExists(lite, originalUrl)
		if err != nil {
			log.Println(fmt.Errorf("error checking if url exists: %w", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if urlExists {
			http.Error(w, "URL already exists", http.StatusBadRequest)
			return
		}
		shortURL := generateUniqueShortUrl(w, originalUrl, lite)

		err = db.StoreURL(lite, shortURL, originalUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]string{
			"ShortURL": shortURL,
		}

		t, err := template.ParseFiles("internal/views/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func generateUniqueShortUrl(w http.ResponseWriter, originalUrl string, lite *sql.DB) string {
	shortURL := url.Shorten(originalUrl)

	var guard = true

	shortURLExists, err := db.GetShortURLExists(lite, shortURL)
	if err != nil {
		log.Println(fmt.Errorf("error checking if short url exists: %w", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	for guard {
		if !shortURLExists {
			guard = false
		} else {
			shortURL = url.Shorten(originalUrl)
			shortURLExists, err = db.GetShortURLExists(lite, shortURL)
			if err != nil {
				log.Println(fmt.Errorf("error checking if short url exists: %w", err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return ""
			}
		}
	}
	return shortURL
}
