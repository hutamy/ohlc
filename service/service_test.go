package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"ohlc/redis"
	"ohlc/service"
	"reflect"
	"testing"

	r "github.com/redis/go-redis/v9"
)

func TestGetOHLC(t *testing.T) {
	// Create a Redis client for testing purposes
	rdb := r.NewClient(&r.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Create a new service instance
	svc := service.NewService(&redis.RedisClient{
		Client: rdb,
	})

	// Define test cases
	testCases := []struct {
		name       string
		stockCode  string
		wantResult map[string]interface{}
		wantErr    error
	}{
		{
			name:      "valid stock code",
			stockCode: "AAPL",
			wantResult: map[string]interface{}{
				"open":  100.0,
				"high":  110.0,
				"low":   90.0,
				"close": 105.0,
			},
			wantErr: nil,
		},
		{
			name:       "invalid stock code",
			stockCode:  "INVALID",
			wantResult: map[string]interface{}{},
			wantErr:    errors.New("OHLC summary not found"),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantErr == nil {
				// Set test data
				val, _ := json.Marshal(tc.wantResult)
				_ = rdb.Set(context.Background(), tc.stockCode, string(val), 0)
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
				fmt.Println(gotResult, tc.wantResult)
				t.Errorf("GetOHLC() = %v, want %v", gotResult, tc.wantResult)
			}
		})
	}
}
