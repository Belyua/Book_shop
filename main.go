package main

import (
	"awesomeProject/controllers"
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	models.ConnectDatabase()

	r.GET("/create", controllers.ShowCreateBookForm)
	r.GET("/delete/", controllers.ShowDeleteBookForm)
	r.GET("/books", controllers.FindBooks)
	r.POST("/delete", controllers.DeleteBook)
	r.POST("/create", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"data": "hello world"})
	})

	address := "localhost:8081"

	fmt.Println("Server is running at:", address)

	r.Run(address)
}
