package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

func (v *ValidError) Error() string {
	return v.Message
}

type ValidErrors []*ValidError

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors

	// 将表单提交的数据与Request进行绑定
	err := c.ShouldBind(v)

	// 如果bind出现错误
	if err != nil {
		// 从context中取出中间件：翻译器
		translator := c.Value("trans")
		// 取出来的中间件翻译器是interface{}类型，这里将其转换为ut.Translator
		trans, _ := translator.(ut.Translator)

		// 判断刚才出现的错误是否是ValidationErrors类型
		verrs, ok := err.(val.ValidationErrors)
		// 如果不是
		if !ok {
			errs = append(errs, &ValidError{
				Key:     "其他错误",
				Message: "请求参数错误或请求体格式错误，详细：" + err.Error(),
			})
			// 返回true代表有错误，但错误不是由于参数校验引起的，即错误是其他原因引起(如json格式不对)
			return true, errs
		}

		// 如果是, 则翻译每个错误并把错误加到errs slice
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		// 返回true代表有错误，并返回errs slice
		return true, errs
	}
	// 返回false代表无错误
	return false, nil
}
