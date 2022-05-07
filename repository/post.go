package repository

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Post struct {
	Id         uint
	TopicId    uint
	Content    string
	CreateTime time.Time
}

func (Post) TableName() string {
	return "post"
}

type PostDao struct {
}

var postDao *PostDao
var postOnce sync.Once

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id uint) (*Post, error) {
	var post Post
	err := db.Where("id = ?", id).Find(&post).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (*PostDao) QueryPostByTopic(topicId uint) ([]*Post, error) {
	var posts []*Post
	err := db.Where("topic_id = ?", topicId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (*PostDao) CreatePost(post *Post) error {
	err := db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (*PostDao) QueryAllPosts() ([]*Post, error) {
	var res []*Post
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*PostDao) DeletePostById(id uint) error {
	err := db.Where("id = ?", id).Delete(&Post{}).Error
	return err
}

func (*PostDao) UpdatePostById(post *Post) error {
	err := db.Where("id = ?", post.Id).Updates(post).Error
	return err
}
