package main

import (
	"github.com/gin-gonic/gin"
)

func getAllExcuses(c *gin.Context) {
	c.JSON(200, gin.H{"excuses": Messages})
}

func addExcuse(c *gin.Context) {
	var newExcuse Message
	if err := c.ShouldBindJSON(&newExcuse); err != nil {
		c.JSON(400, gin.H{"error": "Bad request"})
		return
	}
	Messages = append(Messages, newExcuse)
	c.JSON(200, gin.H{"excuse": newExcuse})
}
