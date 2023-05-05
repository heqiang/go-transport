package main

import (
	"flag"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/service"
	"news_data_transport/transport/config"
	"news_data_transport/transport/service/es"
)

var configFile = flag.String("f", "etc/config.yaml", "Specify the config file")

func toKqConf(c config.KafkaConf) []kq.KqConf {
	var ret []kq.KqConf

	for _, topic := range c.Topics {
		ret = append(ret, kq.KqConf{
			ServiceConf: c.ServiceConf,
			Brokers:     c.Brokers,
			Group:       c.Group,
			Topic:       topic,
			Offset:      c.Offset,
			Conns:       c.Conns,
			Consumers:   c.Consumers,
			Processors:  c.Processors,
			MinBytes:    c.MinBytes,
			MaxBytes:    c.MaxBytes,
			Username:    c.Username,
			Password:    c.Password,
		})
	}

	return ret
}
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	proc.SetTimeToForceQuit(c.GracePeriod)

	group := service.NewServiceGroup()
	defer group.Stop()

	for _, processor := range c.Clusters {
		client, err := elastic.NewClient(
			elastic.SetSniff(false),
			elastic.SetURL(processor.Output.ElasticSearch.Hosts...),
			elastic.SetBasicAuth(processor.Output.ElasticSearch.Username, processor.Output.ElasticSearch.Password),
		)
		logx.Must(err)

		handle := es.NewHandle(client, processor.Output.ElasticSearch.Index, processor.Output.ElasticSearch)
		for _, k := range toKqConf(processor.Input.Kafka) {
			group.Add(kq.MustNewQueue(k, handle))
		}
	}

	group.Start()
}
