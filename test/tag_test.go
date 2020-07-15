package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Blackoutta/blog-service/global"
	v1 "github.com/Blackoutta/blog-service/internal/api/v1"
	"github.com/Blackoutta/blog-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

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
}

func TestGet(t *testing.T) {
	global.Logger = logger.NewLogger(os.Stdout, "", log.LstdFlags)
	for _, tc := range ts {
		// 模拟http response writer 和 *gin.Context
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// 模拟请求和service
		tag := v1.NewTag()
		mockSvc := &mockTagService{}
		tag.BuildService(mockSvc)
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
