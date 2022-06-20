package database

type DatabaseService interface {
	ReadBooks()
	ReadBookIdByBookName()
	ReadBookByBookId()
	CreateBook()
	UpdateBookQuantity()
}
