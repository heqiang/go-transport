package kf

import (
	"fmt"
	"news_data_transport/transport/config"

	"github.com/zeromicro/go-queue/kq"
)

func KfConsumer(c config.Config) {
	q := kq.MustNewQueue(c.Kq, kq.WithHandle(func(k, v string) error {
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
