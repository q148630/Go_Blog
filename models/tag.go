package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

// gorm (hook 函數) 在建立前調用，實現 BeforeCreateInterface
func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	tag.CreatedOn = int(time.Now().Unix())
	return nil
}

// gorm (hook 函數) 在更新前調用，實現 BeforeUpdateInterface
func (tag *Tag) BeforeUpdate(tx *gorm.DB) error {
	tag.ModifiedOn = int(time.Now().Unix())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagByID(id int) *Tag {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return &tag
	}
	return nil
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data *Tag) bool {
	db.Where("id = ?", id).Updates(data)
	return true
}