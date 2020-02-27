package question

import (
	"context"
	"time"
)

const ServiceName = "Question"
type QuestionType int

func (v QuestionType) String() string {
	switch v {
	case QuestionTypeText:
		return "Text"
	case QuestionTypeAudio:
		return "Audio"
	default:
		return "Unknown"
	}
}

const (
QuestionTypeText QuestionType = 0
QuestionTypeAudio QuestionType = 1
)

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


type Question struct {
	Id int64
	UserId int64
	TopicId int64
	Title string
	Content string
	Type QuestionType
	Answers int64
	CreatedAt time.Time
	UpdatedAt time.Time
}


type CreateRequest struct {
	Title string
	Content string
	Type QuestionType
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
	result string
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
	Filter map[string]interface{}
}

func (req *GetRequest) Validate() error {
	return nil
}
type GetResponse struct {
	Result *Question
}


type QueryRequest struct {
	Offset int64
	Limit int32
	Filter map[string]interface{}
	Sort string
}

func (req *QueryRequest) Validate() error {
	return nil
}
type QueryResponse struct {
	Results []*Question
}


type CreatedEvent struct {
	Data *Question
}


type DeletedEvent struct {
	Data *Question
}


type UpdatedEvent struct {
	Id int64
	Attrs map[string]interface{}
}


