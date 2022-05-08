package service

import (
	"Class_2_Homework/repository"
	"errors"
	"strconv"
	"time"
)

func NewTopic(title, content string) (uint, error) {
	var err error
	topic := &repository.Topic{}
	topic.Id, err = repository.NewIdDaoInstance().NewTopicId()
	if err != nil {
		return 0, err
	}
	topic.Content = content
	topic.Title = title
	topic.CreateTime = time.Now()
	err = repository.NewTopicDaoInstance().CreateTopic(topic)
	if err != nil {
		return 0, err
	}
	return topic.Id, nil
}

func UpdateTopic(id, title, content string) error {
	var err error
	topic := &repository.Topic{}
	tId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = repository.NewTopicDaoInstance().CheckTopicById(uint(tId))
	if err != nil {
		return errors.New("topic record not found")
	}
	topic.Id = uint(tId)
	topic.Title = title
	topic.Content = content
	err = repository.NewTopicDaoInstance().UpdateTopic(topic)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTopic(id string) error {
	tId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = repository.NewTopicDaoInstance().CheckTopicById(uint(tId))
	if err != nil {
		return errors.New("topic record not found")
	}
	err = repository.NewTopicDaoInstance().DeletePostById(uint(tId))
	if err != nil {
		return err
	}
	return nil
}

func GetTopics() ([]repository.Topic, error) {
	res := make([]repository.Topic, 0)
	tmp, err := repository.NewTopicDaoInstance().QueryAllTopics()
	if err != nil {
		return nil, err
	}
	for _, v := range tmp {
		res = append(res, *v)
	}
	return res, nil
}
