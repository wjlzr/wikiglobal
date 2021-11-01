package config

import (
	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
	"path/filepath"
	"sync"
)

//配置信息
type Config struct {
	Application      application
	MySQL            mysql
	RedisCluster     rediscluster
	API              api
	Log              log
	Es               esCluster
	UserCenter       userCenter
	Mongo            mongo
	Nsq              nsq
	LocalEs          localEs
	GoogleMap        googleMap
	Encryption       encryption
	WeChat           wechat
	Qq               qq
	Apple            apple
	ElasticSearchApi elasticsearchapi
	Wikifx           wikifx
	Wikibit          wikibit
}

//服务配置
type application struct {
	Mode string `toml:"mode"` //模式
	Host string `toml:"host"` //服务器名
	Name string `toml:"name"` //服务名称
	Port int    `toml:"port"` //端口
}

//mysql配置
type mysql struct {
	DriverName   string `toml:"driver_name"`
	Dsn          string `toml:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
}

//redis配置
type rediscluster struct {
	Addrs       []string `toml:"addrs"`
	Password    string   `toml:"password"`
	DialTimeout int      `toml:"dial_timeout"`
	PoolSize    int      `toml:"pool_size"`
}

//es配置
type esCluster struct {
	Hosts    []string `toml:"hosts"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
}

// locales配置
type localEs struct {
	Hosts    []string `toml:"hosts"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
}

type mongo struct {
	Host string `toml:"host"`
}

//用户中台配置
type userCenter struct {
	TestUrl string `toml:"test_url"`
	ProdUrl string `toml:"prod_url"`
	User    string `toml:"user"`
}

//nsq配置
type nsq struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

//nsq配置
type googleMap struct {
	Key string `toml:"key"`
}

type encryption struct {
	AesSecretKey  string `toml:"aes_secret_key"`
	SecretKey     int    `toml:"secret_key"`
	RideSecretKey int    `toml:"ride_secret_key"`
}

//log
type log struct {
	Path string `toml:"path"`
}

//api
type api struct {
	AllowPathPrefixSkipper []string `toml:"allow_path_prefix_skipper"`
	AuthToken              string   `toml:"auth_token"`
}

// wechat
type wechat struct {
	Gateway   string `toml:"gateway"`
	AppId     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

// qq
type qq struct {
	Gateway string `toml:"gateway"`
	AppId   string `toml:"app_id"`
	AppKey  string `toml:"app_key"`
}

type apple struct {
	KeyId  string `toml:"key_id"`
	TeamId string `toml:"team_id"`
	AppId  string `toml:"app_id"`
}

type elasticsearchapi struct {
	Gateway string `toml:"gateway"`
}

type wikifx struct {
	Gateway  string `toml:"gateway"`
	Gateway1 string `toml:"gateway1"`
	Gateway2 string `toml:"gateway2"`
	Gateway3 string `toml:"gateway3"`
	Gateway4 string `toml:"gateway4"`
	Gateway5 string `toml:"gateway5"`
}

type wikibit struct {
	Gateway string `toml:"gateway"`
}

var (
	cfg  *Config
	once sync.Once
)

//加载配置文件
func LoadConfig() {
	once.Do(func() {
		fp, err := filepath.Abs("./src/config/config.toml")
		//fp, err := filepath.Abs("/data/import/config.toml")
		pp.Println(fp)
		if err == nil {
			_, _ = toml.DecodeFile(fp, &cfg)
		}
	})
}

//获取配置对象
func Conf() *Config {
	return cfg
}

//获取redis集群信息
func (c *Config) RedisClusterConfig() (cs []interface{}, is []int) {
	cs = []interface{}{
		c.RedisCluster.Addrs,
		c.RedisCluster.Password,
	}
	is = []int{
		c.RedisCluster.DialTimeout,
		c.RedisCluster.PoolSize,
	}
	return
}
