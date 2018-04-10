package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/go-redis/redis"
)

const version = "0.0.1"

func main() {

	var httpBindAddr = "0.0.0.0"
	var httpPort = os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	httpAddr := fmt.Sprintf("%s:%s", httpBindAddr, httpPort)
	log.Println("Starting key-count app...")

	var redisAddr = os.Getenv("REDIS_ADDRESS")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	// Create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	mux := http.NewServeMux()
	mux.Handle("/", KeyCountHandler(redisClient))
	mux.HandleFunc("/healthz", HealthCheck)

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = LoggingHandler(mux)

	errChan := make(chan error, 10)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
