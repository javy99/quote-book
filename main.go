package main

import (
	"log"
	"net/http"

	_ "github.com/javy99/quote-book/docs"
	"github.com/javy99/quote-book/internal/handler"
	"github.com/javy99/quote-book/internal/storage"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Quote API
// @version 1.0
// @description REST API for managing quotes
// @host localhost:8080
// @Basepath /
func main() {
	store := storage.NewMemoryStore()
	h := handler.NewQuoteHandler(store)

	// REST API endpoints
	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateQuote(w, r)
		case http.MethodGet:
			h.GetQuotes(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/quotes/random", h.GetRandomQuote)
	http.HandleFunc("/quotes/", h.DeleteQuote)

	// Swagger UI endpoint
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
