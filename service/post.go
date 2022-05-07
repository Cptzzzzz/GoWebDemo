package service

import (
	"Class_2_Homework/repository"
	"strconv"
	"time"
)

func NewPost(topicId, content string) (uint, error) {
	var err error
	post := &repository.Post{}
	post.Id, err = repository.NewIdDaoInstance().NewPostId()
	if err != nil {
		return 0, err
	}
	post.Content = content
	toId, err := strconv.ParseInt(topicId, 10, 32)
	post.TopicId = uint(toId)
	post.CreateTime = time.Now()
	err = repository.NewPostDaoInstance().CreatePost(post)
	if err != nil {
		return 0, err
	}
	return post.Id, nil
}
