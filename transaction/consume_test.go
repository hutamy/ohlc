package transaction

import (
	"context"
	"testing"

	"ohlc/kafka"
	pb "ohlc/proto"
	"ohlc/redis"

	"github.com/bmizerany/assert"
	"google.golang.org/protobuf/proto"
)

func TestCalculate(t *testing.T) {
	// create a summary with some initial values
	summary := &pb.Summary{
		StockCode: "AAPL",
		Prev:      100,
		Open:      100,
		High:      100,
		Low:       100,
		Close:     100,
		Average:   100,
		Value:     100,
		Volume:    100,
	}

	// create a new transaction with a higher price
	tx := &pb.Transaction{
		StockCode: "AAPL",
		Type:      "A",
		Quantity:  0,
		Price:     110,
	}

	// calculate the new summary
	newSummary := Calculate(summary, tx)

	// verify that the new summary has been updated correctly
	assert.Equal(t, "AAPL", newSummary.StockCode)
	assert.Equal(t, int32(100), newSummary.Open)
	assert.Equal(t, int32(100), newSummary.High)
	assert.Equal(t, int32(100), newSummary.Low)
	assert.Equal(t, int32(100), newSummary.Close)
	assert.Equal(t, int32(100), newSummary.Volume)
	assert.Equal(t, int32(100), newSummary.Value)
	assert.Equal(t, int32(100), newSummary.Average)
	assert.Equal(t, int32(110), newSummary.Prev)
}

func TestSetCache(t *testing.T) {
	// create a new redis client
	ctx := context.Background()
	rdb, _ := redis.NewRedisClient(ctx)

	// create a new kafka readerr
	kafkaReader, _ := kafka.NewKafkaConsumer()

	consumer := NewTransactionConsumer(kafkaReader, rdb)

	// create a new transaction
	txMessage := &pb.Transaction{
		StockCode: "AAPL",
		Type:      "A",
		Quantity:  0,
		Price:     110,
	}
	txSummarry := Calculate(&pb.Summary{}, txMessage)
	want, _ := proto.Marshal(txSummarry)

	// set the cache
	consumer.SetCache(ctx, txMessage)

	// get the cached value
	got, _ := rdb.Get(ctx, txMessage.StockCode)
	assert.Equal(t, string(want), got)
}
