package service

import (
	"context"
	"encoding/json"
	"errors"
	"ohlc/redis"
	"reflect"
	"testing"

	pb "ohlc/proto"

	"google.golang.org/protobuf/proto"
)

func TestGetOHLC(t *testing.T) {
	// Create a Redis client for testing purposes
	rdb, _ := redis.NewRedisClient(context.Background())
	defer rdb.Close()

	// Create a new service instance
	svc := NewService(rdb)

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
				Open:  100,
				High:  110,
				Low:   90,
				Close: 105,
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
				resByte, _ := proto.Marshal(tc.wantResult)
				rdb.Set(context.Background(), tc.stockCode.StockCode, resByte)
			}

			// Call GetOHLC method
			gotResult, gotErr := svc.GetOHLC(context.Background(), tc.stockCode)

			// Check if error matches expected error
			if (gotErr == nil && tc.wantErr != nil) || (gotErr != nil && tc.wantErr == nil) || (gotErr != nil && tc.wantErr != nil && gotErr.Error() != tc.wantErr.Error()) {
				t.Errorf("GetOHLC() error = %v, wantErr %v", gotErr, tc.wantErr)
				return
			}

			res := &pb.Summary{}
			want := &pb.Summary{}
			json.Unmarshal([]byte(gotResult.String()), res)
			json.Unmarshal([]byte(tc.wantResult.String()), want)

			// Check if result matches expected results
			if !reflect.DeepEqual(res, want) {
				t.Errorf("GetOHLC() = %v, want %v", gotResult, tc.wantResult)
			}
		})
	}
}
