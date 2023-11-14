package service

import (
	"context"
	"fmt"
	"log"
	"ohlc/redis"

	pb "ohlc/proto"

	"google.golang.org/protobuf/proto"
)

type Service struct {
	pb.UnimplementedOHLCServiceServer
	rdb *redis.RedisClient
}

func NewService(rdb *redis.RedisClient) *Service {
	return &Service{
		rdb: rdb,
	}
}

func (s *Service) GetOHLC(ctx context.Context, req *pb.StockRequest) (*pb.Summary, error) {
	var ohlcSummary string
	var err error

	ohlcSummary, err = s.rdb.Get(ctx, req.StockCode)
	if err != nil {
		return nil, fmt.Errorf("OHLC summary not found")
	}

	jsonOhlcSummary := &pb.Summary{}
	err = proto.Unmarshal([]byte(ohlcSummary), jsonOhlcSummary)
	if err != nil {
		log.Printf("Failed to parse transaction: %v", err)
	}

	return jsonOhlcSummary, nil
}
