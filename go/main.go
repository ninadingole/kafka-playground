package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"time"
)

const (
	topics  = "test"
	groupID = "group-id"
)

type ErrorHandler interface {
	Handle(error, *sarama.ConsumerMessage)
}

func main() {

	config := getConfig()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, groupID, config)
	if err != nil {
		panic(err)
	}

	handler := NewHandler(&DropMessageErrorHandler{})
	go func() {
		for {
			select {
			case <-signals:
				return
			default:
			}

			err := group.Consume(context.Background(), []string{topics}, handler)
			if err != nil {
				panic(err)
			}
		}
	}()
	fmt.Println("Consumer is running")

	<-signals
	fmt.Println("Interrupt is detected")
}

func getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_2_3_0

	config.ClientID = "test"
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Session.Timeout = 10 * time.Second
	config.Consumer.Fetch.Default = 10 * 1024 * 1024

	return config
}

type Handler struct {
	errorHandler ErrorHandler
}

func NewHandler(errorHandler ErrorHandler) *Handler {
	return &Handler{errorHandler: errorHandler}
}

func (h *Handler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *Handler) ConsumeClaim(_ sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		key := msg.Key
		value := msg.Value

		var employee Employee
		err := json.Unmarshal(value, &employee)
		if err != nil {
			if h.errorHandler == nil {
				panic(err)
			}

			h.errorHandler.Handle(err, msg)
		}

		fmt.Println(fmt.Sprintf("key: %s, value: %s", string(key), employee))
	}
	return nil
}
