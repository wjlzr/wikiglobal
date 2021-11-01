package officers

import (
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/global/response"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/esofficers"
	"wiki_global/src/models/mysql/attention"

	"github.com/gin-gonic/gin"
)

// GetList 高管搜索接口
func GetList(c *gin.Context) {

	var searchListRequest es.SearchListRequest
	if err := c.ShouldBindJSON(&searchListRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取组装好的条件
	filter, boolQuery, _ := searchListRequest.FilterCondition(es.OfficersFields)

	// 筛选
	res, count, err := esofficers.ListQuery(c, filter, boolQuery)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	response.PageOK(c, res, count, searchListRequest.PageIndex, searchListRequest.PageSize, "ok")
}

// GetInfo 获取高管详情接口
func GetInfo(c *gin.Context) {

	uuid := c.Request.FormValue("uuid")
	userId := c.Request.FormValue("user_id")
	if uuid == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	res, err := esofficers.DetailQuery(c, uuid)
	if err != nil && err.Error() == "不存在此高管" {
		i18nresponse.Error(c, "500")
		return
	} else if err != nil {
		i18nresponse.Error(c, "404")
		return
	}

	//查询此公司是否关注
	if userId != "0" {
		_, count, err := attention.Attention{UserId: userId, Uuid: uuid, Type: 2}.QueryInfo(c)
		if err != nil {
			i18nresponse.Error(c, "500")
			return
		}
		if count > 0 {
			res.IsAttention = true
		}
	}

	i18nresponse.Success(c, "ok", res)
}

// CompletionSuggest 完成推荐
func CompletionSuggest(c *gin.Context) {

	prefix := c.Request.FormValue("prefix")
	if prefix == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	res, err := esofficers.CompletionSuggestQuery(c, prefix)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", res)
}

// ManualUpdate 手动更新
func ManualUpdate(c *gin.Context) {

	uuid := c.Param("uuid")
	var manualUpdateRequest esofficers.ManualUpdateRequest
	if err := c.ShouldBindJSON(&manualUpdateRequest); err != nil {
		response.Error(c, -1, "缺少参数")
		return
	}

	res, err := esofficers.ManualUpdate(c, uuid, manualUpdateRequest)
	if err != nil {
		response.Error(c, 500, "服务器错误")
		return
	}

	response.Success(c, res, "ok")
}
