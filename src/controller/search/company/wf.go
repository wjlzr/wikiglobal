package company

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"wiki_global/src/global/i18nresponse"
	"wiki_global/src/services/usercenter"
	"wiki_global/src/services/usercenter/wikifx"
	"wiki_global/src/utils"
)

// SurveyList wikifx详情-实勘列表
func SurveyList(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")
	pageIndex := c.Request.FormValue("page_index")
	pageSize := c.Request.FormValue("page_size")

	var query = map[string]interface{}{
		"code":    evaluationCode,
		"country": 156,
		"lan":     languageCode,
		"index":   pageIndex,
		"size":    pageSize,
	}

	info, err := usercenter.Request("surveyList", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// Survey wikifx详情-实勘详情
func Survey(c *gin.Context) {

	sid := c.Request.FormValue("sid")
	languageCode := c.Request.FormValue("language_code")
	version := c.Request.FormValue("version")

	var query = map[string]interface{}{
		"sid":     sid,
		"country": 156,
		"lan":     languageCode,
		"version": version,
		"pro":     4,
	}

	info, err := usercenter.Request("survey", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// Survey wikifx详情-旗舰店广告
func UltimateMember(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	languageCode := c.Request.FormValue("language_code")
	isApp := c.Request.FormValue("is_app")

	var query = map[string]interface{}{
		"traderCode":   evaluationCode,
		"countryCode":  010,
		"languageCode": languageCode,
		"isApp":        isApp,
		"ver":          "2.0.0",
	}

	info, err := usercenter.Request("ultimateMember", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// ResearchIsShow wikifx详情-是否显示红色按钮处功能
func ResearchIsShow(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	isApp := c.Request.FormValue("is_app")

	var query = map[string]interface{}{
		"code":       evaluationCode,
		"isdomestic": "false",
		"isApp":      isApp,
	}

	info, err := usercenter.Request("researchIsShow", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// BrokerHistoryDataExist wikifx详情-经纪商是否有历史数据（详情页是否显示蓝色按钮）
func BrokerHistoryDataExist(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"brokerId": evaluationCode,
	}

	info, err := usercenter.Request("brokerHistoryDataExist", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// Epc wikifx详情-epc
func Epc(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"traderCode":   evaluationCode,
		"languageCode": utils.GetLang(c),
		"userId":       "",
		//"countryCode":  utils.CountryCode(c),
		"countryCode": 702,
		"ip":          c.ClientIP(),
	}

	info, err := usercenter.Request("epc", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// Epc wikifx详情-提示语
func Tips(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"traderCode":   evaluationCode,
		"languageCode": utils.GetLang(c),
		"targetType":   1,
		"countryCode":  utils.CountryCode(c),
	}

	info, err := usercenter.Request("tips", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// TraderCount wikifx详情-提示语-统计
func TraderCount(c *gin.Context) {

	var traderCountReq wikifx.TraderCountReq
	if err := c.ShouldBindJSON(&traderCountReq); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	traderCountReq.SandBox = 0
	traderCountReq.UserId = ""
	traderCountReq.Platform = 3
	traderCountReq.Ip = c.ClientIP()
	traderCountReq.Country = utils.CountryCode(c)
	traderCountReq.Lang = utils.GetLang(c)

	query := structs.Map(traderCountReq)

	info, err := usercenter.Request("tradeCount", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// TraderAccount wikifx详情-账户类型
func TraderAccount(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"code":         evaluationCode,
		"languageCode": utils.GetLang(c),
		"countryCode":  utils.CountryCode(c),
	}

	info, err := usercenter.Request("tradeAccount", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// MT4 wikifx详情-MT4/5
func MT4(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"code":         evaluationCode,
		"languageCode": utils.GetLang(c),
		"countryCode":  utils.CountryCode(c),
		"remote_addr":  c.ClientIP(),
	}

	info, err := usercenter.Request("mt4", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// MT4 wikifx详情-MT4/5
func Fake(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")

	var query = map[string]interface{}{
		"code":         evaluationCode,
		"languageCode": utils.GetLang(c),
		"countryCode":  utils.CountryCode(c),
		"remote_addr":  c.ClientIP(),
	}

	info, err := usercenter.Request("fake", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// ResearchInfo wikifx详情-研究院数据（饼图数据）
func ResearchInfo(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	type1 := c.Request.FormValue("type")
	isApp := c.Request.FormValue("is_app")

	var query = map[string]interface{}{
		"code":         evaluationCode,
		"languageCode": utils.GetLang(c),
		"countryCode":  utils.CountryCode(c),
		"isapp":        isApp,
		"type":         type1,
	}

	info, err := usercenter.Request("researchInfo", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// BrokerHistoryData wikifx详情-经纪商历史走势
func BrokerHistoryData(c *gin.Context) {

	var brokerHistoryDataReq wikifx.BrokerHistoryDataReq
	if err := c.ShouldBindJSON(&brokerHistoryDataReq); err != nil {
		i18nresponse.Error(c, "1010004")
		return
	}

	query := structs.Map(brokerHistoryDataReq)

	pp.Println(query)

	info, err := usercenter.Request("brokerHistoryData", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}

// ExposureList wikifx详情-曝光
func ExposureList(c *gin.Context) {

	evaluationCode := c.Request.FormValue("evaluation_code")
	pageIndex := c.Request.FormValue("page_index")
	pageSize := c.Request.FormValue("page_size")

	var query = map[string]interface{}{
		"traderCode":   evaluationCode,
		"languageCode": utils.GetLang(c),
		"countryCode":  utils.CountryCode(c),
		"pageIndex":    pageIndex,
		"pageSize":     pageSize,
		"uid":          "",
		"platform":     1,
	}

	info, err := usercenter.Request("exposureList", query)
	if err != nil {
		i18nresponse.Error(c, "500")
		return
	}

	i18nresponse.Success(c, "ok", info)
}
