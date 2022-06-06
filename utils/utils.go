package utils

import (
	"errors"
	"library/Basic-Golang-Api/data"
)

// get data by id (usually we can take for it from database)
func BookById(id string) (*data.Book, error) {
	for i, b := range data.Books {
		if b.ID == id {
			return &data.Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}
