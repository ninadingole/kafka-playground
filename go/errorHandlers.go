package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type DropMessageErrorHandler struct {
}

func (d *DropMessageErrorHandler) Handle(err error, _ *sarama.ConsumerMessage) {
	if err != nil {
		return
	}
}

type DlqErrorHandler struct {
	producer sarama.SyncProducer
	topic    string
}

func NewDLQMessageHandler(producer sarama.SyncProducer, topic string) *DlqErrorHandler {
	return &DlqErrorHandler{producer: producer, topic: topic}
}

func (d *DlqErrorHandler) Handle(err error, msg *sarama.ConsumerMessage) {
	if err != nil && d.producer != nil {
		_, _, err = d.producer.SendMessage(&sarama.ProducerMessage{
			Topic: d.topic,
			Value: sarama.ByteEncoder(msg.Value),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
