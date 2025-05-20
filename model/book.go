package model

import "sync"

// Book adalah model untuk data buku
type Book struct {
	ID	 int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

// BookStore adalah penyimpanan untuk buku
type BookStore struct {
	books []Book
	mu    sync.RWMutex
}

var (
	storeInstance *BookStore
	once         sync.Once
)

// GetBookStore mengembalikan instance BookStore
func GetBookStore() *BookStore {
	once.Do(func() {
		storeInstance = &BookStore{
			books: []Book{},
		}
	})
	return storeInstance
}

// GetBooks mengembalikan daftar semua buku
func (bs *BookStore) GetBooks() []Book {
	return bs.books
}

// GetBookByID mengembalikan buku berdasarkan ID
func (bs *BookStore) GetBookByID(id int) (*Book, bool) {
	for _, book := range bs.books {
		if book.ID == id {
			return &book, true
		}
	}
	return nil, false
}

// AddBook menambahkan buku baru ke dalam daftar
func (bs *BookStore) AddBook(book Book) {
	book.ID = len(bs.books) + 1
	bs.books = append(bs.books, book)
}

// UpdateBook memperbarui data buku berdasarkan ID
func (bs *BookStore) UpdateBook(id int, updatedBook Book) bool {
	for i, book := range bs.books {
		if book.ID == id {
			updatedBook.ID = id
			bs.books[i] = updatedBook
			return true
		}
	}
	return false
}

// DeleteBook menghapus buku berdasarkan ID
func (bs *BookStore) DeleteBook(id int) bool {
	for i, book := range bs.books {
		if book.ID == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
			return true
		}
	}
	return false
}