package service

import (
	"github.com/Blackoutta/blog-service/internal/dao"
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/Blackoutta/blog-service/pkg/app"
	"golang.org/x/net/context"
)

// requests
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1 2"`
}

type GetTagRequest struct {
	Id uint32 `form:"id" binding:"gt=0"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=10"`
	State uint8  `form:"state,default=0" binding:"oneof=0 1 2"`
}

type CreateTagRequest struct {
	Name      string `json:"name" example:"some_tag_name" binding:"required,min=3,max=100"`
	CreatedBy string `json:"created_by" example:"some_user_name" binding:"required,min=3,max=100"`
	State     uint8  `json:"state,default=1" example:"1" binding:"oneof=1 2"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id"`
	Name       string `json:"name" binding:"required,min=3,max=100"`
	State      *uint8 `json:"state" binding:"required,oneof=1 2"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type UpdateTagRequestSwagger struct {
	Name       string `json:"name"`
	State      *uint8 `json:"state"`
	ModifiedBy string `json:"modified_by"`
}

type ChangeStateRequest struct {
	ID         uint32 `form:"id"`
	State      *uint8 `json:"state" binding:"required,oneof=1 2"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type ChangeStateRequestSwagger struct {
	State      *uint8
	ModifiedBy string
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"gt=0"`
}

type TagService struct {
	Ctx context.Context
	Dao dao.TagDao
}

func (svc *TagService) AddDao(dao dao.TagDao) {
	svc.Dao = dao
}

// service
func (svc *TagService) CountTag(param *CountTagRequest) (int, error) {
	return svc.Dao.CountTag(param.Name, param.State)
}

func (svc *TagService) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return svc.Dao.GetTag(param.Id)
}

func (svc *TagService) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.Dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *TagService) CreateTag(param *CreateTagRequest) error {
	return svc.Dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *TagService) UpdateTag(param *UpdateTagRequest) error {
	return svc.Dao.UpdateTag(param.ID, param.Name, *param.State, param.ModifiedBy)
}

func (svc *TagService) ChangeState(param *ChangeStateRequest) error {
	return svc.Dao.ChangeState(param.ID, *param.State, param.ModifiedBy)
}

func (svc *TagService) DeleteTag(param *DeleteTagRequest) error {
	return svc.Dao.DeleteTag(param.ID)
}
