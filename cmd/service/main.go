package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"ohlc/config"
	"ohlc/redis"
	"ohlc/service"

	pb "ohlc/proto"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.SetConfig()
	ctx := context.Background()
	rdb, err := redis.NewRedisClient(ctx, cfg.RedisHost, cfg.RedisPassword, cfg.RedisPort)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = rdb.Close(); err != nil {
			log.Fatalln("Error closing redis:", err)
		}
	}()

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	service := service.NewService(rdb)
	pb.RegisterOHLCServiceServer(server, service)

	fmt.Println("Server is listening on port 50051...")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
