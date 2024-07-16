package books

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	books_data "go-microservices/internal/data"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetBooks(t *testing.T) {
	// Mock the server
	r := SetUpRouter()
	r.GET("/books", GetBooks)

	// Make the request & send it to the mocked server
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Read the response
	var books []books_data.Book
	json.Unmarshal(w.Body.Bytes(), &books)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestPostBook(t *testing.T) {
// 	// Set up the route
// 	r := SetUpRouter()
// 	r.GET("/books", GetBooks)

// 	// Make the request
// 	req, _ := http.NewRequest("GET", "/books", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	// Read the response
// 	var books []books_data.Book
// 	json.Unmarshal(w.Body.Bytes(), &books)

// 	assert.Equal(t, http.StatusOK, w.Code)
// }
