package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/IBM/sarama"
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
	client, err := sarama.NewClient(brokers, nil)
	if err != nil {
		panic(err)
	}
	output := make(chan string)
	closed, err := consumeMsg(client, consumeTopic, output)
	if err != nil {
		panic(err)
	}

	go func() {
		// process messages from output channel
		for msg := range output {
			fmt.Println("received message: " + msg)
		}
	}()

	<-closed
}

func consumeMsg(client sarama.Client, topic string, output chan string) (chan struct{}, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, errors.New("cannot create client")
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return nil, errors.New("cannot get partitions from topic " + topic)
	}

	partitionConsumerWg := &sync.WaitGroup{}
	closing := make(chan struct{})
	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			e := client.Close()
			close(closing)
			if e != nil && e != sarama.ErrClosedClient {
				fmt.Println("cannot close client: " + e.Error())
			}
			return nil, errors.New("failed to start consumer for partition: " + string(partition))
		}

		partitionConsumerWg.Add(1)
		go func() {
			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					fmt.Println("cannot close partition consumer: " + err.Error())
				}
				partitionConsumerWg.Done()
				fmt.Println("partition consumer stopped")
			}()
			kafkaMsgs := partitionConsumer.Messages()
			for {
				select {
				case kafkaMsg := <-kafkaMsgs:
					if kafkaMsg == nil {
						fmt.Println("kafka msg is closed, stopping partition consumer")
						return
					}
					payload := string(kafkaMsg.Value) + ", partition: " + strconv.Itoa(int(partition))
					output <- payload
				case <-closing:
					fmt.Println("subsriber is closing, stopping partition consumer")
					return
				}
			}
		}()
	}

	closed := make(chan struct{})
	go func() {
		partitionConsumerWg.Wait()
		close(closed)
	}()
	return closed, nil
}
