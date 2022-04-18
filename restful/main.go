package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", rootHandler)
	v1.GET("/:id/:title", booksHandler)
	v1.GET("query", queryHandler)
	v1.POST("/books", postBooksHandler)

	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(200, gin.H{
		"message": "Hello World",
		"id":      id,
		"title":   title,
	})
}

//? localhost:8080/1/abc

func queryHandler(c *gin.Context) {
	query := c.Query("query")
	price := c.DefaultQuery("price", "0")

	c.JSON(200, gin.H{
		"message": "Hello World",
		"query":   query,
		"price":   price,
	})
}

//? localhost:8080/query?query=golang&price=100
type Book struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubBooks string      `json:"sub_books" binding:"required"`
}

func postBooksHandler(c *gin.Context) {
	var book Book
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
