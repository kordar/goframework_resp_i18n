package goframework_resp_i18n

import "github.com/kordar/goframework_resp"

type UnauthorizedJsonI18n struct {
	I18nMessage func(message string, t string, c interface{}) string
	goframework_resp.UnauthorizedResponse
}

func (s UnauthorizedJsonI18n) Result(c interface{}, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		if tmessage := s.I18nMessage(value, "unauthorized", c); tmessage != "" {
			goframework_resp.GetResultCallFunc()(c, s.HttpStatus(), goframework_resp.Code("unauthorized"), tmessage, data, count)
			return
		}
	}

	goframework_resp.GetResultCallFunc()(c, s.HttpStatus(), goframework_resp.Code("unauthorized"), s.I18nMessage("unauthorized", "unauthorized", c), data, count)
}
