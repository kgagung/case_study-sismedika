package model

// book.go adalah model untuk data buku
type Book struct {
	ID	 int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	PublishedYear int    `json:"published_year"`
}