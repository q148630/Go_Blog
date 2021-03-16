package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedBy   string `json:"created_by"`
	ModifiedBy  string `json:"modified_by"`
	State       int    `json:"state"`
}

func (article *Article) BeforeCreate(tx *gorm.DB) error {
	article.CreatedOn = int(time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(tx *gorm.DB) error {
	article.ModifiedOn = int(time.Now().Unix())
	return nil
}

func ExistArticleByID(id int) *Article {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return &article
	}
	return nil
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return 
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Preload("Tag").Where("id = ?", id).First(&article)
	return
}

func EditArticle(id int, data *Article) bool {
	db.Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:       data["tag_id"].(int),
		Title:       data["title"].(string),
		Description: data["desc"].(string),
		Content:     data["content"].(string),
		CreatedBy:   data["created_by"].(string),
		State:       data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}