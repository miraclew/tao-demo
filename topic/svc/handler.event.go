package topicsvc

import (
	"context"
	"e.coding.net/miraclew/douyin/locator"
	"github.com/miraclew/tao-demo/topic"
)

type eventHandler struct {
	Service topic.Service
}

func (h eventHandler) Register() {
	// example:
	locator.TopicEvent().HandleCreated(h.onTopicCreated)
}

// example:
func (h eventHandler) onTopicCreated(ctx context.Context, req *topic.CreatedEvent) error {
	return nil
}
