package nestjsredis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func NewService(options *redis.Options) *Service {
	rdb := redis.NewClient(options)
	return &Service{
		client:       rdb,
		subscription: make(map[string]Handler),
	}
}

func (s *Service) addListener(command string, handler Handler) {
	channel := createChannel(command)
	_, ok := s.subscription[command]
	if !ok {
		s.subscription[channel.income] = handler
		go s.subscribe(channel, handler)
	}
}

func (s *Service) subscribe(channel *channel, handler Handler) {
	pubsub := s.client.Subscribe(channel.income)

	ch := pubsub.Channel()
	for msg := range ch {
		go s.handleMessage(msg, channel, handler)
	}
}

func (s *Service) handleMessage(msg *redis.Message, channel *channel, handler Handler) {
	message := &message{}
	err := json.Unmarshal([]byte(msg.Payload), message)
	if err != nil {
		fmt.Println(err)
	}
	context := Context{
		message: message,
		id:      message.Id,
		responseWriter: func(response string) {
			err := s.client.Publish(channel.outcome, response).Err()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	handler(&context)
}

func (s *Service) MessageHandler(command string, handler Handler) {
	s.addListener(command, handler)
}
