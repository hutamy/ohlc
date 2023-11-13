package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumerClient struct {
	reader *kafka.Reader
}

func NewKafkaConsumer() (*KafkaConsumerClient, error) {
	kafkaConn := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "ohlc",
	})

	return &KafkaConsumerClient{
		reader: kafkaConn,
	}, nil
}

func (k *KafkaConsumerClient) Close() error {
	return k.reader.Close()
}

func (k *KafkaConsumerClient) ReadMessage(ctx context.Context) ([]byte, error) {
	m, err := k.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return m.Value, nil
}
