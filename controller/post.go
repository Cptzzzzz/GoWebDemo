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
	resp := Response{Data: map[string]uint{"post_id": id}}
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = -1
	} else {
		resp.Code = 0
		resp.Msg = "发布成功"
	}
	c.JSON(200, resp)
}

func UpdatePost(c *gin.Context) {

}
func DeletePost(c *gin.Context) {
	c.JSON(200, Response{})
}
func GetPost(c *gin.Context) {
	c.JSON(200, Response{})
}
