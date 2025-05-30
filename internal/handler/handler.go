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

// CreateQuote godoc
// @Summary Add a new quote
// @Description Add a quote by providing author and quote text
// @Tags quotes
// @Accept json
// @Produce json
// @Param quote body model.Quote true "Quote object"
// @Success 201 {object} model.Quote
// @Failure 400 {string} string "Invalid input"
// @Router /quotes [post]
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

// GetQuotes godoc
// @Summary Get all quotes
// @Description Retrieve all quotes or filter by author using query parameter
// @Tags quotes
// @Produce json
// @Param author query string false "Author filter"
// @Success 200 {array} model.Quote
// @Router /quotes [get]
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

// GetRandomQuote godoc
// @Summary Get a random quote
// @Description Retrieve one random quote from the store
// @Tags quotes
// @Produce json
// @Success 200 {object} model.Quote
// @Failure 404 {string} string "No quotes found"
// @Router /quotes/random [get]
func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.Store.GetRandom()
	if err != nil {
		http.Error(w, "No quotes found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(quote)
}

// DeleteQuote godoc
// @Summary Delete a quote by ID
// @Description Remove a quote by its numeric ID
// @Tags quotes
// @Param id path int true "Quote ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Quote not found"
// @Router /quotes/{id} [delete]
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
