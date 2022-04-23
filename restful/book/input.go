package book

import "encoding/json"

type Book struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubBooks string      `json:"sub_books" binding:"required"`
}
