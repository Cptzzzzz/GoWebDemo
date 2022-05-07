package repository

import "sync"

type Id struct {
	PostId  uint
	TopicId uint
}

func (Id) TableName() string {
	return "id"
}

type IdDao struct {
}

var idDao *IdDao
var idOnce sync.Once

func NewIdDaoInstance() *IdDao {
	idOnce.Do(
		func() {
			idDao = &IdDao{}
		})
	return idDao
}
func (*IdDao) NewPostId() (uint, error) {
	var id Id
	db.First(&id)
	res := id.PostId + 1
	err := db.Model(&id).Where("post_id = ?", id.PostId).Update("post_id", id.PostId+1).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (*IdDao) NewTopicId() (uint, error) {
	var id Id
	db.First(&id)
	res := id.TopicId + 1
	err := db.Model(&id).Where("post_id = ?", id.PostId).Update("topic_id", id.TopicId+1).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}
