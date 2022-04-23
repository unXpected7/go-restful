package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api/book"
)

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{
		"message": "Hello World",
		"id":      id,
		"title":   title,
	})
}

//? localhost:8080/1/abc

func QueryHandler(c *gin.Context) {
	query := c.Query("query")
	price := c.DefaultQuery("price", "0")

	c.JSON(200, gin.H{
		"message": "Hello World",
		"query":   query,
		"price":   price,
	})
}

//? localhost:8080/query?query=golang&price=100

func PostBooksHandler(c *gin.Context) {
	var book book.Book
	if err := c.ShouldBindJSON(&book); err != nil {

		errorMessages := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s:%s", err.ActualTag(), err.Field())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(400, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(200, gin.H{
		"title":     book.Title,
		"price":     book.Price,
		"sub_books": book.SubBooks,
	})
}
