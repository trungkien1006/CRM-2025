package main

import (
	"context"
	"fmt"
	"kafka_server/create_consumer"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka config
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0 // Phiên bản Kafka
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// Kafka broker
	brokers := []string{"kafka:9092"}
	// brokers := []string{"172.26.168.7:9092"}

	// Tạo WaitGroup để đợi các goroutine kết thúc
	var wg sync.WaitGroup

	// Tạo context với tín hiệu hủy
	ctx, cancel := context.WithCancel(context.Background())

	// Khởi tạo consumer goroutines
	wg.Add(2) // Có 2 goroutines cần đợi

	fmt.Println("Consumer is starting...")

	go create_consumer.ConsumerInit(brokers, "chat-group", "chat-topic", config, ctx)
	go create_consumer.ConsumerInit(brokers, "invoice-group", "invoice-topic", config, ctx)

	// Đợi tín hiệu từ hệ thống để dừng
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	// Chờ tín hiệu hủy hoặc SIGINT
	<-sigchan

	// Khi nhận tín hiệu hủy, thực hiện hủy context để dừng các goroutine
	fmt.Println("Shutting down consumers...")
	cancel()

	// Đợi tất cả các goroutine kết thúc
	wg.Wait()

	fmt.Println("Consumer is end...")
}
