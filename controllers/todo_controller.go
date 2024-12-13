package controllers

import (
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos) // Use config.DB to access the database
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Description: input.Description, Completed: input.Completed}
	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	parmID := c.Param("id")
	if err := config.DB.First(&todo, parmID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&todo).Updates(input) // Use config.DB to update the record
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := config.DB.First(&todo, c.Param("id")).Error; err != nil { // Use config.DB to find the record
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	config.DB.Delete(&todo) // Use config.DB to delete the record
	c.JSON(http.StatusOK, gin.H{"data": "Todo deleted"})
}
