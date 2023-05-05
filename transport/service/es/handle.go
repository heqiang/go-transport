package es

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/executors"
	"github.com/zeromicro/go-zero/core/logx"
	"news_data_transport/transport/config"
)

var NetWorkError = errors.New("网络错误")

type (
	InsertDoc struct {
		client   *elastic.Client
		index    string
		inserter *executors.ChunkExecutor
	}
	valueWithIndex struct {
		index string
		val   string
	}
)

func NewHandle(client *elastic.Client, index string, c config.EsConf) *InsertDoc {
	writer := &InsertDoc{
		client: client,
		index:  index,
	}
	writer.inserter = executors.NewChunkExecutor(writer.execute, executors.WithChunkBytes(c.MaxChunkBytes))
	return writer
}

func (m *InsertDoc) IndexExists(index string) (bool, error) {
	return m.client.IndexExists(index).Do(context.Background())
}

func (m *InsertDoc) CreateIndex(index string) (err error) {
	exists, err := m.IndexExists(index)
	if err != nil {
		return NetWorkError
	}
	if exists {
		return err
	} else {
		_, err = m.client.CreateIndex(index).Body(getMapping()).Do(context.Background())
		if err != nil {
			return err
		}
		return nil
	}

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
