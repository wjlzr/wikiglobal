package usercenter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
	"net/http"
	"wiki_global/src/config"
	"wiki_global/src/utils/log"
)

const (
	getEvaluationInfo    = "Wikibit/evaluation/getEvaluationInfo"
	getGroupRelationShip = "Wikibit/evaluation/getgrouprelationship"
)

// GetEvaluationInfo 受评方详情
func GetEvaluationInfo(languageCode, evaluationCode, currency string) (result *EvaluationInfoDataResponse, err error) {

	request, err := request(http.MethodGet, config.Conf().Wikibit.Gateway+fmt.Sprintf(getEvaluationInfo+"?languageCode=%s&evaluationCode=%s&currency=%s&countryCode=%s", languageCode, evaluationCode, currency, "156"), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetEvaluationInfo 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(request)
	var v EvaluationInfoResponse
	_ = json.Unmarshal(content, &v)
	if v.Code != 200 || v.Success != true {
		log.Logger().Info("UserCenter 获取受评方失败 response：", zap.Any("response", v))
		return nil, errors.New(v.Msg)
	}
	pp.Println("区块详情")
	pp.Println(v.Data)
	return &v.Data, nil
}

// GetGroupRelationShip 获取集团关系
func GetGroupRelationShip(languageCode, evaluationCode string) (result *EvaluationInfoDataResponse, err error) {

	request, err := request(http.MethodGet, config.Conf().Wikibit.Gateway+fmt.Sprintf(getGroupRelationShip+"?languageCode=%s&evaluationCode=%s&countryCode=%s", languageCode, evaluationCode, "156"), nil)
	if err != nil {
		log.Logger().Error("UserCenter GetGroupRelationShip 请求 err：", zap.Error(err))
		return nil, err
	}

	content := responseHandle(request)
	var v EvaluationInfoResponse
	_ = json.Unmarshal(content, &v)
	if v.Code != 200 || v.Success != true {
		log.Logger().Info("UserCenter 获取集团关系失败 response：", zap.Any("response", v))
		return nil, errors.New(v.Msg)
	}
	pp.Println("集团关系")
	pp.Println(v.Data)
	return &v.Data, nil
}
