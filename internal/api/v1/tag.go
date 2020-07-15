package v1

import (
	"github.com/Blackoutta/blog-service/global"
	"github.com/Blackoutta/blog-service/internal/dao"
	"github.com/Blackoutta/blog-service/internal/service"
	"github.com/Blackoutta/blog-service/pkg/app"
	"github.com/Blackoutta/blog-service/pkg/convert"
	"github.com/Blackoutta/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Service service.TagService
}

func NewTag() Tag {
	return Tag{}
}

func (t *Tag) BuildService(svc service.TagService) {
	t.Service = svc
}

// @Summary 获取单个标签
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
	param := service.GetTagRequest{
		Id: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)
	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if t.Service == nil {
		t.BuildService(service.NewTagService(c, dao.NewTagDao(global.DBEngine)))
	}

	tag, err := t.Service.GetTag(&param)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
			return
		}
		global.Logger.Errorf("svc.GetTag errs: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(&tag)
	return
}

// @Summary 获取多个标签
// @Produce json
// @Param name query string false "标签名称, 最大长度100"
// @Param state query int false "标签状态, 0:全部  1:启用  2:禁用"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	// 初始化请求体和响应体
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	// 将用户的请求与请求体进行绑定
	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewTagService(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	totalRows, err := svc.CountTag(&service.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary 创建标签
// @Produce json
// @Param CreateTagRequest body service.CreateTagRequest true "创建标签json请求体"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)

	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.NewTagService(c)
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签状态
// @Produce json
// @Param id path int true "标签id"
// @Param ChangeStateRequest body service.ChangeStateRequestSwagger true "更新标签状态json请求体"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags/state/{id} [patch]
func (t Tag) ChangeState(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUint32()
	param := service.ChangeStateRequest{}
	response := app.NewResponse(c)

	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	param.ID = id
	svc := service.NewTagService(c)
	err := svc.ChangeState(&param)
	if err != nil {
		global.Logger.Errorf("svc.ChangeState err:", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 更新标签
// @Produce json
// @Param id path int true "标签id"
// @Param UpdateTagRequest body service.UpdateTagRequestSwagger true "更新标签结构体"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)

	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	param.ID = convert.StrTo(c.Param("id")).MustUint32()

	svc := service.NewTagService(c)
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err:", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部服务错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUint32(),
	}
	response := app.NewResponse(c)

	invalid, errs := app.BindAndValid(c, &param)
	if invalid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	param.ID = convert.StrTo(c.Param("id")).MustUint32()

	svc := service.NewTagService(c)
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err:", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
