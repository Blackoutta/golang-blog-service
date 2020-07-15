package test

import (
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/Blackoutta/blog-service/internal/service"
	"github.com/Blackoutta/blog-service/pkg/app"
)

type mockTagService struct{}

func (svc *mockTagService) CountTag(param *service.CountTagRequest) (int, error) {
	return 1, nil
}

func (svc *mockTagService) GetTag(param *service.GetTagRequest) (*model.Tag, error) {
	return &model.Tag{
		Name:  "success",
		State: 1,
	}, nil
}

func (svc *mockTagService) GetTagList(param *service.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return []*model.Tag{
		{
			Name:  "success",
			State: 1,
		},
	}, nil
}

func (svc *mockTagService) CreateTag(param *service.CreateTagRequest) error {
	return nil
}

func (svc *mockTagService) UpdateTag(param *service.UpdateTagRequest) error {
	return nil
}

func (svc *mockTagService) ChangeState(param *service.ChangeStateRequest) error {
	return nil
}

func (svc *mockTagService) DeleteTag(param *service.DeleteTagRequest) error {
	return nil
}
