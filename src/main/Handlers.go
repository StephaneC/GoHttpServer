package main

import (
	"net/http"
	"encoding/json"
)

/**
 * JSON PING response
 */
func PingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}
