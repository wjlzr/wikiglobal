package mongo

import (
	"context"
	"sync"
	"time"
	"wiki_global/src/config"
	"wiki_global/src/utils/log"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/zap"
)

var (
	err    error
	once   sync.Once
	client *mongo.Client
)

//集群连接
func Connect() {

	once.Do(func() {
		url := config.Conf().Mongo.Host
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// 进行连接并获取到 client
		if client, err = mongo.Connect(ctx, options.Client().ApplyURI(url)); err != nil {
			panic("Mongodb connect err：" + err.Error())
		}

		// ping ip 检查
		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			log.Logger().Error("Mongodb ping ip err: "+url, zap.Error(err))
			panic("Mongodb ping ip err" + err.Error())
		}
	})
}

// Mongo服务
func Client(dbName string) *mongo.Database {
	return client.Database(dbName)
}
