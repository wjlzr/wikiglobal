package share

import (
	"github.com/gin-gonic/gin"
	"wiki_global/src/common/convert"
	"wiki_global/src/global/i18nresponse"
)

// GetUrl 获取链接
func GetUrl(c *gin.Context) {

	type1 := convert.StrToInt(c.Request.FormValue("type"))
	shareUrl := ""
	if type1 == 0 {
		i18nresponse.Error(c, "1010004")
		return
	}
	if type1 == 1 {
		shareUrl = "https://www.wikiglobal.com/cn/company/"
	} else {
		shareUrl = "https://www.wikiglobal.com/cn/executive/"
	}
	i18nresponse.Success(c, "ok", struct {
		ShareUrl string `json:"share_url"`
	}{ShareUrl: shareUrl})
}
