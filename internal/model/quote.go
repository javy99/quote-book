package model

// Quote represents a quote entry with an ID, author, and the quote text.
type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
