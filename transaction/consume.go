package transaction

import (
	"context"
	"log"
	"ohlc/kafka"
	"ohlc/redis"
	"ohlc/util"

	pb "ohlc/proto"

	"google.golang.org/protobuf/proto"
)

type TransactionConsumer struct {
	reader *kafka.KafkaConsumerClient
	rdb    *redis.RedisClient
}

func NewTransactionConsumer(reader *kafka.KafkaConsumerClient, rdb *redis.RedisClient) *TransactionConsumer {
	return &TransactionConsumer{
		reader: reader,
		rdb:    rdb,
	}
}

func (t *TransactionConsumer) Run(ctx context.Context) {
	for {
		msg, err := t.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Failed to read message from Kafka: %v", err)
			continue
		}

		ohlcMsg := &pb.Transaction{}
		if err := proto.Unmarshal(msg, ohlcMsg); err != nil {
			log.Printf("Failed to unmarshal OHLC message: %v", err)
			continue
		}

		t.SetCache(ctx, ohlcMsg)
	}
}

func (t *TransactionConsumer) SetCache(ctx context.Context, ohlcMsg *pb.Transaction) {
	if ohlcMsg == nil {
		return
	}

	value := &pb.Summary{}
	strValue, err := t.rdb.Get(ctx, ohlcMsg.StockCode)
	if err == nil {
		if err = proto.Unmarshal([]byte(strValue), value); err != nil {
			log.Printf("Failed to unmarshal summary: %v", err)
			return
		}
	}

	value = Calculate(value, ohlcMsg)
	bytes, err := proto.Marshal(value)
	if err != nil {
		log.Printf("Failed to marshal summary: %v", err)
	}

	if err = t.rdb.Set(ctx, ohlcMsg.StockCode, bytes); err != nil {
		log.Printf("Failed to set summary: %v", err)
	}
}

func Calculate(summary *pb.Summary, tx *pb.Transaction) *pb.Summary {
	if tx == nil {
		return summary
	}

	summary.StockCode = tx.StockCode
	if tx.Quantity == 0 {
		summary.Prev = tx.Price
	} else if tx.Quantity > 0 && tx.Type != "A" {
		if summary.Open == 0 {
			summary.Open = tx.Price
		} else {
			summary.Close = tx.Price
			summary.High = util.Max(summary.High, tx.Price)
			summary.Low = util.Min(summary.Low, tx.Price)
		}
		summary.Volume += tx.Quantity
		summary.Value += tx.Quantity * tx.Price
		summary.Average = summary.Value / summary.Volume
	}

	return summary
}
