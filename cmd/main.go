package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"strconv"
)

const url = "localhost:29092"

func pubMessage(topic, message string) {
	// Create connection to broker
	brokersUrl := []string{url}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		print("1 ")
		log.Fatalln(err)
	}
	defer conn.Close()

	// Create a message for a specific topic
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Publish Message
	partition, offset, err := conn.SendMessage(msg)
	if err != nil {
		print("2 ")
		log.Fatalln(err)
	}
	fmt.Printf("Message '%s' is stored in topic (%s) / partition (%d) / offset (%d)\n", message, topic, partition, offset)
}

func subMessage(topic string, partition int32, c chan *sarama.ConsumerMessage) {
	// Create subscriber connection
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer([]string{url}, config)
	if err != nil {
		log.Fatalln(err)
	}

	// Subscribe to connection
	consumer, err := conn.ConsumePartition(topic, partition, sarama.OffsetOldest)
	for {
		msg := <-consumer.Messages()
		c <- msg
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "pub" {
			println("publish")
			if len(os.Args) > 3 {
				topic := os.Args[2]
				message := os.Args[3]
				pubMessage(topic, message)
			}
		} else if os.Args[1] == "sub" {
			println("subscribe")
			if len(os.Args) > 3 {
				topic := os.Args[2]
				partition, _ := strconv.Atoi(os.Args[3])
				signals := make(chan os.Signal)
				chanMessage := make(chan *sarama.ConsumerMessage, 256)
				go subMessage(topic, int32(partition), chanMessage)
			subLoop:
				for {
					select {
					case msg := <-chanMessage:
						log.Printf("New Message from kafka, message: %v", string(msg.Value))
					case sig := <-signals:
						if sig == os.Interrupt {
							break subLoop
						}
					}
				}
			}
		}
	}
}
