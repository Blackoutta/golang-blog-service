package service

import (
	"github.com/Blackoutta/blog-service/global"
	"github.com/Blackoutta/blog-service/internal/dao"
	"golang.org/x/net/context"
)

type tagService struct {
	ctx context.Context
	dao *dao.Dao
}

func NewTagService(ctx context.Context) TagService {
	svc := tagService{
		ctx: ctx,
	}
	svc.dao = dao.New(global.DBEngine)
	return &svc
}
