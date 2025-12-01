package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{ID: "1", Title: "Go语言", Author: "西门庆", Price: 79.00},
	{ID: "2", Title: "GoWeb语言", Author: "西门庆", Price: 79.00},
	{ID: "3", Title: "Python语言", Author: "西门庆", Price: 79.00},
	{ID: "4", Title: "Javascript语言", Author: "西门庆", Price: 79.00},
}

func main() {
	fmt.Printf("Hello world")
	r := gin.Default()

	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.GET("/books/:id", func(c *gin.Context) {

		id := c.Param("id")

		for _, book := range books {
			if book.ID == id {
				c.JSON(http.StatusOK, book)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})

	})

	r.Run(":8080")
}
