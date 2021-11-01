package download

import (
	"github.com/gin-gonic/gin"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/models/mysql/download"
)

func GetDownload(c *gin.Context) {

	var receiver download.Download
	result, err := receiver.QueryInfo()
	if err != nil {
		i18nresponse.Error(c, "1010011")
		return
	}

	i18nresponse.Success(c, "ok", result)
}
