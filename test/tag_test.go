package test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Blackoutta/blog-service/global"
	v1 "github.com/Blackoutta/blog-service/internal/api/v1"
	"github.com/Blackoutta/blog-service/internal/model"
	"github.com/Blackoutta/blog-service/internal/service"
	"github.com/Blackoutta/blog-service/pkg/app"
	"github.com/Blackoutta/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type testcase struct {
	code int
	want gin.H
}

type testsuite []testcase

var ts testsuite = testsuite{
	testcase{
		code: 200,
		want: gin.H{
			"name": "success",
			"state": 
		},
	},
}

func TestGet(t *testing.T) {
	global.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	tag := v1.NewTag()
	mockSvc := &mockTagService{}
	tag.BuildService(mockSvc)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/tags/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	c.Request = req
	c.Params = append(c.Params, gin.Param{
		Key:   "id",
		Value: "0",
	})

	tag.Get(c)
	assert.Equal(t, 200, w.Result().StatusCode)

	var got gin.H
	err = json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	want := "success"
	assert.Equal(t, want, got["name"])
}

type mockTagService struct{}

func (svc *mockTagService) CountTag(param *service.CountTagRequest) (int, error) {
	return 1, nil
}

func (svc *mockTagService) GetTag(param *service.GetTagRequest) (*model.Tag, error) {
	return &model.Tag{
		Name: "success",
		State: 1,
	}, nil
}

func (svc *mockTagService) GetTagList(param *service.TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return []*model.Tag{
		&model.Tag{
			Name: "success",
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
