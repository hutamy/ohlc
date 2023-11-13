package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"ohlc/redis"
	"ohlc/service"
	"reflect"
	"testing"

	pb "ohlc/proto"
)

func TestGetOHLC(t *testing.T) {
	// Create a Redis client for testing purposes
	rdb, _ := redis.NewRedisClient(context.Background())

	// Create a new service instance
	svc := service.NewService(rdb)

	// Define test cases
	testCases := []struct {
		name       string
		stockCode  *pb.StockRequest
		wantResult *pb.Summary
		wantErr    error
	}{
		{
			name:      "valid stock code",
			stockCode: &pb.StockRequest{StockCode: "AAPL"},
			wantResult: &pb.Summary{
				Open:  100.0,
				High:  110.0,
				Low:   90.0,
				Close: 105.0,
			},
			wantErr: nil,
		},
		{
			name:       "invalid stock code",
			stockCode:  &pb.StockRequest{StockCode: "INVALID"},
			wantResult: nil,
			wantErr:    errors.New("OHLC summary not found"),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr == nil {
				// Set test data
				val, _ := json.Marshal(tc.wantResult)
				_ = rdb.Set(context.Background(), tc.stockCode.StockCode, string(val))
			}

			// Call GetOHLC method
			gotResult, gotErr := svc.GetOHLC(context.Background(), tc.stockCode)

			// Check if error matches expected error
			if (gotErr == nil && tc.wantErr != nil) || (gotErr != nil && tc.wantErr == nil) || (gotErr != nil && tc.wantErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("GetOHLC() error = %v, wantErr %v", gotErr, tc.wantErr)
				return
			}

			// Check if result matches expected result
			if !reflect.DeepEqual(gotResult, tc.wantResult) {
				t.Errorf("GetOHLC() = %v, want %v", gotResult, tc.wantResult)
			}
		})
	}
}
