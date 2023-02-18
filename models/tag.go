package models

import (
	"time"
	"gorm.io/gorm"
)

type AddTagFrom struct {
    Name string `form:"name" json:"name" binding:"required,gte=1,lte=130"`
    CreatedBy string `form:"created_by" json:"created_by" binding:"required,gte=0,lte=130"`
    State int `form:"state" json:"state" binding:"required,gte=0,lte=1"`
}

type GetTagsForm struct {
    Name string `query:"name" form:"name" json:"name" binding:"gte=0,lte=130"`
    State int `query:"state" form:"state" json:"state" binding:"gte=0,lte=1"`
}

type EditTagForm struct {
    Name string `query:"name" form:"name" json:"name" binding:"gte=0,lte=130"`
    ModifiedBy string `query:"modified_by" form:"modified_by" json:"modified_by" binding:"required,gte=0,lte=130"`
}

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
    
    return
}

func GetTagTotal(maps interface {}) (count int64){
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

func ExistTagById(id int) bool {
    var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
    if tag.ID > 0 {
        return true
    }

    return false
}

func AddTag(name string, state int, createdBy string) bool {
    db.Create(&Tag {
        Name : name,
        State : state,
        CreatedBy : createdBy,
    })

    return true
}

func EditTag(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Model(&Tag{}).Where("id = ?", id).Delete(&Tag{})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.DB) error {
    tag.CreatedOn = time.Now()
	tag.ModifiedOn = time.Now()

    return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.DB) error {
    tag.ModifiedOn = time.Now()

    return nil
}