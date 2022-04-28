package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "TDD With Python", Author: "Jazzy Jeff", Quantity: 2},
	{ID: 2, Title: "Prometheus Up and Running", Author: "Jazzy Jeff", Quantity: 5},
	{ID: 3, Title: "Learning Golang", Author: "Tech with Tim", Quantity: 6},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}
