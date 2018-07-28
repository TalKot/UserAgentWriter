package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/kshvakov/clickhouse"
	producer "gitlab.appsflyer.com/wa/go-af-kafka/producer"
)

type eventStructer struct {
	Date      string `json:"date"`
	Referer   string `json:"referer"`
	EventType string `json:"eventType"`
	UserAgent string `json:"userAgent"`
}

type KafkaConfig struct {
	KafkaBroker []string
	KafkaTopic  string
	KafkaMetric string
}

var c = KafkaConfig{}
var kafkaProducerConfigured = producer.KafkaProducer{}

func setupConsumer() {
	c.getConfig()
}
func writeMessage(m message) {
	producer := c.connect()
	producer.ProduceMessages(m.toString())
	fmt.Println("Wrote to kafka", m.toString())
}

func (c KafkaConfig) connect() *producer.KafkaProducer {
	//creation of new producer
	producer, producerErr := producer.New(c.KafkaBroker, c.KafkaTopic, c.KafkaMetric)
	if producerErr != nil {
		fmt.Println("Error starting Kafka Producer: ", producerErr)
	}
	return producer
}

func LoadConfiguration() KafkaConfig {
	var config KafkaConfig
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func (c *KafkaConfig) getConfig() {
	//producer config
	(*c) = LoadConfiguration()
}
