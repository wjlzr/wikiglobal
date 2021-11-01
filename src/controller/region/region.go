package region

import (
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/models/mysql/region"

	"github.com/gin-gonic/gin"
)

// GetInfo 获取地区信息
func GetInfo(c *gin.Context) {

	result, err := region.Region{}.GetInfo()
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}
	i18nresponse.Success(c, "ok", result)
}
