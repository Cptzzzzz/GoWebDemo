package controller

import (
	"Class_2_Homework/service"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func AddPost(c *gin.Context) {
	topicId := c.PostForm("topic_id")
	content := c.PostForm("content")
	id, err := service.NewPost(topicId, content)
	resp := Response{}
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = -1
	} else {
		resp.Code = 0
		resp.Msg = "publish successfully"
		resp.Data = map[string]uint{"post_id": id}
	}
	c.JSON(200, resp)
}

func UpdatePost(c *gin.Context) {
	id := c.PostForm("id")
	topicId := c.PostForm("topic_id")
	content := c.PostForm("content")
	err := service.UpdatePost(id, topicId, content)
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
func DeletePost(c *gin.Context) {
	id := c.PostForm("id")
	err := service.DeletePost(id)
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
func GetPost(c *gin.Context) {
	var err error
	resp := Response{}
	resp.Data, err = service.GetPosts()
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	} else {
		resp.Code = 0
		resp.Msg = "get successfully"
	}
	c.JSON(200, resp)
}
