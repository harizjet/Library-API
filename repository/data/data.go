package data

type BooksReponse struct {
	Books []BookResponse
}

type BookResponse struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Author   string `json:"author"`
}

type BookID struct {
	ID string `json:"id"`
}
