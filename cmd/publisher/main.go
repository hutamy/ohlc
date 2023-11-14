package main

import (
	"context"
	"log"
	"ohlc/kafka"
	"ohlc/transaction"
)

func main() {
	ctx := context.Background()

	conn, err := kafka.NewKafkaPublisher(ctx)
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatalln("Error closing connection:", err)
		}
	}()
	publisher := transaction.NewTrasactionPublisher(conn)
	publisher.Run(ctx)
}
