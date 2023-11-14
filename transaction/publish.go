package transaction

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"ohlc/kafka"
	"path/filepath"
	"strings"

	pb "ohlc/proto"

	"google.golang.org/protobuf/proto"
)

type Transaction struct {
	Type             string `json:"type"`
	Price            int    `json:"price,string"`
	StockCode        string `json:"stock_code"`
	Quantity         int    `json:"quantity,string"`
	ExecutedQuantity int    `json:"executed_quantity,string"`
	ExecutionPrice   int    `json:"execution_price,string"`
}

type TransactionPublisher struct {
	kafkaConn *kafka.KafkaPublisherClient
}

func NewTransactionPublisher(kafkaConn *kafka.KafkaPublisherClient) *TransactionPublisher {
	return &TransactionPublisher{
		kafkaConn: kafkaConn,
	}
}

func (t *TransactionPublisher) Run(ctx context.Context) {
	// Read the transaction data from the .ndjson files in the `subsetdata` folder.
	files, err := ioutil.ReadDir("subsetdata")
	if err != nil {
		log.Fatalf("Failed to read subsetdata folder: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".ndjson" {
			return
		}

		filePath := filepath.Join("subsetdata", file.Name())
		fileData, err := ioutil.ReadFile(filepath.Clean(filePath))
		if err != nil {
			log.Printf("Failed to read file %s: %v", filePath, err)
			return
		}

		transactions := strings.Split(string(fileData), "\n")
		t.Process(ctx, transactions)
	}
}

func (t *TransactionPublisher) Process(ctx context.Context, transactions []string) {
	for _, txStr := range transactions {
		if txStr == "" {
			continue
		}

		var tx Transaction
		if err := json.Unmarshal([]byte(txStr), &tx); err != nil {
			log.Printf("Failed to parse transaction: %v", err)
			continue
		}

		ohlcMsg := &pb.Transaction{
			Type:      tx.Type,
			Price:     int32(tx.Price),
			StockCode: tx.StockCode,
			Quantity:  int32(tx.Quantity),
		}

		if tx.Type == "E" || tx.Type == "P" {
			ohlcMsg.Price = int32(tx.ExecutionPrice)
			ohlcMsg.Quantity = int32(tx.ExecutedQuantity)
		}

		ohlcBytes, err := proto.Marshal(ohlcMsg)
		if err != nil {
			log.Printf("Failed to marshal OHLC message: %v", err)
			continue
		}

		err = t.kafkaConn.Publish(ctx, ohlcBytes)
		if err != nil {
			log.Printf("Failed to publish OHLC message to Kafka: %v", err)
			continue
		}
	}
}
