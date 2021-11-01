package utils

import (
	"net/http"
	"wiki_global/src/common/constant/lang"

	"github.com/gin-gonic/gin"
)

//错误返回的json
func ResponseErrorJson(c *gin.Context, code int64, lang string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  lang, //service.Errors(code, lang),
	})
	c.Abort()
}

//正确返回的json
func ResponseJson(c *gin.Context, msg, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

//
func ResponseErrorFormatJson(c *gin.Context, code int64, lang string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "", //fmt.Sprintf(service.Errors(code, lang), v...),
	})
}

//获取语言
func GetLang(c *gin.Context) string {
	language := c.Request.Header.Get("Language")
	if len(language) > 0 {
		return language
	}
	return lang.ZH_CN
}

//获取版本号
func GetVersion(c *gin.Context) string {
	version := c.Request.Header.Get("version")
	if len(version) > 0 {
		return version
	}
	return "1.0.1"
}
