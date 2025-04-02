package main

import (
	"github.com/IBM/sarama"
	"kafka_server/create_consumer"
)

func main() {
	// Kafka config
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0 // Phiên bản Kafka
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// Kafka broker
	brokers := []string{"kafka:9093"}

	go create_consumer.ConsumerInit(brokers, "chat-group", "chat-topic", config)
	go create_consumer.ConsumerInit(brokers, "invoice-group", "invoice-topic", config)
}
