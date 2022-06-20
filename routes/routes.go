package routes

import (
	"library/Basic-Golang-Api/adapter/database"
	"library/Basic-Golang-Api/adapter/database/postgres14"
	"library/Basic-Golang-Api/data"
	"net/http"

	"library/Basic-Golang-Api/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	DatabaseService database.DatabaseService
}

func (r Router) getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Books)
}

func (r Router) getBookIdByBookName(c *gin.Context) {
	name := c.Param("name")
	name, err := r.DatabaseService.CreateBook()
}

func (r Router) getBookByBookId(c *gin.Context) {
	id := c.Param("id")
	book, err := utils.BookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func (r Router) createBookOrderByBookId(c *gin.Context) {
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

func (r Router) createBook(c *gin.Context) {
	var newBook data.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	data.Books = append(data.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func (r Router) returnBookOrderByBookId(c *gin.Context) {
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

func InitRoutes(cfg config.Config) *gin.Engine {
	// connect to db adapter
	postgresService := postgres14.NewService(cfg)

	// initiate Router func
	routerHand := Router{
		DatabaseService: postgresService,
	}

	// gin route
	router := gin.Default()
	router.GET("/books", routerHand.getBooks)
	router.GET("/books", routerHand.getBookIdByBookName)
	router.GET("/books/:id", routerHand.getBookByBookId)
	router.POST("/books", routerHand.createBook)
	router.PATCH("/checkout", routerHand.createBookOrderByBookId)
	router.PATCH("/return", routerHand.returnBookOrderByBookId)

	return router
}
