package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var (
	clusterClient *redis.ClusterClient
)

//初始化集群
func NewRedisClusterClient(cs []interface{}, is []int) (err error) {
	if len(cs) != 2 && len(is) != 2 {
		return
	}

	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       cs[0].([]string),
		Password:    cs[1].(string),
		DialTimeout: time.Second * time.Duration(is[0]),
		PoolSize:    is[1],
	})

	_, err = clusterClient.Ping().Result()
	return
}

//
func RedisClusterClient() *redis.ClusterClient {
	return clusterClient
}
