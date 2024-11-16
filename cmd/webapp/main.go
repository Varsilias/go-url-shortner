package main

import (
	"database/sql"
	"fmt"
	"github.com/Varsilias/go-url-shortner/internal/controllers"
	"github.com/Varsilias/go-url-shortner/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer func(sqlite *sql.DB) {
		err := sqlite.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sqlite)

	if err := db.CreateTable(sqlite); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("URL Path", request.URL.Path)
		if request.URL.Path == "/" {
			controllers.ShowIndex(writer, request)
		} else {
			controllers.Redirect(sqlite)(writer, request)
		}
	})
	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	log.Fatal(http.ListenAndServe(":9000", nil))
}
