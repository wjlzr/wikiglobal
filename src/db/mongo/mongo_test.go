package mongo

import (
	"testing"
	"wiki_global/src/config"
	"wiki_global/src/utils/log"

	"github.com/k0kubun/pp"
)

func TestConnect(t *testing.T) {
	//初始化配置文件
	config.LoadConfig()

	//初始化log
	log.Init("logs")
	//mongo连接
	Connect()
	aa := Client("ml_company")
	pp.Println(aa)
}
