package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"time"
)

type (
	EsConf struct {
		Hosts         []string
		Index         string
		DocType       string `json:",default=doc"`
		TimeZone      string `json:",optional"`
		MaxChunkBytes int    `json:",default=15728640"` // default 15M
		Compress      bool   `json:",default=false"`
		Username      string `json:",optional"`
		Password      string `json:",optional"`
	}
	Cluster struct {
		ElasticSearch EsConf
		Kafka         KafkaConf
	}
	KafkaConf struct {
		service.ServiceConf
		Brokers    []string
		Group      string
		Topics     []string
		Offset     string `json:",options=first|last,default=last"`
		Conns      int    `json:",default=1"`
		Consumers  int    `json:",default=8"`
		Processors int    `json:",default=8"`
		MinBytes   int    `json:",default=10240"`    // 10K
		MaxBytes   int    `json:",default=10485760"` // 10M
		Username   string `json:",optional"`
		Password   string `json:",optional"`
	}
	Config struct {
		ServiceName string        `json:"serviceName"`
		GracePeriod time.Duration `json:",default=10s"`
		Port        string        `json:"Port"`
		Kq          kq.KqConf     `json:"Kafka"`
		Clusters    []Cluster
	}
)
