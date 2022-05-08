package main

import (
	"Class_2_Homework/controller"
	"Class_2_Homework/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	v := r.Group("/api")
	{
		v1 := v.Group("/topic")
		{
			v1.POST("/add", controller.AddTopic)
			v1.POST("/update", controller.UpdateTopic)
			v1.POST("/delete", controller.DeleteTopic)
			v1.GET("/get", controller.GetTopic)
		}
		v2 := v.Group("/post")
		{
			v2.POST("/add", controller.AddPost)
			v2.POST("/update", controller.UpdatePost)
			v2.POST("/delete", controller.DeletePost)
			v2.GET("/get", controller.GetPost)
		}
	}
	err := r.Run(":80")
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}
