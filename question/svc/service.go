package questionsvc

import (
	"context"
	"github.com/miraclew/tao-demo/locator"
	"github.com/miraclew/tao-demo/question"
	"github.com/labstack/echo/v4"
)

type DefaultService struct {
}

func NewService() *DefaultService {s := &DefaultService{}
	locator.Register(s.Name()+"Service", s)
	locator.Register(s.Name()+"Event", &question.EventSubscriber{Subscriber: locator.Subscriber()})

	return s
}

func (s *DefaultService) Name() string {
	return "Question"
}

func (s *DefaultService) RegisterEventHandler() {
	eh := eventHandler{s}
	eh.Register()
}

func (s *DefaultService) RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	h := handler{Service: s}
	h.RegisterRoutes(e, m...)
}
func (s *DefaultService) Create(ctx context.Context, req *question.CreateRequest) (*question.CreateResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Delete(ctx context.Context, req *question.DeleteRequest) (*question.DeleteResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Update(ctx context.Context, req *question.UpdateRequest) (*question.UpdateResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Get(ctx context.Context, req *question.GetRequest) (*question.GetResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Query(ctx context.Context, req *question.QueryRequest) (*question.QueryResponse, error) {
	panic("not implemented")
}