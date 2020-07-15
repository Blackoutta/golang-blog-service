package dao

import (
	"github.com/Blackoutta/blog-service/global"
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/Blackoutta/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type TagDao interface {
	CountTag(name string, state uint8) (int, error)
	GetTag(id uint32) (*model.Tag, error)
	GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error)
	CreateTag(name string, state uint8, createdBy string) error
	UpdateTag(id uint32, name string, state uint8, modifiedBy string) error
	ChangeState(id uint32, state uint8, modifiedBy string) error
	DeleteTag(id uint32) error
}

type tagDao struct {
	engine *gorm.DB
}

func NewTagDao(engine *gorm.DB) TagDao {
	return &tagDao{
		engine: global.DBEngine,
	}
}

func (d *tagDao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *tagDao) GetTag(id uint32) (*model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Get(d.engine)
}

func (d *tagDao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *tagDao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return tag.Create(d.engine)
}

func (d *tagDao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
		"name":        name,
	}

	return tag.Update(d.engine, values)
}

func (d *tagDao) ChangeState(id uint32, state uint8, modifiedBy string) error {
	tag := model.Tag{
		State: state,
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}
	return tag.ChangeState(d.engine)

}

func (d *tagDao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	return tag.Delete(d.engine)
}
