package service

import (
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/Blackoutta/blog-service/pkg/app"
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

type TagService interface {
	CountTag(param *CountTagRequest) (int, error)
	GetTag(param *GetTagRequest) (*model.Tag, error)
	GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error)
	CreateTag(param *CreateTagRequest) error
	UpdateTag(param *UpdateTagRequest) error
	ChangeState(param *ChangeStateRequest) error
	DeleteTag(param *DeleteTagRequest) error
}

// service
func (svc *tagService) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *tagService) GetTag(param *GetTagRequest) (*model.Tag, error) {
	return svc.dao.GetTag(param.Id)
}

func (svc *tagService) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *tagService) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *tagService) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, *param.State, param.ModifiedBy)
}

func (svc *tagService) ChangeState(param *ChangeStateRequest) error {
	return svc.dao.ChangeState(param.ID, *param.State, param.ModifiedBy)
}

func (svc *tagService) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
