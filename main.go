package main

import (
	"github.com/gin-gonic/gin"

	books "go-microservices/internal/app"
)

func main() {
	router := gin.Default()
	router.GET("/books", books.GetBooks)
	router.GET("/books/:id", books.GetBookByID)
	router.POST("/books", books.PostBook)

	router.Run("localhost:8080")
}
