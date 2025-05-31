package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/javy99/quote-book/internal/handler"
	"github.com/javy99/quote-book/internal/model"
	"github.com/javy99/quote-book/internal/storage"
)

func makeRequest(handler http.HandlerFunc, method, path string, body io.Reader) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	handler(w, r)
	return w
}

func setupHandler() (*handler.QuoteHandler, *storage.MemoryStore) {
	store := storage.NewMemoryStore()
	return handler.NewQuoteHandler(store), store
}

func TestCreateQuote(t *testing.T) {
	h, _ := setupHandler()
	quote := model.Quote{Author: "Javy", Quote: "Relativity"}
	body, _ := json.Marshal(quote)
	res := makeRequest(h.CreateQuote, http.MethodPost, "/quotes", bytes.NewBuffer(body))

	if res.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", res.Code)
	}
}

func TestCreateQuote_InvalidJSON(t *testing.T) {
	h, _ := setupHandler()
	res := makeRequest(h.CreateQuote, http.MethodPost, "/quotes", strings.NewReader("invalid"))

	if res.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", res.Code)
	}
}

func TestGetQuotes(t *testing.T) {
	h, store := setupHandler()
	store.Add(model.Quote{Author: "Javy", Quote: "Relativity"})

	res := makeRequest(h.GetQuotes, http.MethodGet, "/quotes", nil)
	if res.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", res.Code)
	}
}

func TestGetQuotesByAuthor(t *testing.T) {
	h, store := setupHandler()
	store.Add(model.Quote{Author: "Javy", Quote: "Relativity"})

	req := httptest.NewRequest(http.MethodGet, "/quotes?author=Javy", nil)
	w := httptest.NewRecorder()
	h.GetQuotes(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetRandomQuote(t *testing.T) {
	h, store := setupHandler()
	store.Add(model.Quote{Author: "Javy", Quote: "Relativity"})

	res := makeRequest(h.GetRandomQuote, http.MethodGet, "/quotes/random", nil)
	if res.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", res.Code)
	}
}

func TestGetRandomQuote_NoQuotes(t *testing.T) {
	h, _ := setupHandler()
	res := makeRequest(h.GetRandomQuote, http.MethodGet, "/quotes/random", nil)
	if res.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", res.Code)
	}
}

func TestDeleteQuote_MissingID(t *testing.T) {
	h, _ := setupHandler()
	res := makeRequest(h.DeleteQuote, http.MethodDelete, "/quotes", nil)
	if res.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", res.Code)
	}
}

func TestDeleteQuote_InvalidID(t *testing.T) {
	h, _ := setupHandler()
	res := makeRequest(h.DeleteQuote, http.MethodDelete, "/quotes/abc", nil)
	if res.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", res.Code)
	}
}

func TestDeleteQuote_ValidID(t *testing.T) {
	h, store := setupHandler()
	quote := store.Add(model.Quote{Author: "Javy", Quote: "Test"})

	quoteID := fmt.Sprintf("%d", quote.ID)

	res := makeRequest(h.DeleteQuote, http.MethodDelete, "/quotes/"+quoteID, nil)

	if res.Code != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", res.Code)
	}
}

func TestDeleteQuote_NotFound(t *testing.T) {
	h, _ := setupHandler()
	res := makeRequest(h.DeleteQuote, http.MethodDelete, "/quotes/100", nil)
	if res.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", res.Code)
	}
}
