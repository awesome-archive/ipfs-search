package config

import (
	"time"
)

// Default() returns default configuration
func Default() *Config {
	return &Config{
		Tika{
			IpfsTikaURL:     "http://localhost:8081",
			IpfsTikaTimeout: 300 * time.Duration(time.Second),
			MetadataMaxSize: 50 * 1024 * 1024,
		},
		IPFS{
			IpfsAPI:     "localhost:5001",
			IpfsTimeout: 360 * time.Duration(time.Second),
		},
		ElasticSearch{
			ElasticSearchURL: "http://localhost:9200",
		},
		AMQP{
			AMQPURL: "amqp://guest:guest@localhost:5672/",
		},
		Crawler{
			HashWait:    time.Duration(100 * time.Millisecond),
			FileWait:    time.Duration(100 * time.Millisecond),
			HashWorkers: 140,
			FileWorkers: 120,
			RetryWait:   2 * time.Duration(time.Second),
			PartialSize: 262144,
		},
	}
}
