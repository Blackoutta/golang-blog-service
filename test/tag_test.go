package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/Blackoutta/blog-service/global"
	v1 "github.com/Blackoutta/blog-service/internal/api/v1"
	"github.com/Blackoutta/blog-service/internal/service"
	"github.com/Blackoutta/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

func TestGet(t *testing.T) {
	type testcase struct {
		data string
		code int
		want string
	}

	type testsuite []testcase

	var ts testsuite = testsuite{
		{
			data: "1",
			code: 200,
			want: "success",
		},
		{
			data: "0",
			code: 400,
			want: "入参错误",
		},
		{
			data: "10000",
			code: 404,
			want: gorm.ErrRecordNotFound.Error(),
		},
	}

	global.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags)
	for _, tc := range ts {
		// 模拟http response writer 和 *gin.Context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 模拟请求和service
		tag := v1.NewTagAPI()
		tag.AddService(service.TagService{Ctx: c})
		tag.Service.AddDao(&mockTagDao{})
		req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/api/v1/tags/{id}", nil)
		if err != nil {
			t.Fatal(err)
		}
		c.Request = req
		c.Params = append(c.Params, gin.Param{
			Key:   "id",
			Value: tc.data,
		})

		// 执行方法
		tag.Get(c)

		// 断言http响应
		assert.Equal(t, tc.code, w.Result().StatusCode)
		assert.Contains(t, w.Body.String(), tc.want)
	}
}

func TestCreateTag(t *testing.T) {
	type testcase struct {
		data string
		code int
	}

	type testsuite []testcase

	var ts testsuite = testsuite{
		{
			data: "1",
			code: 200,
		},
		{
			data: "2",
			code: 200,
		},
		{
			data: "3",
			code: 400,
		},
	}

	global.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags)
	for _, tc := range ts {
		// 模拟http response writer 和 *gin.Context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 模拟请求和service
		tag := v1.NewTagAPI()
		tag.AddService(service.TagService{Ctx: c})
		tag.Service.AddDao(&mockTagDao{})

		body := fmt.Sprintf(`{
			"created_by": "some_user_name",
			"name": "some_tag_name",
			"state": %v }`, tc.data)
		rc := ioutil.NopCloser(strings.NewReader(body))
		req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/api/v1/tags", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		c.Request = req
		c.Request.Body = rc

		// 执行方法
		tag.Create(c)

		// 断言http响应
		assert.Equal(t, tc.code, w.Result().StatusCode)
	}
}
