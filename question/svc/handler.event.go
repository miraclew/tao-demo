package questionsvc

import (
	"context"
	"e.coding.net/miraclew/douyin/locator"
	"github.com/miraclew/tao-demo/question"
)

type eventHandler struct {
	Service question.Service
}

func (h eventHandler) Register() {
	// example:
	locator.QuestionEvent().HandleCreated(h.onQuestionCreated)
}

// example:
func (h eventHandler) onQuestionCreated(ctx context.Context, req *question.CreatedEvent) error {
	return nil
}
