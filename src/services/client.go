package services

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"wiki_global/src/utils/log"

	"github.com/k0kubun/pp"

	"go.uber.org/zap"
)

//统一请求分发
func Request(method, url string, body io.Reader) (request *http.Request, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Logger().Error("services http request err：", zap.Error(err))
		return request, err
	}
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

//返回参数统一处理
func ResponseHandle(request *http.Request) (content []byte, err error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		pp.Println(err.Error())
		return content, err

	}
	defer resp.Body.Close()
	content, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("OpenApi返回值：%s \n", string(content))
	return content, nil
}
