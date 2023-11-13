package main

import (
	"context"
	"encoding/json"
	"net/http"
	"ohlc/config"
	"ohlc/redis"
	"time"
)

func main() {
	cfg := config.SetConfig()
	ctx := context.Background()
	_, err := redis.NewRedisClient(ctx, cfg.RedisHost, cfg.RedisPassword, cfg.RedisPort)
	if err != nil {
		panic(err)
	}

	// Create a new web server
	http.HandleFunc("/ohlc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true}); err != nil {
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
