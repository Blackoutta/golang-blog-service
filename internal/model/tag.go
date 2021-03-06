package model

import (
	"github.com/Blackoutta/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (a Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

const ShowAllStates uint8 = 2

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	// 如果用户查询时输入了name, 根据name进行筛选
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	// 根据用户输入的state进行筛选
	if t.State != 0 {
		db = db.Where("state = ?", t.State)
	}

	// 使用model对数据库进行查询并将结果写入到count中
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag Tag
	err := db.Where("is_del = ?", 0).Where("id = ?", t.Model.ID).Find(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	if t.State != 0 {
		db = db.Where("state = ?", t.State)
	}

	if err = db.Where("is_del = ?", 0).Order("id DESC").Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) ChangeState(db *gorm.DB) error {
	err := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
