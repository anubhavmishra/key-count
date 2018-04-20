package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type KeyCountResponse struct {
	LinksCount int `json:"links_count"`
}

type keyCountHandler struct {
	RedisClient *redis.Client
	KeyName     string
}

func (g *keyCountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var count int

	// Get key count
	redisResponse := g.RedisClient.Keys(fmt.Sprintf("%s*", g.KeyName))
	results, err := redisResponse.Result()
	if err != nil {
		log.Printf("Failed to get keys from redis: %v", err)
		http.Error(w, "Failed to get keys from redis", http.StatusInternalServerError)
	}
	count = len(results)

	response := KeyCountResponse{
		LinksCount: count,
	}
	json.NewEncoder(w).Encode(response)
	return
}

func KeyCountHandler(redisClient *redis.Client, key string) http.Handler {
	return &keyCountHandler{
		RedisClient: redisClient,
		KeyName:     key,
	}
}
