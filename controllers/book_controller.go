package controllers

import (
	"errors"
	"net/http"

	BookModel "jmc/bootcamp/models"

	"github.com/gin-gonic/gin"
)


func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, BookModel.ArrayBooks)
}

func getBookById(id string) (*BookModel.Book, error) {
	for i, b := range BookModel.ArrayBooks {
		if b.ID == id {
			return &BookModel.ArrayBooks[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func BookById(c *gin.Context) {
	bookId := c.Param("bookId")
	book, err := getBookById(bookId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook BookModel.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	BookModel.ArrayBooks = append(BookModel.ArrayBooks, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func ChecoutBook(c *gin.Context) {
	bookId, ok := c.GetQuery("bookId")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(bookId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if  book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}