package service

import (
	"github.com/Blackoutta/blog-service/internal/dao"
	"golang.org/x/net/context"
)

type tagService struct {
	ctx context.Context
	dao dao.TagDao
}

func NewTagService(ctx context.Context, dao dao.TagDao) TagService {
	svc := tagService{
		ctx: ctx,
	}
	svc.dao = dao
	return &svc
}
