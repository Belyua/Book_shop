package controllers

import (
	"awesomeProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Books",
		"Books": books,
	})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ShowCreateBookForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

//func ShowDeleteBookForm(c *gin.Context) {
//	c.HTML(http.StatusOK, "delete.html", nil)
//}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{
			"Error": err.Error(),
		})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.Redirect(http.StatusSeeOther, "/books")
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func ShowDeleteBookForm(c *gin.Context) {
	// Get list of books for the dropdown
	var books []models.Book
	if err := models.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.HTML(http.StatusOK, "delete.html", gin.H{
		"Books": books,
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.PostForm("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, "/books")
}
