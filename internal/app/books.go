package books

import (
	"net/http"

	books_data "go-microservices/internal/data"

	"github.com/gin-gonic/gin"
)

// getBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books_data.Books)
}

// postBook adds a book from JSON received in the request body.
func PostBook(c *gin.Context) {
	var newBook books_data.Book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	books_data.Books = append(books_data.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of books, looking for
	// an book whose ID value matches the parameter.
	for _, a := range books_data.Books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
