package usercenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
	"wiki_global/src/config"
	"wiki_global/src/services/usercenter/wikifx"
	"wiki_global/src/utils/log"
)

const (
	traderDetail = "trader/detail"
	spreadCode   = "skyeye/spread/code"
	//spreadCode = "mt4api/skyeye/spread/code"
	spreadHighLow = "SkyEye/GetSpreadHighLowByCode"
	rankEnv       = "trader/rank/env"
	historyScores = "trader/rank/score"
)

// GetTraderInfo 交易商详情
func GetTraderInfo(languageCode, traderCode string) (result *EvaluationInfoDataResponse, err error) {

	req, err := request(http.MethodGet, config.Conf().Wikifx.Gateway+fmt.Sprintf(traderDetail+"?languageCode=%s&traderCode=%s&countryCode=%s", languageCode, traderCode, "156"), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetTraderInfo 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	var v EvaluationInfoDataResponse
	_ = json.Unmarshal(content, &v)
	if v.Succeed != true {
		log.Logger().Info("UserCenter 获取交易商详情失败 response：", zap.Any("response", v))
		return nil, errors.New(v.Message)
	}
	pp.Println("交易商详情")
	pp.Println(v)
	return &v, nil
}

// GetSpreadCode 交易商的点差
func GetSpreadCode(code, languageCode string) (result []interface{}, err error) {

	req, err := request(http.MethodGet, config.Conf().Wikifx.Gateway1+fmt.Sprintf(spreadCode+"?code=%s&languageCode=%s", code, languageCode), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetSpreadCode 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	_ = json.Unmarshal(content, &result)
	return result, nil
}

// GetSpreadHighLowByCode 交易商的点差详情
func GetSpreadHighLowByCode(code, name string, type1 int64) (result interface{}, err error) {

	req, err := request(http.MethodGet, config.Conf().Wikifx.Gateway2+fmt.Sprintf(spreadHighLow+"?code=%s&name=%s&type=%d", code, name, type1), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetSpreadHighLowByCode 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	_ = json.Unmarshal(content, &result)
	return result, nil
}

// TradingEnvironment 交易商的交易环境
func TradingEnvironment(code, languageCode string) (result interface{}, err error) {

	req, err := request(http.MethodGet, config.Conf().Wikifx.Gateway+fmt.Sprintf(rankEnv+"?traderCode=%s&languageCode=%s&countryCode=%s", code, languageCode, "156"), nil)
	if err != nil {
		log.Logger().Error("UserCenter TradingEnvironment 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	_ = json.Unmarshal(content, &result)
	return result, nil
}

// HistoryScores 交易商-交易环境-历史评分
func HistoryScores(code, languageCode string) (result interface{}, err error) {

	req, err := request(http.MethodGet, config.Conf().Wikifx.Gateway+fmt.Sprintf(historyScores+"?traderCode=%s&languageCode=%s&countryCode=%s", code, languageCode, "156"), nil)
	if err != nil {
		log.Logger().Error("UserCenter HistoryScores 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	_ = json.Unmarshal(content, &result)
	return result, nil
}

// Request wikifx接口统一请求分发
func Request(method string, query map[string]interface{}) (result interface{}, err error) {

	methodRequest, site, params := pretreatment(method, query)

	req, err := request(methodRequest, site, params)
	if err != nil {
		log.Logger().Error("UserCenter "+method+" 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(req)
	_ = json.Unmarshal(content, &result)
	return result, nil
}

// pretreatment 请求预处理
func pretreatment(method string, query map[string]interface{}) (string, string, io.Reader) {

	url := wikifx.Url()

	if url[method][0] == http.MethodGet {
		return url[method][0], url[method][1] + url[method][2] + HandleUrl(query), nil
	} else {
		jsonStr, _ := json.Marshal(query)
		return url[method][0], url[method][1] + url[method][2], bytes.NewBuffer(jsonStr)
	}
}

// HandleUrl url参数处理
func HandleUrl(query map[string]interface{}) string {
	var buf bytes.Buffer
	for k, v := range query {
		switch v.(type) {
		case int:
			buf.WriteString("&" + k + "=" + strconv.Itoa(v.(int)))
		case string:
			buf.WriteString("&" + k + "=" + v.(string))
		case int64:
			buf.WriteString("&" + k + "=" + strconv.FormatInt(v.(int64), 10))
		}
	}
	return "?" + buf.String()[1:]
}
