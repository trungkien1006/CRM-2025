package create_consumer

import (
	"context"
	"fmt"
	"log"

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

func ConsumerInit(brokers []string, group string, topic string, config *sarama.Config, ctx context.Context) {
	fmt.Println("Consumer is initing...")

	// Tạo consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Không thể tạo consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// Bắt đầu đọc message
	go func() {
		for {
			err := consumerGroup.Consume(ctx, []string{topic}, ConsumerGroupHandler{})
			if err != nil {
				log.Printf("Lỗi trong consumer group: %v", err)
			}
			if ctx.Err() != nil {
				fmt.Println("Consumer stopped due to context cancellation.")
				return // Chấm dứt khi context bị hủy
			}
		}
	}()

	// Đợi cho đến khi context bị hủy
	<-ctx.Done()
	fmt.Println("Consumer shutting down gracefully.")
}