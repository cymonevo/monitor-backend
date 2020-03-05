package provider

import (
	"sync"

	"github.com/cymonevo/monitor-backend/internal/elastic"
)

var (
	esClient     elastic.Client
	syncEsClient sync.Once
)

func GetESClient() elastic.Client {
	syncEsClient.Do(func() {
		cfg := GetAppConfig().ESConfig
		esClient = elastic.New(cfg)
	})
	return esClient
}
