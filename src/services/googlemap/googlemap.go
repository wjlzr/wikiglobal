package googlemap

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"wiki_global/src/config"
	"wiki_global/src/services"
	"wiki_global/src/utils/log"

	"github.com/buger/jsonparser"
	"go.uber.org/zap"
)

//const getWay = "https://maps.googleapis.com/maps/api/geocode/json"
const getWay = "https://www.wikibtc.com/api/v1/googleMap/findCoordinateByAddress"

type Response struct {
	Code int64      `json:"code"`
	Msg  string     `json:"msg"`
	Data Coordinate `json:"data"`
}

type Coordinate struct {
	Lat float64 `json:"lat"` // 维度
	Lng float64 `json:"lng"` // 经度
}

type Request struct {
	Address string `json:"address"`
}

// FindCoordinateByAddress discard(已废弃)
func FindCoordinateByAddressDiscard(address string) (coordinate Coordinate, err error) {

	// 地址为空则直接返回
	if address == "" {
		return coordinate, nil
	}

	response, err := services.Request(http.MethodGet, getWay+"?address="+url.QueryEscape(address)+"&key="+config.Conf().GoogleMap.Key, nil)
	if err != nil {
		log.Logger().Error("地址转经纬度 err: ", zap.Error(err))
	}
	content, err := services.ResponseHandle(response)
	if err != nil {
		return coordinate, nil
	}
	status, err := jsonparser.GetString(content, "status")
	if err != nil {
		log.Logger().Error("jsonparser 解析错误 err: ", zap.Error(err))
	}

	if status != "OK" {
		errorMessage, _ := jsonparser.GetString(content, "error_message")
		log.Logger().Error("googlemap 调用地理编码失败 err: ", zap.Error(errors.New(errorMessage)))
		return Coordinate{}, errors.New("googlemap 调用地理编码失败")
	}

	value, _, _, err := jsonparser.Get(content, "results", "[0]", "geometry", "location")
	if err != nil {
		log.Logger().Error("googlemap 地理编码转换失败 err: ", zap.Error(err))
		return Coordinate{}, errors.New("googlemap 地理编码转换失败")
	}
	_ = json.Unmarshal(value, &coordinate)
	return
}

// FindCoordinateByAddress
func FindCoordinateByAddress(address string) (coordinate Coordinate, err error) {

	// 地址为空则直接返回
	if address == "" {
		return coordinate, nil
	}

	jsonStr, _ := json.Marshal(Request{Address: address})
	response, err := services.Request(http.MethodPost, getWay, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Logger().Error("googlemap FindCoordinateByAddress 请求 err：", zap.Error(err))
		return coordinate, nil
	}

	content, err := services.ResponseHandle(response)
	if err != nil {
		return coordinate, nil
	}
	var res Response
	_ = json.Unmarshal(content, &res)
	if res.Code != 200 {
		return coordinate, nil
	}
	return res.Data, nil
}
