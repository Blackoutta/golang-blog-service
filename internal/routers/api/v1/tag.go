package v1

import (
	"github.com/Blackoutta/blog-service/global"
	"github.com/Blackoutta/blog-service/internal/service"
	"github.com/Blackoutta/blog-service/pkg/app"
	"github.com/Blackoutta/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// Produce json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "标签名称" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failture 400 {object} errcode.Error "请求错误"
// @Failture 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	invalid, errs := app.BindAndValid(c, &param)
	if invalid == true {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
