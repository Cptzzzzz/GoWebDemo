package controller

import (
	"Class_2_Homework/service"
	"github.com/gin-gonic/gin"
)

func AddTopic(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	id, err := service.NewTopic(title, content)
	resp := Response{}
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = -1
	} else {
		resp.Msg = "add successfully"
		resp.Code = 0
		resp.Data = map[string]uint{"topic_id": id}
	}
	c.JSON(200, resp)
}

func UpdateTopic(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	err := service.UpdateTopic(id, title, content)
	resp := Response{}
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Code = 0
		resp.Msg = "update successfully"
	}
	c.JSON(200, resp)
}
func DeleteTopic(c *gin.Context) {
	id := c.PostForm("id")
	err := service.DeleteTopic(id)
	resp := Response{}
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Code = 0
		resp.Msg = "delete successfully"
	}
	c.JSON(200, resp)
}
func GetTopic(c *gin.Context) {
	var err error
	resp := Response{}
	resp.Data, err = service.GetTopics()
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Code = 0
		resp.Msg = "get successfully"
	}
	c.JSON(200, resp)
}
