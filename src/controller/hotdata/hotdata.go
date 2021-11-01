package hotdata

import (
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/global/response"
	"wiki_global/src/models/mysql/hotdata"

	"github.com/gin-gonic/gin"
)

// GetList 获取热门公司或者高管
func GetList(c *gin.Context) {

	var hotDataRequest hotdata.HotDataRequest
	if err := c.ShouldBindJSON(&hotDataRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	result, count, err := hotDataRequest.QueryList(c)
	if err != nil {
		i18nresponse.Error(c, "1010013")
		return
	}

	response.PageOK(c, result, count, hotDataRequest.PageIndex, hotDataRequest.PageSize, "ok")
}
