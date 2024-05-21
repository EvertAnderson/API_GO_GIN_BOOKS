package controllers

import (
	"github.com/EvertAnderson/configs"
	"github.com/EvertAnderson/models"
	"github.com/gin-gonic/gin"
)

type BookRequestBody struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	Summary       string `json:"summary"`
	PublishedYear uint   `json:"publishedyear"`
}

func BookCreate(c *gin.Context) {
	body := BookRequestBody{}

	c.BindJSON(&body)

	book := &models.Book{
		Title:         body.Title,
		Author:        body.Author,
		Genre:         body.Genre,
		Summary:       body.Summary,
		PublishedYear: body.PublishedYear,
	}

	result := configs.DB.Create(&book)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &book)
}

func BookGet(c *gin.Context) {
	var books []models.Book
	configs.DB.Find(&books)
	c.JSON(200, &books)
	return
}

func BookGetById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	configs.DB.First(&book, id)
	c.JSON(200, &book)
	return
}

func BookUpdate(c *gin.Context) {

	id := c.Param("id")
	var book models.Book
	configs.DB.First(&book, id)

	body := BookRequestBody{}
	c.BindJSON(&body)
	data := &models.Book{
		Title:   body.Title,
		Author:  body.Author,
		Genre:   body.Genre,
		Summary: body.Summary,
	}

	result := configs.DB.Model(&book).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &book)
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	configs.DB.Delete(&book, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
