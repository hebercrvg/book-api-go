package server

import (
	"books/controllers"
	"books/infra/database"
	"books/models"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		books := main.Group("books")
		{
			books.POST("/", controllers.CreateBook)
			books.GET("/", controllers.GetBooks)
			books.GET("/:id", controllers.GetBook)
		}
		comments := books.Group("comments")
		{
			comments.POST("/", controllers.CreateComment)
			comments.GET("/:bookId", controllers.GetComments)
		}
		main.GET("/benchmark-go", func(c *gin.Context) {
			db := database.GetDatabase()

			row := models.Benchmark{
				Title:       "Test",
				Description: "Test",
			}

			err := db.Create(&row).Error

			if err != nil {
				c.JSON(500, err)
				return
			}

			var rows []models.Benchmark

			err = db.Find(&rows).Error

			if err != nil {
				c.JSON(500, err)
				return
			}

			c.JSON(200, rows)
		})
	}

	return router
}
