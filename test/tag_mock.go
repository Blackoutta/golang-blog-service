package test

import (
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/jinzhu/gorm"
)

type mockTagDao struct{}

func (d *mockTagDao) CountTag(name string, state uint8) (int, error) {
	return 1, nil
}

func (d *mockTagDao) GetTag(id uint32) (*model.Tag, error) {
	if id > 9999 {
		return &model.Tag{}, gorm.ErrRecordNotFound
	}
	return &model.Tag{
		Name: "success",
	}, nil
}

func (d *mockTagDao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	return []*model.Tag{
		{
			Name: "success",
		},
	}, nil
}

func (d *mockTagDao) CreateTag(name string, state uint8, createdBy string) error {
	return nil
}

func (d *mockTagDao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	return nil
}

func (d *mockTagDao) ChangeState(id uint32, state uint8, modifiedBy string) error {
	return nil
}

func (d *mockTagDao) DeleteTag(id uint32) error {
	return nil
}
