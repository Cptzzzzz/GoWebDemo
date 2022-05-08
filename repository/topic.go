package repository

import (
	"sync"
	"time"
)

type Topic struct {
	Id         uint
	Title      string
	Content    string
	CreateTime time.Time
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao struct {
}

var topicDao *TopicDao
var topicOnce sync.Once

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id uint) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (*TopicDao) CreateTopic(topic *Topic) error {
	err := db.Create(topic).Error
	if err != nil {
		return err
	}
	return nil
}

func (*TopicDao) QueryAllTopics() ([]*Topic, error) {
	var res []*Topic
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*TopicDao) DeletePostById(id uint) error {
	err := db.Where("id = ?", id).Delete(&Topic{}).Error
	return err
}

func (*TopicDao) UpdateTopic(topic *Topic) error {
	err := db.Where("id = ?", topic.Id).Updates(topic).Error
	return err
}

func (*TopicDao) CheckTopicById(id uint) error {
	return db.Where("id = ?", id).First(&Topic{}).Error
}
