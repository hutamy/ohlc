package main

import (
	"context"
	"log"
	"ohlc/kafka"
	"ohlc/redis"
	"ohlc/transaction"
)

func main() {
	ctx := context.Background()
	rdb, err := redis.NewRedisClient(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = rdb.Close(); err != nil {
			log.Fatalln("Error closing redis:", err)
		}
	}()

	conn, err := kafka.NewKafkaPublisher(ctx)
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatalln("Error closing connection:", err)
		}
	}()
	publisher := transaction.NewTrasactionPublisher(conn)
	publisher.Run(ctx)

	reader, err := kafka.NewKafkaConsumer()
	defer func() {
		if err = reader.Close(); err != nil {
			log.Fatalln("Error closing reader:", err)
		}
	}()

	consumer := transaction.NewTrasactionConsumer(reader, rdb)
	consumer.Run(ctx)
}
