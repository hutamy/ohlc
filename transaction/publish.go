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
	Type      string `json:"type"`
	Price     int    `json:"execution_price,string"`
	StockCode string `json:"stock_code"`
	Quantity  int    `json:"executed_quantity,string"`
}

type TrasactionPublisher struct {
	kafkaConn *kafka.KafkaPublisherClient
}

func NewTrasactionPublisher(kafkaConn *kafka.KafkaPublisherClient) *TrasactionPublisher {
	return &TrasactionPublisher{
		kafkaConn: kafkaConn,
	}
}

func (t *TrasactionPublisher) Run(ctx context.Context) {
	// Read the transaction data from the .ndjson files in the `subsetdata` folder.
	files, err := ioutil.ReadDir("subsetdata")
	if err != nil {
		log.Fatalf("Failed to read subsetdata folder: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".ndjson" {
			continue
		}

		filePath := filepath.Join("subsetdata", file.Name())
		fileData, err := ioutil.ReadFile(filepath.Clean(filePath))
		if err != nil {
			log.Printf("Failed to read file %s: %v", filePath, err)
			continue
		}

		transactions := strings.Split(string(fileData), "\n")
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

			ohlcBytes, err := proto.Marshal(ohlcMsg)
			if err != nil {
				log.Printf("Failed to marshal OHLC message: %v", err)
				return
			}

			err = t.kafkaConn.Publish(ctx, ohlcBytes)
			if err != nil {
				log.Printf("Failed to publish OHLC message to Kafka: %v", err)
				return
			}
		}
	}
}
