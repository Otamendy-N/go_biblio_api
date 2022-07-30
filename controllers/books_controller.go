package controllers

import (
	"go_biblio_api/catalog"
	"go_biblio_api/entities"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	collection *catalog.BookCollection
}

type BookControllerInterface interface {
	paginateBooks(c *gin.Context)
	createBook(c *gin.Context)
	ConfigureRouter(router *gin.Engine)
}

var collection = catalog.BookCollection{}

func paginateBooks(c *gin.Context) {
	books := collection.GetAll()
	c.JSON(200, books)
}

func createBook(c *gin.Context) {
	var book entities.Book
	c.BindJSON(&book)
	collection.Add(&book)
	c.JSON(200, book)
}

func ConfigureRouter(router *gin.Engine) {
	router.GET("/books", paginateBooks)
	router.POST("/books", createBook)
}
