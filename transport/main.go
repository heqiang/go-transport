package main

import (
	"flag"
	"fmt"
	"news_data_transport/transport/config"
	"news_data_transport/transport/service/kf"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "Specify the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	kf.KfConsumer(c)

	g := gin.Default()

	g.Run(fmt.Sprintf(":%s", c.Port))

}
