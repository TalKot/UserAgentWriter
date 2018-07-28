package main

import (
	config "WA_Training_Matcher/config"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/uber-go/kafka-client/kafka"
	consumer "gitlab.appsflyer.com/wa/go-af-kafka/consumer"
)

var c *config.Conf

func settingConsumerSettings() {
	fmt.Println("Kafka Consumer setup...")
	c = config.InitConf()
}

func startConsumer() {
	fmt.Println("Init Kafka Consumer...")

	// mapping from cluster name to list of broker ip addresses
	cluster := "wa-test"
	brokers := map[string][]string{
		cluster: c.KafkaBroker,
	}
	// mapping from topic name to cluster that has that topic
	topicClusterAssignment := map[string][]string{
		c.KafkaTopicInbound: []string{cluster},
	}
	config := &kafka.ConsumerConfig{
		TopicList: kafka.ConsumerTopicList{
			kafka.ConsumerTopic{ // Consumer Topic is a combination of topic + dead-letter-queue
				Topic: kafka.Topic{ // Each topic is a tuple of (name, clusterName)
					Name:    c.KafkaTopicInbound,
					Cluster: cluster,
				},
			},
		},
		GroupName:   c.KafkaConsumerGroup,
		Concurrency: 100, // number of go routines processing messages in parallel
	}

	config.Offsets.Initial.Offset = sarama.OffsetNewest
	// sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	consumer, consumerErr := consumer.New(brokers, topicClusterAssignment, config)
	if consumerErr != nil {
		fmt.Println("Error starting Kafka Consumer: ", consumerErr)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	fmt.Println("Starting WA Go Matcher")
	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if !ok {
				return // channel closed
			}
			saveMessage(string(msg.Value()))
			msg.Ack()
		case <-sigCh:
			consumer.Stop()
			<-consumer.Closed()
		}
	}

}
