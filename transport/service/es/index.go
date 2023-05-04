package es

import (
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
)

type (
	IndexFormat func(m map[string]interface{}) string
	IndexFunc   func() string

	Index struct {
		client       *elastic.Client
		indexFormat  IndexFormat
		indices      map[string]lang.PlaceholderType
		lock         sync.RWMutex
		singleFlight syncx.SingleFlight
	}
)

//func NewIndex(client *elastic.Client, indexFormat string, loc *time.Location) *Index {
//	return &Index{
//		client:       client,
//		indexFormat:  buildIndexFormatter(indexFormat, loc),
//		indices:      make(map[string]lang.PlaceholderType),
//		singleFlight: syncx.NewSingleFlight(),
//	}
//}
