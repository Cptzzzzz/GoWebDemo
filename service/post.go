package service

import (
	"Class_2_Homework/repository"
	"errors"
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
	tId, err := strconv.ParseInt(topicId, 10, 32)
	post.TopicId = uint(tId)
	err = repository.NewTopicDaoInstance().CheckTopicById(uint(tId))
	if err != nil {
		return 0, errors.New("topic record not found")
	}
	post.CreateTime = time.Now()
	err = repository.NewPostDaoInstance().CreatePost(post)
	if err != nil {
		return 0, err
	}
	return post.Id, nil
}

func UpdatePost(id, topicId, content string) error {
	var err error
	post := &repository.Post{}
	tId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = repository.NewPostDaoInstance().CheckPostById(uint(tId))
	if err != nil {
		return errors.New("post record not found")
	}
	post.Id = uint(tId)
	tId, err = strconv.ParseInt(topicId, 10, 64)
	if err != nil {
		return err
	}
	post.TopicId = uint(tId)
	err = repository.NewTopicDaoInstance().CheckTopicById(uint(tId))
	if err != nil {
		return errors.New("topic record not found")
	}
	post.Content = content
	err = repository.NewPostDaoInstance().UpdatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id string) error {
	tId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = repository.NewPostDaoInstance().CheckPostById(uint(tId))
	if err != nil {
		return errors.New("post record not found")
	}
	err = repository.NewPostDaoInstance().DeletePostById(uint(tId))
	if err != nil {
		return err
	}
	return nil
}

func GetPosts() ([]repository.Post, error) {
	res := make([]repository.Post, 0)
	tmp, err := repository.NewPostDaoInstance().QueryAllPosts()
	if err != nil {
		return nil, err
	}
	for _, v := range tmp {
		res = append(res, *v)
	}
	return res, nil
}
