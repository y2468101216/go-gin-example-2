package models

import (
	"time"
	"gorm.io/gorm"
)

type GetArticlesForm struct {
	TagID int `query:"tag_id" form:"tag_id" json:"tag_id" binding:"gte=1"`
	State int `query:"state" form:"state" json:"state" binding:"gte=0,lte=1"`
}

type AddArticleForm struct {
	TagID int `query:"tag_id" form:"tag_id" json:"tag_id" binding:"required,gte=1"`
	Title string `query:"title" form:"title" json:"title" binding:"required,gte=1"`
	Desc string `query:"desc" form:"desc" json:"desc" binding:"required,gte=1"`
	Content string `query:"content" form:"content" json:"content" binding:"required,gte=1"`
	CreatedBy string `query:"created_by" form:"created_by" json:"created_by" binding:"required,gte=1"`
	State int `query:"state" form:"state" json:"state" binding:"required,gte=0,lte=1"`
}

type EditArticleForm struct {
	TagID int `query:"tag_id" form:"tag_id" json:"tag_id" binding:"gte=1"`
	Title string `query:"title" form:"title" json:"title" binding:"gte=1"`
	Desc string `query:"desc" form:"desc" json:"desc" binding:"gte=1"`
	Content string `query:"content" form:"content" json:"content" binding:"gte=1"`
	ModifiedBy string `query:"modified_by" form:"modified_by" json:"modified_by" binding:"gte=1"`
	State int `query:"state" form:"state" json:"state" binding:"gte=0,lte=1"`
}

type Article struct {
    Model

    TagID int `json:"tag_id" gorm:"index"`
    Tag   Tag `json:"tag"`

    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
	DeletedOn gorm.DeletedAt `json:"deleted_on"`
}

func ExistArticleByID(id int) bool {
    var article Article
    db.Select("id").Where("id = ?", id).First(&article)

    if article.ID > 0 {
        return true
    }

    return false
}

func GetArticleTotal(maps interface {}) (count int64){
    db.Model(&Article{}).Where(maps).Count(&count)

    return
}

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

    return
}

func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).Preload("Tag").First(&article)

    return 
}

func EditArticle(id int, data interface {}) bool {
    db.Model(&Article{}).Where("id = ?", id).Updates(data)

    return true
}

func AddArticle(data AddArticleForm) bool {
    db.Create(&Article {
        TagID : data.TagID,
        Title : data.Title,
        Desc : data.Desc,
        Content : data.Content,
        CreatedBy : data.CreatedBy,
        State : data.State,
    })

    return true
}

func DeleteArticle(id int) bool {
    db.Where("id = ?", id).Delete(&Article{})

    return true
}

func (article *Article) BeforeCreate(scope *gorm.DB) error {
    article.CreatedOn = time.Now()
	article.ModifiedOn = time.Now()

    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.DB) error {
    article.ModifiedOn = time.Now()

    return nil
}

func CleanAllArticle() bool {
    db.Unscoped().Where("deleted_on is not null").Delete(&Article{})

    return true
}