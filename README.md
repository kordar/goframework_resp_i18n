# goframework_resp_i18n

对[`goframework_resp`](https://github.com/kordar/goframework_resp)的多语言实现版本。

## 安装

```go
go get https://github.com/kordar/goframework_resp_i18n v0.0.1
```

## 使用

将`goframework_resp`中的类型替换成`i18n`实现类

```go
response.RegRespFunc(response.SuccessType, responseI18n.SuccessResultI18n{I18nMessage: i18nFunc})
response.RegRespFunc(response.ErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18nFunc, GetTrans: gettrans})
response.RegRespFunc(response.ValidErrorType, responseI18n.ErrorResultI18n{I18nMessage: i18nFunc, GetTrans: gettrans})
response.RegRespFunc(response.OutputType, responseI18n.OutputResponseI18n{I18nMessage: i18nFunc})
response.RegRespFunc(response.UnauthorizedType, responseI18n.UnauthorizedJsonI18n{I18nMessage: i18nFunc})
```

- `gin`实现

```go
var i18nFunc = func(message string, messagetype string, c interface{}) string {
    ctx := c.(*gin.Context)
    locale := getlocale(ctx)
    if messagetype == response.SuccessType {
		// TODO 使用gocfg实现多语言管理
        return gocfg.GetSectionValue(locale, fmt.Sprintf("response.success.%s", message), "language")
    } else if messagetype == response.ErrorType {
        return gocfg.GetSectionValue(locale, fmt.Sprintf("response.errors.%s", message), "language")
    } else {
        return gocfg.GetSectionValue(locale, fmt.Sprintf("response.common.%s", message), "language")
    }
}

// 获取翻译器
func gettrans(c interface{}) (trans ut.Translator, found bool) {
    ctx := c.(*gin.Context)
    locale := getlocale(ctx)
    return gotrans.Get().GetTranslator(GetRealLocale(locale))
}
```
