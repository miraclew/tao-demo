package topic

import (
	"context"
	"time"
)

const ServiceName = "Topic"


type Service interface {
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error)
}


type Event interface {
	HandleCreated(f func(ctx context.Context, req *CreatedEvent) error)
	HandleDeleted(f func(ctx context.Context, req *DeletedEvent) error)
	HandleUpdated(f func(ctx context.Context, req *UpdatedEvent) error)
}


type Empty struct {
}


type Topic struct {
	Id int64 `db:"Id"`
	Title string `db:"Title"`
	Icon string `db:"Icon"`
	Questions int64 `db:"Questions"`
	Hot int64 `db:"Hot"`
	CreatedAt time.Time `db:"CreatedAt"`
	UpdatedAt time.Time `db:"UpdatedAt"`
}


type CreateRequest struct {
	Topic *Topic
}

func (req *CreateRequest) Validate() error {
	return nil
}
type CreateResponse struct {
	Id int64
}


type DeleteRequest struct {
	Id int64
}

func (req *DeleteRequest) Validate() error {
	return nil
}
type DeleteResponse struct {
	Result string
}


type UpdateRequest struct {
	Id int64
	Attrs map[string]interface{}
}

func (req *UpdateRequest) Validate() error {
	return nil
}
type UpdateResponse struct {
	Result string
}


type GetRequest struct {
	Id int64
}

func (req *GetRequest) Validate() error {
	return nil
}
type GetResponse struct {
	Result *Topic
}


type QueryRequest struct {
	Offset int64
	Limit int32
	Filters map[string]interface{}
	Sort string
}

func (req *QueryRequest) Validate() error {
	return nil
}
type QueryResponse struct {
	Results []*Topic
}


type CreatedEvent struct {
	Data *Topic
}


type DeletedEvent struct {
	Data *Topic
}


type UpdatedEvent struct {
	Id int64
	Attrs map[string]interface{}
}


