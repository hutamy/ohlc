package main

import (
	"encoding/json"
	"net/http"
	"ohlc/config"
)

func main() {
	config.SetConfig()
	// Create a new web server
	http.HandleFunc("/ohlc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("ok")
	})

	// Start the web server and listen for incoming requests
	http.ListenAndServe(":8080", nil)
}
