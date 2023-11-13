package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"ohlc/config"
	"ohlc/redis"
	"ohlc/service"
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

	service := service.NewService(rdb)
	http.HandleFunc("/ohlc", func(w http.ResponseWriter, r *http.Request) {
		// Get the stock code from the query parameters
		stockCode := r.URL.Query().Get("stock_code")

		var ohlcSummary map[string]interface{}
		ohlcSummary, err = service.GetOHLC(ctx, stockCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(ohlcSummary); err != nil {
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
