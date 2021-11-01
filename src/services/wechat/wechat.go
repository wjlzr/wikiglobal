package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"wiki_global/src/config"
	"wiki_global/src/services"
	"wiki_global/src/utils/log"
)

func GetAccessToken(code string) (*accessTokenResponse, error) {

	request, err := services.Request(http.MethodGet, config.Conf().WeChat.Gateway+fmt.Sprintf(accessToken+"?appid=%s&secret=%s&code=%s&grant_type=%s", config.Conf().WeChat.AppId, config.Conf().WeChat.AppSecret, code, grantType), nil)
	if err != nil {
		log.Logger().Error("wechat GetAccessToken 请求 err：", zap.Error(err))
		return nil, err
	}

	content, _ := services.ResponseHandle(request)
	var v accessTokenResponse
	_ = json.Unmarshal(content, &v)
	if v.ErrCode != 0 || v.ErrMsg != "" {
		log.Logger().Info("UserCenter GetAccessToken 获取微信accessToken response：", zap.Any("response", v))
		return nil, errors.New("500")
	}
	return &v, nil
}
