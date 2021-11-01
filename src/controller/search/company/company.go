package company

import (
	"github.com/k0kubun/pp"
	"wiki_global/src/common/convert"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/global/response"
	"wiki_global/src/models/es"
	"wiki_global/src/models/es/escompany"
	"wiki_global/src/models/mysql/attention"
	"wiki_global/src/services/usercenter"

	"github.com/gin-gonic/gin"
)

// GetList 公司搜索接口
func GetList(c *gin.Context) {

	var searchListRequest es.SearchListRequest
	if err := c.ShouldBindJSON(&searchListRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}
	pp.Println("公司列表参数")
	pp.Println(searchListRequest)
	// 获取组装好的条件
	filter, boolQuery, scoreQuery := searchListRequest.FilterConditionV2(es.CompanyFields, searchListRequest.Language)

	// 筛选
	res, count, err := escompany.ListQuery(c, searchListRequest.Language, filter, boolQuery, scoreQuery)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	response.PageOK(c, res, count, searchListRequest.PageIndex, searchListRequest.PageSize, "ok")
}

// GetInfo 获取公司详情接口
func GetInfo(c *gin.Context) {

	uuid := c.Request.FormValue("uuid")
	userId := c.Request.FormValue("user_id")
	if uuid == "" {
		i18nresponse.Error(c, "1010004")
		return
	}

	res, err := escompany.DetailQuery(c, uuid)
	if err != nil && err.Error() == "不存在此公司" {
		i18nresponse.Error(c, "404")
		return
	} else if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	//查询此公司是否关注
	if userId != "" {
		_, count, err := attention.Attention{UserId: userId, Uuid: uuid, Type: 1}.QueryInfo(c)
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

// 获取筛选条件聚合
func GetCondition(c *gin.Context) {

	var searchListRequest es.SearchListRequest
	if err := c.ShouldBindJSON(&searchListRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	// 获取组装好的条件
	filter, boolQuery, _ := searchListRequest.FilterCondition(es.CompanyFields)

	res, err := escompany.GetCondition(c, filter, boolQuery)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
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

	res, err := escompany.CompletionSuggestQuery(c, prefix)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", res)
}

// ManualUpdate 手动更新
func ManualUpdate(c *gin.Context) {

	uuid := c.Param("uuid")
	var manualUpdateRequest escompany.ManualUpdateRequest
	if err := c.ShouldBindJSON(&manualUpdateRequest); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	res, err := escompany.ManualUpdate(c, uuid, manualUpdateRequest)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", res)
}

// GetEvaluationInfo wikibit详情
func GetEvaluationInfo(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")
	currency := ""
	if languageCode == "zh-cn" {
		currency = "cny"
	} else {
		currency = "usd"
	}

	info, err := usercenter.GetEvaluationInfo(languageCode, evaluationCode, currency)
	if err != nil || info.Succeed != true {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info.Result)
}

// GetTraderInfo wikifx详情
func GetTraderInfo(c *gin.Context) {

	traderCode := c.Request.FormValue("trader_code")
	languageCode := c.Request.FormValue("language_code")

	info, err := usercenter.GetTraderInfo(languageCode, traderCode)
	if err != nil || info.Succeed != true {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info.Result)
}

// GetSpreadCode wikifx点差
func GetSpreadCode(c *gin.Context) {

	traderCode := c.Request.FormValue("trader_code")
	languageCode := c.Request.FormValue("language_code")

	info, err := usercenter.GetSpreadCode(traderCode, languageCode)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// GetSpreadHighLowByCode wikifx点差-详情
func GetSpreadHighLowByCode(c *gin.Context) {

	traderCode := c.Request.FormValue("trader_code")
	name := c.Request.FormValue("name")
	type1 := c.Request.FormValue("type")

	info, err := usercenter.GetSpreadHighLowByCode(traderCode, name, convert.StrToInt64(type1))
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// GetGroupRelationShip wikibit详情-集团关系
func GetGroupRelationShip(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")

	info, err := usercenter.GetGroupRelationShip(languageCode, evaluationCode)
	if err != nil || info.Succeed != true {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info.Result)
}

// TradingEnvironment wikifx详情-交易环境
func TradingEnvironment(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")

	info, err := usercenter.TradingEnvironment(evaluationCode, languageCode)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// HistoryScores wikifx详情-交易环境-历史评分
func HistoryScores(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")

	info, err := usercenter.HistoryScores(evaluationCode, languageCode)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}
