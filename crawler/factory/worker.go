package factory

import (
	"context"
	"github.com/ipfs-search/ipfs-search/crawler"
	"github.com/streadway/amqp"
)

// CrawlFunc returns a function crawling a particular indexable with
// given context
type CrawlFunc func(i *crawler.Indexable) func(context.Context) error

// Worker performs crawling based on a single AMQP message
type Worker struct {
	*crawler.Crawler

	Delivery  *amqp.Delivery
	CrawlFunc CrawlFunc
}

// Work takes a message with JSON body, converts it to a crawlable and
// calls CrawlFunc on it.
func (c *Worker) Work(ctx context.Context) error {
	// Create an Indexable from the message's body
	i, err := c.IndexableFromJSON(c.Delivery.Body)
	if err != nil {
		return err
	}

	// Call crawler function with context
	return c.CrawlFunc(i)(ctx)
}
