package storage_test

import (
	"testing"

	"github.com/javy99/quote-book/internal/model"
	"github.com/javy99/quote-book/internal/storage"
)

func TestAddQuote(t *testing.T) {
	store := storage.NewMemoryStore()
	quote := model.Quote{Author: "Javy", Quote: "Test Quote"}
	saved := store.Add(quote)

	if saved.ID != 1 {
		t.Errorf("Expected ID 1, got %d", saved.ID)
	}
}

func TestGetAllQuotes(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Add(model.Quote{Author: "Javy", Quote: "Quote"})

	all := store.GetAll()
	if len(all) != 1 {
		t.Errorf("Expected 1 quote, got %d", len(all))
	}
}

func TestGetByAuthor(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Add(model.Quote{Author: "Javy", Quote: "Quote"})

	found := store.GetByAuthor("Javy")
	if len(found) != 1 {
		t.Errorf("Expected 1 quote by Javy, got %d", len(found))
	}
}

func TestGetRandomQuote(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Add(model.Quote{Author: "Javy", Quote: "Quote"})

	_, err := store.GetRandom()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestGetRandomQuote_Empty(t *testing.T) {
	store := storage.NewMemoryStore()

	_, err := store.GetRandom()
	if err == nil {
		t.Errorf("Expected error on empty store")
	}
}

func TestDeleteByID_Success(t *testing.T) {
	store := storage.NewMemoryStore()
	quote := store.Add(model.Quote{Author: "Javy", Quote: "Test"})

	ok := store.DeleteByID(quote.ID)
	if !ok {
		t.Errorf("Expected delete to return true")
	}

	if len(store.GetAll()) != 0 {
		t.Errorf("Expected store to be empty after delete")
	}
}

func TestDeleteByID_NotFound(t *testing.T) {
	store := storage.NewMemoryStore()
	ok := store.DeleteByID(999)

	if ok {
		t.Errorf("Expected delete to return false for non-existent ID")
	}
}
