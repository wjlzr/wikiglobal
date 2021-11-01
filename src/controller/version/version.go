package version

import (
	"github.com/gin-gonic/gin"
	"wiki_global/src/global/i18nresponse"
	version2 "wiki_global/src/models/mysql/version"
)

// Version 版本详情
func Version(c *gin.Context) {

	client := c.Request.FormValue("client")
	if client == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	var v version2.Version
	v.Client = client
	result, err := v.FindOne()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", result)
}
