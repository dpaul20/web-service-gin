package controllers

import (
	"errors"
	"net/http"
	"web-service-gin/dtos"
	"web-service-gin/middlewares"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"msg":    `Books`,
		"data":   books,
	})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate body
	var body dtos.CreateBookInput
	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]middlewares.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = middlewares.ErrorMsg{Field: fe.Field(), Message: middlewares.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	// Create book
	book := models.Book{Title: body.Title, Author: body.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"msg":    `Book Created`,
		"data":   book,
	})
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"msg":    `Book Deleted`,
		"data":   book,
	})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input dtos.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"msg":    `Book found`,
		"data":   book,
	})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"msg":    `Book Deleted`,
		"data":   "",
	})
}
