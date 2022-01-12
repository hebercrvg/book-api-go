package controllers

import (
	"books/infra/database"
	"books/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	bookId, err := strconv.Atoi(c.Param("bookId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot get bookId: "+err.Error())
		return
	}

	var comments []models.BookComment

	db := database.GetDatabase()

	err = db.Find(&comments, models.BookComment{BookID: bookId}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot get comments: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}

func CreateComment(c *gin.Context) {
	var comment models.BookComment

	err := c.ShouldBindJSON(&comment)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot bind json: "+err.Error())
		return
	}

	db := database.GetDatabase()

	err = db.Create(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, "Cannot create comment: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, &comment)
}
