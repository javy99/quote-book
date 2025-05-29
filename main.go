package main

import (
	"log"
	"net/http"

	"github.com/javy99/quote-book/internal/handler"
	"github.com/javy99/quote-book/internal/storage"
)

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

	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
