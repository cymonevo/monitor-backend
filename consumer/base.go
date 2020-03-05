package consumer

import "github.com/cymonevo/monitor-backend/internal/mq"

type BaseConsumerHandler interface {
	GetConsumers() []mq.BaseConsumer
}
