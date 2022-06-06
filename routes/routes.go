package routes

import (
	"library/Basic-Golang-Api/data"
	"library/Basic-Golang-Api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := utils.BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBookOrder(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	book, err := utils.BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook data.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	data.Books = append(data.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func returnBookOrder(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	book, err := utils.BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusCreated, book)
}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", createBookOrder)
	router.PATCH("/return", returnBookOrder)

	return router
}
