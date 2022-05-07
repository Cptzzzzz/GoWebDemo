package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddTopic(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	fmt.Println(title, content)
	c.JSON(200, Response{})
}

func UpdateTopic(c *gin.Context) {
	c.JSON(200, Response{})
}
func DeleteTopic(c *gin.Context) {
	c.JSON(200, Response{})
}
func GetTopic(c *gin.Context) {
	c.JSON(200, Response{})
}
