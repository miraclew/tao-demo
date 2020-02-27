package topicsvc

import (
	"context"
	"github.com/miraclew/tao-demo/locator"
	"github.com/miraclew/tao-demo/topic"
	"github.com/labstack/echo/v4"
)

type DefaultService struct {
}

func NewService() *DefaultService {s := &DefaultService{}
	locator.Register(s.Name()+"Service", s)
	locator.Register(s.Name()+"Event", &topic.EventSubscriber{Subscriber: locator.Subscriber()})

	return s
}

func (s *DefaultService) Name() string {
	return "Topic"
}

func (s *DefaultService) RegisterEventHandler() {
	eh := eventHandler{s}
	eh.Register()
}

func (s *DefaultService) RegisterRouter(e *echo.Echo, m ...echo.MiddlewareFunc) {
	h := handler{Service: s}
	h.RegisterRoutes(e, m...)
}
func (s *DefaultService) Create(ctx context.Context, req *topic.CreateRequest) (*topic.CreateResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Delete(ctx context.Context, req *topic.DeleteRequest) (*topic.DeleteResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Update(ctx context.Context, req *topic.UpdateRequest) (*topic.UpdateResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Get(ctx context.Context, req *topic.GetRequest) (*topic.GetResponse, error) {
	panic("not implemented")
}
func (s *DefaultService) Query(ctx context.Context, req *topic.QueryRequest) (*topic.QueryResponse, error) {
	panic("not implemented")
}