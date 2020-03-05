package model

import (
	"github.com/cymonevo/monitor-backend/internal/database"
	"github.com/cymonevo/monitor-backend/internal/mq"
	"github.com/cymonevo/monitor-backend/internal/redis"
)

type Factory interface {
	NewGetByIDModel() *GetArticleModel
	NewHealthModel() *HealthModel
}

type articleFactory struct {
	db        database.Client
	redis     redis.Client
	publisher mq.Publisher
}

func NewArticleFactory(db database.Client, redis redis.Client, publisher mq.Publisher) Factory {
	return &articleFactory{
		db:        db,
		redis:     redis,
		publisher: publisher,
	}
}

func (f *articleFactory) NewGetByIDModel() *GetArticleModel {
	return &GetArticleModel{}
}

func (f *articleFactory) NewHealthModel() *HealthModel {
	return &HealthModel{
		dbClient:    f.db,
		redisClient: f.redis,
		publisher:   f.publisher,
	}
}
