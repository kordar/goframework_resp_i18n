package goframework_resp_i18n

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/goframework_resp"
)

type ErrorResultI18n struct {
	GetTrans    func(c interface{}) (trans ut.Translator, found bool)
	I18nMessage func(message string, t string, c interface{}) string
	goframework_resp.ErrorResult
}

func (e ErrorResultI18n) Result(c interface{}, message interface{}, data interface{}, count int64) {

	// 处理文本 message
	if err, ok := message.(string); ok && err != "" {
		if tmessage := e.I18nMessage(err, "error", c); tmessage != "" {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("error"), tmessage, data, count)
		} else {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("error"), err, data, count)
		}
		return
	}

	// 处理 validate error
	if err, ok := message.(validator.ValidationErrors); ok {
		nMessage := e.I18nMessage("error.valid", "error", c)
		if trans, found := e.GetTrans(c); found {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("valid"), nMessage, err.Translate(trans), count)
		} else {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("valid"), nMessage, data, count)
		}
		return
	}

	// 处理 error
	if value, ok := message.(error); ok {
		if tmessage := e.I18nMessage(value.Error(), "error", c); tmessage != "" {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("error"), tmessage, data, count)
		} else {
			goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("error"), value.Error(), data, count)
		}

		return
	}

	goframework_resp.GetResultCallFunc()(c, e.HttpStatus(), goframework_resp.Code("error"), e.I18nMessage("error", "error", c), data, count)
}
