package answersvc

import (
	"context"
	"e.coding.net/miraclew/douyin/locator"
	"github.com/miraclew/tao-demo/answer"
)

type eventHandler struct {
	Service answer.Service
}

func (h eventHandler) Register() {
	// example:
	locator.AnswerEvent().HandleCreated(h.onAnswerCreated)
}

// example:
func (h eventHandler) onAnswerCreated(ctx context.Context, req *answer.CreatedEvent) error {
	return nil
}
