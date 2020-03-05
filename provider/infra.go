package provider

import (
	"github.com/cymonevo/monitor-backend/internal/redis"
	"sync"

	"github.com/cymonevo/monitor-backend/internal/database"
	"github.com/cymonevo/monitor-backend/internal/mq"
)

var (
	dbClient     database.Client
	syncDbClient sync.Once

	redisClient     redis.Client
	syncRedisClient sync.Once

	mqServer     mq.Server
	syncMqClient sync.Once
)

func GetDBClient() database.Client {
	syncDbClient.Do(func() {
		cfg := GetAppConfig().DBConfig
		dbClient = database.New(cfg)
	})
	return dbClient
}

func GetRedisClient() redis.Client {
	syncRedisClient.Do(func() {
		cfg := GetAppConfig().RedisConfig
		redisClient = redis.New(&cfg)
	})
	return redisClient
}

func GetMQClient() mq.Server {
	syncMqClient.Do(func() {
		mqServer = mq.New()
	})
	return mqServer
}
