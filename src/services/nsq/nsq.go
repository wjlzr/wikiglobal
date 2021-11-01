package nsq

import (
	"wiki_global/src/config"
)

type Nsq struct {
	Topic string
	Host  string
	Port  string
}

func NewNsq() Nsq {
	return Nsq{
		Host: config.Conf().Nsq.Host,
		Port: config.Conf().Nsq.Port,
	}
}
