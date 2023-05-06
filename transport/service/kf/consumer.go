package kf

import (
	"fmt"

	"github.com/zeromicro/go-queue/kq"
)

func Consumer() {
	var c kq.KqConf
	q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
