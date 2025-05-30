package model

// Quote represents a quote entry with an ID, author, and quote text.
// It is used in JSON requests and responses.
// swagger:model
type Quote struct {
	// ID is the unique identifier of the quote.
	// example: 1
	ID int `json:"id"`
	// Author is the name of the person who said the quote.
	// example: Albert Einstein
	Author string `json:"author"`
	// Quote is the text of the quote.
	// example: Life is like riding a bicycle. To keep your balance, you must keep moving.
	Quote string `json:"quote"`
}
