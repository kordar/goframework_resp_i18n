package goframework_resp_i18n

import "github.com/kordar/goframework_resp"

type OutputResponseI18n struct {
	I18nMessage func(message string, t string, c interface{}) string
	goframework_resp.OutputResponse
}

func (o OutputResponseI18n) Result(c interface{}, message interface{}, data interface{}, count int64) {
	// TODO implement me
	/*
		if value, ok := data.(IOutput); ok && value != nil {
			if value.Type() == "browser" {
				// output web
				for k, v := range value.Header() {
					c.Header(k, v)
				}
				_, _ = c.Writer.Write(value.Data())
				return
			}

			msg := "output.success"
			if tmessage, found := message.(string); found && tmessage != "" {
				msg = tmessage
				if tmessage2 := o.I18nMessage(tmessage, "success", c); tmessage2 != "" {
					msg = tmessage2
				}
			}

			c.JSON(o.HttpStatus(), map[string]interface{}{
				"code": success, "message": o.I18nMessage(msg, "success", c), "data": value.Data(), "params": value.Params(),
			})

			return
		}
	*/

	goframework_resp.GetResultCallFunc()(c, o.HttpStatus(), goframework_resp.Code("fail"), o.I18nMessage("output.fail", "error", c), nil, -1)

}
