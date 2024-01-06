package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

var (
	brokers      = []string{"kafka:9092"}
	consumeTopic = "events"
	publishTopic = "events-processed"
)

type event struct {
	ID int `json:"id"`
}

type processedEvent struct {
	ProcessedID int       `json:"processed_id"`
	Time        time.Time `json:"time"`
}

func main() {
	fmt.Println("Start publishing msg to kafka")
	producer, err := createPublisher()
	if err != nil {
		panic(err)
	}
	body := bodyFrom(os.Args)
	msg := &sarama.ProducerMessage{
		Topic: publishTopic,
		Value: sarama.ByteEncoder(body),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message sent to Kafka, paritition: %v, offset: %v\n", partition, offset)
}

func createPublisher() (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Kafka producer")
	}
	return producer, nil
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
