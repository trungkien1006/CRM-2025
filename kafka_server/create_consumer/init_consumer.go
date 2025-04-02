package create_consumer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)


// ConsumerGroupHandler xử lý message trong consumer group
type ConsumerGroupHandler struct{}

// Setup được gọi khi một consumer group session bắt đầu
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup được gọi khi một consumer group session kết thúc
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Nhận tin nhắn: %s từ partition %d, offset %d\n", string(message.Value), message.Partition, message.Offset)
		session.MarkMessage(message, "") // Đánh dấu message đã được xử lý
	}
	return nil
}

func ConsumerInit(brokers []string, group string, topic string, config *sarama.Config) {
	// Tạo consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Không thể tạo consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			err := consumerGroup.Consume(ctx, []string{topic}, ConsumerGroupHandler{})
			if err != nil {
				log.Printf("Lỗi trong consumer group: %v", err) // Log lỗi nhưng không dừng app
			}

			if ctx.Err() != nil {
				return // Chỉ thoát nếu context bị huỷ (shutdown)
			}

			log.Println("💡 Rebalance xảy ra, consumer đang chạy lại...")
		}
	}()

	// Xử lý tín hiệu shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	<-sigchan

	fmt.Println("Shutting down consumer...")
	cancel()
	wg.Wait()
}