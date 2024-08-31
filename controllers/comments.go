package controllers

import (
	"comments_service/models"

	"github.com/gin-gonic/gin"
)

var Comments []models.Comment

func GetComments(c *gin.Context) {
	c.JSON(200, gin.H{"success gett": Comments})
}
func PostComment(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Comments = append(Comments, comment)
	c.JSON(200, gin.H{"Success": "Comment posted"})
}
