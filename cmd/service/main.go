package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ohlc/config"
	"ohlc/redis"
	"time"
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

	// Create a new web server
	http.HandleFunc("/ohlc", func(w http.ResponseWriter, r *http.Request) {
		// Get the stock code from the query parameters
		stockCode := r.URL.Query().Get("stock_code")

		var ohlcSummary string
		ohlcSummary, err = rdb.Get(ctx, stockCode)
		if err != nil {
			// If the OHLC summary is not found, return an error message
			http.Error(w, "OHLC summary not found", http.StatusNotFound)
			return
		}

		var jsonOhlcSummary map[string]interface{}
		err = json.Unmarshal([]byte(ohlcSummary), &jsonOhlcSummary)
		if err != nil {
			log.Printf("Failed to parse transaction: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(jsonOhlcSummary); err != nil {
			panic(err)
		}
	})

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
