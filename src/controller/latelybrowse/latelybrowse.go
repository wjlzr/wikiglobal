package latelybrowse

import (
	"wiki_global/src/common/convert"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/global/response"
	"wiki_global/src/models/mysql/latelybrowse"
	"wiki_global/src/utils"

	"github.com/gin-gonic/gin"
)

// GetList 获取最近浏览的企业或者高管
func GetList(c *gin.Context) {

	var latelyBrowse latelybrowse.LatelyBrowseRequest
	if err := c.ShouldBindJSON(&latelyBrowse); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	//latelyBrowse.UserId = convert.Int64ToString(utils.GetUserId(c))
	result, count, err := latelyBrowse.QueryList(c)
	if err != nil {
		i18nresponse.Error(c, "1010013")
		return
	}

	response.PageOK(c, result, count, latelyBrowse.PageIndex, latelyBrowse.PageSize, "ok")
}

// 添加最近浏览的公司/高管
func Add(c *gin.Context) {

	var latelyBrowse latelybrowse.LatelyBrowse
	if err := c.ShouldBindJSON(&latelyBrowse); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	latelyBrowse.UserId = convert.Int64ToString(utils.GetUserId(c))
	if _, err := latelyBrowse.Insert(); err != nil {
		i18nresponse.Error(c, "1010014")
		return
	}

	i18nresponse.Success(c, "ok", nil)
}
