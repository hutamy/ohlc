package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ohlc/redis"

	pb "ohlc/proto"
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

	var jsonOhlcSummary *pb.Summary
	err = json.Unmarshal([]byte(ohlcSummary), &jsonOhlcSummary)
	if err != nil {
		log.Printf("Failed to parse transaction: %v", err)
	}

	return jsonOhlcSummary, nil
}
