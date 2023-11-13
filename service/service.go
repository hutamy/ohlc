package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ohlc/redis"
)

type Service struct {
	rdb *redis.RedisClient
}

func NewService(rdb *redis.RedisClient) *Service {
	return &Service{rdb}
}

func (s *Service) GetOHLC(ctx context.Context, stockCode string) (map[string]interface{}, error) {
	var ohlcSummary string
	var err error

	ohlcSummary, err = s.rdb.Get(ctx, stockCode)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("OHLC summary not found")
	}

	var jsonOhlcSummary map[string]interface{}
	err = json.Unmarshal([]byte(ohlcSummary), &jsonOhlcSummary)
	if err != nil {
		log.Printf("Failed to parse transaction: %v", err)
	}

	return jsonOhlcSummary, nil
}
