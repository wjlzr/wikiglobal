package data

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"wiki_global/src/config"
	"wiki_global/src/services"
	"wiki_global/src/utils/log"
)

const getMultiple = "wikicore/getMultiple"

type searchRequest struct {
	Language string   `json:"language"`
	Country  string   `json:"country"`
	Codes    []string `json:"codes"`
}

type SearchResponse struct {
	Result  []Result `json:"result"`
	Succeed bool     `json:"succeed"`
	Message string   `json:"message"`
}

type Result struct {
	Code            string `json:"code"`
	IsVr            bool   `json:"isVr"`
	Color           string `json:"color"`
	Annotation      string `json:"annotation"`
	Logo            string `json:"logo"`
	RegisterCountry string `json:"registerCountry"`
	LocalFullName   string `json:"localFullName"`
}

func GetSearchData(codes []string, language string) (s SearchResponse, err error) {

	jsonStr, _ := json.Marshal(searchRequest{Language: language, Country: "156", Codes: codes})
	request, err := services.Request(http.MethodPost, config.Conf().ElasticSearchApi.Gateway+getMultiple, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("data GetSearchData 请求 err：", zap.Error(err))
		return SearchResponse{}, err
	}

	content, _ := services.ResponseHandle(request)
	_ = json.Unmarshal(content, &s)
	if s.Succeed != true {
		log.Logger().Info("data GetSearchData Get Error response：", zap.Any("response", s))
		return SearchResponse{}, errors.New(s.Message)
	}
	return s, nil
}
