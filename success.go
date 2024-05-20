package goframework_resp_i18n

import "github.com/kordar/goframework_resp"

// ---------- i18n success --------------

type SuccessResultI18n struct {
	I18nMessage func(message string, t string, c interface{}) string
	goframework_resp.SuccessResult
}

func (s SuccessResultI18n) Result(c interface{}, message interface{}, data interface{}, count int64) {
	if value, ok := message.(string); ok && value != "" {
		if tmessage := s.I18nMessage(value, "success", c); tmessage != "" {
			goframework_resp.GetResultCallFunc()(
				c,
				s.HttpStatus(),
				goframework_resp.Code("success"),
				tmessage,
				data,
				count,
			)
			return
		}
	}

	goframework_resp.GetResultCallFunc()(
		c,
		s.HttpStatus(),
		goframework_resp.Code("success"),
		s.I18nMessage("success", "success", c),
		data,
		count,
	)
}
