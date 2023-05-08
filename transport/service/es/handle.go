package es

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/syncx"
	"news_data_transport/transport/config"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/logx"
)

var NetWorkError = errors.New("网络错误")

type (
	InsertDoc struct {
		client       *elastic.Client
		index        string
		inserter     *executors.ChunkExecutor
		indices      map[string]lang.PlaceholderType
		lock         sync.RWMutex
		singleFlight syncx.SingleFlight
	}
	valueWithIndex struct {
		index string
		val   string
	}
)

func NewHandle(client *elastic.Client, index string, c config.EsConf) *InsertDoc {
	writer := &InsertDoc{
		client:       client,
		index:        index,
		indices:      make(map[string]lang.PlaceholderType),
		singleFlight: syncx.NewSingleFlight(),
	}
	writer.inserter = executors.NewChunkExecutor(writer.execute, executors.WithChunkBytes(c.MaxChunkBytes))
	return writer
}

func (m *InsertDoc) CreateIndex(index string) (err error) {
	_, err = m.singleFlight.Do(index, func() (any, error) {
		m.lock.RLock()
		defer m.lock.RUnlock()

		if _, ok := m.indices[index]; ok {
			return nil, nil
		}
		existsService := elastic.NewIndicesExistsService(m.client)
		existsService.Index([]string{index})
		exist, err := existsService.Do(context.Background())
		if err != nil {
			return nil, NetWorkError
		}
		if exist {
			return nil, err
		}

		createService := m.client.CreateIndex(index).Body(getMapping())
		if err := fx.DoWithRetry(func() error {
			_, err := createService.Do(context.Background())
			return err
		}); err != nil {
			return nil, err
		}

		return nil, nil

	})

	return err

}
func (m *InsertDoc) write(index, val string) error {
	return m.inserter.Add(valueWithIndex{
		index: index,
		val:   val,
	}, len(val))
}

func (m *InsertDoc) Consume(_, val string) error {
	err := m.CreateIndex(m.index)
	if err != nil {
		return err
	}

	var d map[string]interface{}
	if err := jsoniter.Unmarshal([]byte(val), &d); err != nil {
		return err
	}

	bs, err := jsoniter.Marshal(d)
	if err != nil {
		return err
	}
	return m.write(m.index, string(bs))
}

func (m *InsertDoc) execute(vals []interface{}) {
	var bulk = m.client.Bulk()
	for _, val := range vals {
		pair := val.(valueWithIndex)
		req := elastic.NewBulkIndexRequest().Index(pair.index)
		req = req.Doc(pair.val)
		bulk.Add(req)
	}
	resp, err := bulk.Do(context.Background())
	if err != nil {
		logx.Error(err)
		return
	}

	// bulk error in docs will report in response items
	if !resp.Errors {
		return
	}

	for _, imap := range resp.Items {
		for _, item := range imap {
			if item.Error == nil {
				continue
			}

			logx.Error(item.Error)
		}
	}
}
