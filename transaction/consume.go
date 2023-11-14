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

type TrasactionConsumer struct {
	reader *kafka.KafkaConsumerClient
	rdb    *redis.RedisClient
}

func NewTrasactionConsumer(reader *kafka.KafkaConsumerClient, rdb *redis.RedisClient) *TrasactionConsumer {
	return &TrasactionConsumer{
		reader: reader,
		rdb:    rdb,
	}
}

func (t *TrasactionConsumer) Run(ctx context.Context) {
	for {
		msg, err := t.reader.ReadMessage(context.Background())
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

func (t *TrasactionConsumer) SetCache(ctx context.Context, ohlcMsg *pb.Transaction) {
	if ohlcMsg == nil {
		return
	}

	var value *pb.Summary
	var strValue string
	var err error

	strValue, err = t.rdb.Get(ctx, ohlcMsg.StockCode)
	if err == nil {
		if err = proto.Unmarshal([]byte(strValue), value); err != nil {
			log.Printf("Failed to unmarshal summary: %v", err)
			return
		}
	}

	value = t.Calculate(value, ohlcMsg)
	if err = t.rdb.Set(ctx, ohlcMsg.StockCode, value); err != nil {
		log.Printf("Failed to set summary: %v", err)
	}
}

func (t *TrasactionConsumer) Calculate(summary *pb.Summary, tx *pb.Transaction) *pb.Summary {
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
	}

	return summary
}
