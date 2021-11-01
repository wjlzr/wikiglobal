package wikifx

import (
	"net/http"
	"wiki_global/src/config"
)

func Url() map[string][]string {

	var Url = map[string][]string{
		// wikifx-交易所的实勘列表
		"surveyList": {http.MethodGet, config.Conf().Wikifx.Gateway3, "WikiSurvey/survey/trader"},
		// wikifx-交易所的实勘详情
		"survey": {http.MethodGet, config.Conf().Wikifx.Gateway3, "WikiSurvey/survey"},
		// wikifx-交易所的旗舰店广告
		"ultimateMember": {http.MethodGet, config.Conf().Wikifx.Gateway, "ultimatemember/ad/list"},
		// wikifx-交易所的是否显示红色按钮处功能
		"researchIsShow": {http.MethodGet, config.Conf().Wikifx.Gateway, "research/isshow"},
		// wikifx-经纪商是否有历史数据（详情页是否显示蓝色按钮）
		"brokerHistoryDataExist": {http.MethodPost, config.Conf().Wikifx.Gateway4, "brustInfo/brokerHistoryDataExist"},
		// wikifx-epc
		"epc": {http.MethodGet, config.Conf().Wikifx.Gateway, "trader/epc"},
		// wikifx-epc
		"tips": {http.MethodGet, config.Conf().Wikifx.Gateway, "trader/tips/new"},
		// wikifx-提示语-统计
		"tradeCount": {http.MethodPost, config.Conf().Wikifx.Gateway3, "Statistics/trader/count"},
		// wikifx-账户类型
		"tradeAccount": {http.MethodGet, config.Conf().Wikifx.Gateway3, "Wikifx2/trader/account"},
		// wikifx-MT4/5
		"mt4": {http.MethodGet, config.Conf().Wikifx.Gateway3, "Wikifx2/trader/mt4/detail"},
		// wikifx-查找冒充交易商
		"fake": {http.MethodGet, config.Conf().Wikifx.Gateway3, "Wikifx2/trader/fake"},
		// wikifx-研究院数据（饼图数据）
		"researchInfo": {http.MethodGet, config.Conf().Wikifx.Gateway3, "Wikifx2/research/get"},
		// wikifx-经纪商历史走势
		"brokerHistoryData": {http.MethodPost, config.Conf().Wikifx.Gateway4, "brustInfo/brokerHistoryData"},
		// wikifx-曝光
		"exposureList": {http.MethodGet, config.Conf().Wikifx.Gateway3, "WikiForum/topic/trader/listForWeb"},
	}
	return Url
}
