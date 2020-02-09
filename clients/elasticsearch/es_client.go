package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10 * time.Second),
		elastic.SetSniff(false),
		
	)

	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) Index(interface{})(*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}