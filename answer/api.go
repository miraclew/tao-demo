package answer

import (
	"context"
	"time"
)

const ServiceName = "Answer"


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


type Answer struct {
	Id int64 `db:"Id"`
	QuestionID int64 `db:"QuestionID"`
	Content string `db:"Content"`
	UserId int64 `db:"UserId"`
	Likes int64 `db:"Likes"`
	CreatedAt time.Time `db:"CreatedAt"`
}


type CreateRequest struct {
	QuestionID int64
	Content string
	AudioUrl string
	AudioDuration int64
	Images string
	Positive bool
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
	Values map[string]interface{}
}

func (req *UpdateRequest) Validate() error {
	return nil
}
type UpdateResponse struct {
	Result string
}


type GetRequest struct {
	Id int64
	Filter map[string]interface{}
}

func (req *GetRequest) Validate() error {
	return nil
}
type GetResponse struct {
	Result *Answer
}


type QueryRequest struct {
	Ids []int64
	Filter map[string]interface{}
	Offset int64
	Limit int32
	Sort string
}

func (req *QueryRequest) Validate() error {
	return nil
}
type MyAnswer struct {
	Answer *Answer
	Liked bool
}


type QueryResponse struct {
	Results []*MyAnswer
}


type CreatedEvent struct {
	Data *Answer
}


type DeletedEvent struct {
	Data *Answer
}


type UpdatedEvent struct {
	Id int64
	Values map[string]interface{}
}


