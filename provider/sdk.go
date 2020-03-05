package provider

import (
	"sync"

	"github.com/cymonevo/monitor-backend/sdk"
	"github.com/cymonevo/monitor-backend/sdk/kloudless"
)

var (
	kloudlessClient     kloudless.Client
	syncKloudlessClient sync.Once
)

func GetKloudlessClient() kloudless.Client {
	syncKloudlessClient.Do(func() {
		//TODO: Get from config file
		cfg := kloudless.Config{
			Config: sdk.Config{Timeout: 20, URL: "https://api.kloudless.com/v1"},
			APIKey: "insert key here",
			AccID:  0,
		}
		kloudlessClient = kloudless.New(cfg)
	})
	return kloudlessClient
}
