package es

import (
	"context"
	"sync"
	"wiki_global/src/config"
	"wiki_global/src/utils/log"

	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

var (
	err     error
	once    sync.Once
	client  *elastic.Client
	client1 *elastic.Client
)

//集群连接
func Connect() {
	once.Do(func() {
		urls := config.Conf().Es.Hosts
		client, err = elastic.NewClient(elastic.SetURL(urls...), elastic.SetBasicAuth(config.Conf().Es.Username, config.Conf().Es.Password), elastic.SetSniff(false))
		if err != nil {
			panic("Elasticsearch connect err：" + err.Error())
		}

		// ping ip 检查
		for _, url := range urls {
			if _, _, err := client.Ping(url).Do(context.Background()); err != nil {
				log.Logger().Error("Elasticsearch ping ip err: "+url, zap.Error(err))
				panic("Elasticsearch ping ip err" + err.Error())
			}
		}
	})
}

// ES服务
func Client() *elastic.Client {
	return client
}

//集群连接
func Connect1() {
	once.Do(func() {
		urls := config.Conf().LocalEs.Hosts
		client1, err = elastic.NewClient(elastic.SetURL(urls...), elastic.SetBasicAuth(config.Conf().LocalEs.Username, config.Conf().LocalEs.Password), elastic.SetSniff(false))
		if err != nil {
			panic("Elasticsearch connect err：" + err.Error())
		}

		// ping ip 检查
		for _, url := range urls {
			if _, _, err := client1.Ping(url).Do(context.Background()); err != nil {
				log.Logger().Error("Elasticsearch ping ip err: "+url, zap.Error(err))
				panic("Elasticsearch ping ip err" + err.Error())
			}
		}
	})
}

// ES服务
func Client1() *elastic.Client {
	return client1
}
