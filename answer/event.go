package answer

import (
    "context"
    "e.coding.net/miraclew/tao/pkg/broker"
    "e.coding.net/miraclew/tao/pkg/ac"
    "encoding/json"
    "fmt"
)

// Typed event subscriber proxy
type EventSubscriber struct {
    broker.Subscriber
}


func (e *EventSubscriber) HandleHandleCreated(f func(ctx context.Context, req *CreatedEvent) error) {
    fmt.Println("event subscriber: register Answer.HandleCreated")
    _, _ = e.Subscribe("Answer.CreatedEvent", func(topic string, msg []byte) error {
        var req = new(CreatedEvent)
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("Answer"), req)
    })
}

func (e *EventSubscriber) HandleHandleDeleted(f func(ctx context.Context, req *DeletedEvent) error) {
    fmt.Println("event subscriber: register Answer.HandleDeleted")
    _, _ = e.Subscribe("Answer.DeletedEvent", func(topic string, msg []byte) error {
        var req = new(DeletedEvent)
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("Answer"), req)
    })
}

func (e *EventSubscriber) HandleHandleUpdated(f func(ctx context.Context, req *UpdatedEvent) error) {
    fmt.Println("event subscriber: register Answer.HandleUpdated")
    _, _ = e.Subscribe("Answer.UpdatedEvent", func(topic string, msg []byte) error {
        var req = new(UpdatedEvent)
        err := json.Unmarshal(msg, req)
        if err != nil {
            return err
        }
        return f(ac.NewInternal("Answer"), req)
    })
}

