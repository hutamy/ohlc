package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var (
	broker = "localhost:9092"
	topic  = "ohlc"
)

type KafkaPublisherClient struct {
	conn *kafka.Conn
}

func NewKafkaPublisher(ctx context.Context) (*KafkaPublisherClient, error) {
	kafkaConn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
	}

	return &KafkaPublisherClient{
		conn: kafkaConn,
	}, nil
}

func (k *KafkaPublisherClient) Close() error {
	return k.conn.Close()
}

func (k *KafkaPublisherClient) Publish(ctx context.Context, data []byte) error {
	_, err := k.conn.WriteMessages(
		kafka.Message{Value: data},
	)

	return err
}
