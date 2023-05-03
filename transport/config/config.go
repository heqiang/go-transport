package config

import (
	"github.com/zeromicro/go-queue/kq"
)

type Config struct {
	ServiceName string    `json:"serviceName"`
	Port        string    `json:"Port"`
	Kq          kq.KqConf `json:"Kafka"`
	Es          EsConf    `json:"Elasticsearch"`
}

type EsConf struct {
	Hosts         []string
	Index         string
	DocType       string `json:",default=doc"`
	TimeZone      string `json:",optional"`
	MaxChunkBytes int    `json:",default=15728640"` // default 15M
	Compress      bool   `json:",default=false"`
	Username      string `json:",optional"`
	Password      string `json:",optional"`
}
