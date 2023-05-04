package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/core/service"
	"news_data_transport/transport/config"
)

var configFile = flag.String("f", "etc/config.yaml", "Specify the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	proc.SetTimeToForceQuit(c.GracePeriod)

	group := service.NewServiceGroup()
	defer group.Stop()

	//for _,processor:=range c.Clusters{
	//	client,err:=elastic.NewClient(
	//		elastic.SetSniff(false),
	//		elastic.SetURL(processor.ElasticSearch.Hosts...),
	//		elastic.SetBasicAuth(c.Kq.Username,c.Kq.Password),
	//		)
	//	logx.Must(err)
	//
	//	writer, err := es.NewWriter(processor.ElasticSearch)
	//	logx.Must(err)
	//
	//	var loc *time.Location
	//	tz := processor.ElasticSearch.TimeZone
	//	if len(tz) > 0 {
	//		loc, err = time.LoadLocation(tz)
	//		logx.Must(err)
	//	} else {
	//		loc = time.Local
	//	}
	//
	//
	//	indexer := es.
	//
	//
	//}
}
