package locator

import (
	"github.com/miraclew/tao-demo/config"
	"github.com/miraclew/tao-demo/answer"
	"github.com/miraclew/tao-demo/question"
	"github.com/miraclew/tao-demo/topic"
	"e.coding.net/miraclew/tao/pkg/broker"
	"fmt"
)

const PublisherName = "EventPublisher"
const SubscriberName = "EventSubscriber"

var registry = make(map[string]interface{})

func Register(serviceName string, svc interface{}) {
	fmt.Println("locator: register " + serviceName)
	registry[serviceName] = svc
}

func Publisher() broker.Publisher {
	return registry[PublisherName].(broker.Publisher)
}

func Subscriber() broker.Subscriber {
	return registry[SubscriberName].(broker.Subscriber)
}

func RegisterConfig(cfg *config.Config) {
	registry["Config"] = cfg
}

func Config() *config.Config {
	return registry["Config"].(*config.Config)
}


func Answer() answer.Service {
	v, ok := registry["AnswerService"]
	if !ok {
		panic("AnswerService not register")
	}
	return v.(answer.Service)
}

func AnswerEvent() answer.Event {
	v, ok := registry["AnswerEvent"]
	if !ok {
		panic("AnswerEvent not register")
	}
	return v.(answer.Event)
}

func Question() question.Service {
	v, ok := registry["QuestionService"]
	if !ok {
		panic("QuestionService not register")
	}
	return v.(question.Service)
}

func QuestionEvent() question.Event {
	v, ok := registry["QuestionEvent"]
	if !ok {
		panic("QuestionEvent not register")
	}
	return v.(question.Event)
}

func Topic() topic.Service {
	v, ok := registry["TopicService"]
	if !ok {
		panic("TopicService not register")
	}
	return v.(topic.Service)
}

func TopicEvent() topic.Event {
	v, ok := registry["TopicEvent"]
	if !ok {
		panic("TopicEvent not register")
	}
	return v.(topic.Event)
}
