package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/javy99/quote-book/internal/model"
	"github.com/javy99/quote-book/internal/storage"
)

// QuoteHandler handles HTTP requests for quotes.
type QuoteHandler struct {
	Store *storage.MemoryStore
}

// NewQuoteHandler creates a new QuoteHandler with the given store.
func NewQuoteHandler(store *storage.MemoryStore) *QuoteHandler {
	return &QuoteHandler{Store: store}
}

// CreateQuote handles the creation of a new quote.
func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote model.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	saved := h.Store.Add(quote)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(saved)
}

// GetQuotes retrieves all quotes or filters them by author if they 'author' query parameter is provided.
func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var quotes []model.Quote
	if author != "" {
		quotes = h.Store.GetByAuthor(author)
	} else {
		quotes = h.Store.GetAll()
	}
	json.NewEncoder(w).Encode(quotes)
}

// GetRandomQuote retrieves a random quote from the store.
func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.Store.GetRandom()
	if err != nil {
		http.Error(w, "No quotes found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quote)
}

// DeleteQuote deletes a quote by its ID.
func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if !h.Store.DeleteByID(id) {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
