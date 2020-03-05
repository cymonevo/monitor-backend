package provider

import (
	"sync"

	"github.com/cymonevo/monitor-backend/module/article/model"
)

var (
	articleFactory     model.Factory
	syncArticleFactory sync.Once
)

func GetArticleFactory() model.Factory {
	syncArticleFactory.Do(func() {
		articleFactory = model.NewArticleFactory(GetDBClient(), GetRedisClient(), GetPublisher())
	})
	return articleFactory
}
