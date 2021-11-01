package main

import (
	"fmt"
	"wiki_global/src/config"
	"wiki_global/src/db/es"
	"wiki_global/src/db/mysql"
	"wiki_global/src/router"
	"wiki_global/src/utils/log"
	"wiki_global/src/utils/token"

	"go.uber.org/zap"
)

//主函数入口
func main() {
	//初始化配置文件
	config.LoadConfig()

	//初始化log
	log.Init("logs")

	//mysql初始化
	err := mysql.Init(
		config.Conf().MySQL.DriverName,
		config.Conf().MySQL.Dsn,
		config.Conf().MySQL.MaxOpenConns,
		config.Conf().MySQL.MaxIdleConns,
	)
	if err != nil {
		log.Logger().Error(" mysql connect error", zap.Error(err))
		return
	}

	//redis集群连接
	//err = redis.NewRedisClusterClient(config.Conf().RedisClusterConfig())
	//if err != nil {
	//	log.Logger().Error("redis cluster connect error", zap.Error(err))
	//	return
	//}

	//es集群连接
	es.Connect()

	//mongo连接
	//mongo.Connect()

	//初始化认证token
	token := new(token.Token)
	token.InitToken(config.Conf().API.AuthToken)
	//gin路由引擎配置
	engine := router.RouterEngine(token, log.Logger())
	//启动服务
	engine.Run(fmt.Sprintf("%s:%d",
		config.Conf().Application.Host,
		config.Conf().Application.Port))
}
