package controllers

import (
	"books/infra/database"
	"books/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(context *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(http.StatusBadRequest, "cannot bind json: "+err.Error())
		return
	}

	result := db.Create(&book)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, "cannot create book: "+err.Error())
		return
	}

	context.JSON(http.StatusCreated, book)
}

func GetBooks(context *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Preload("Comments").Find(&books).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, "Cannot get books: "+err.Error())
		return
	}

	context.JSON(http.StatusOK, books)
}

func GetBook(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, "Cannot get book id: "+err.Error())
		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, id).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, "Cannot get book: "+err.Error())
		return
	}

	context.JSON(http.StatusOK, book)
}
