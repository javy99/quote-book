package storage

import (
	"errors"
	"math/rand"
	"sync"

	"github.com/javy99/quote-book/internal/model"
)

// MemoryStore is an in-memory storage for quotes.
type MemoryStore struct {
	mu     sync.Mutex
	quotes []model.Quote
	nextID int
}

// NewMemoryStore initializes a new MemoryStore with an empty quotes slice and a random seed for ID generation.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		quotes: make([]model.Quote, 0),
		nextID: 1,
	}
}

// Add adds a new quote to the store and returns it with an assigned ID.
func (s *MemoryStore) Add(quote model.Quote) model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()
	quote.ID = s.nextID
	s.nextID++
	s.quotes = append(s.quotes, quote)
	return quote
}

// GetAll retrieves all quotes from the store.
func (s *MemoryStore) GetAll() []model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]model.Quote(nil), s.quotes...)
}

// GetByAuthor retrieves all quotes by a specific author from the store.
func (s *MemoryStore) GetByAuthor(author string) []model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()
	var filtered []model.Quote
	for _, q := range s.quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}
	return filtered
}

// GetRandom returns a randomly selected quote from the store.
// Returns an error if no quotes are available.
func (s *MemoryStore) GetRandom() (model.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.quotes) == 0 {
		return model.Quote{}, errors.New("no quotes available")
	}
	return s.quotes[rand.Intn(len(s.quotes))], nil
}

// DeleteByID removes a quote with the specified ID from the store.
// Returns true if the quote was found and deleted, false otherwise.
func (s *MemoryStore) DeleteByID(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}
	return false
}
