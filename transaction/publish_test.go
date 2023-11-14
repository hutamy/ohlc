package transaction

import (
	"context"
	"encoding/json"
	"ohlc/kafka"
	pb "ohlc/proto"
	"testing"

	"github.com/bmizerany/assert"
	"google.golang.org/protobuf/proto"
)

func TestTransactionPublisher_Process(t *testing.T) {
	ctx := context.Background()
	// create a new kafka connection
	kafkaConn, _ := kafka.NewKafkaPublisher(ctx)

	// create a new kafka reader
	kafkaReader, _ := kafka.NewKafkaConsumer()
	publisher := NewTransactionPublisher(kafkaConn)

	// create a new transaction
	transaction := `{"type":"A","order_book":"102","price":"8000","stock_code":"AAPL","quantity":"0"}`
	publisher.Process(ctx, []string{transaction})

	var tx Transaction
	json.Unmarshal([]byte(transaction), &tx)
	want := &pb.Transaction{
		Type:      tx.Type,
		Price:     int32(tx.Price),
		StockCode: tx.StockCode,
		Quantity:  int32(tx.Quantity),
	}

	if tx.Type == "E" || tx.Type == "P" {
		want.Price = int32(tx.ExecutionPrice)
		want.Quantity = int32(tx.ExecutedQuantity)
	}

	retrieveMsg, _ := kafkaReader.ReadMessage(ctx)
	got := &pb.Transaction{}
	proto.Unmarshal(retrieveMsg, got)

	assert.Equal(t, want.Type, got.Type)
	assert.Equal(t, want.Price, got.Price)
	assert.Equal(t, want.StockCode, got.StockCode)
	assert.Equal(t, want.Quantity, got.Quantity)
}
