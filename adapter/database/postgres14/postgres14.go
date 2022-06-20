package postgres14

import "database/sql"

type postgres14Impl struct {
	DbEngine *sql.DB
}

func (p *postgres14Impl) ReadBooks() {
	// (*data.BooksReponse, error)
}

func (p *postgres14Impl) ReadBookIdByBookName() {
	// (*data.BookResponse, error)
}

func (p *postgres14Impl) ReadBookByBookId() {
	// (*data.BookResponse, error)
}

func (p *postgres14Impl) CreateBook() {
	// (*data.BookResponse, error)
}

func (p *postgres14Impl) UpdateBookQuantity() {
	// (*data.BookResponse, error)
}
