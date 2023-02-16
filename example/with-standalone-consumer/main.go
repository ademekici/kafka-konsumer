package main

import (
	"fmt"
	"github.com/Trendyol/kafka-konsumer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	consumer, _ := kafka.NewConsumer(kafka.ConsumerConfig{
		Concurrency: 1,
		Reader: kafka.ReaderConfig{
			Brokers: []string{"localhost:29092"},
			Topic:   "standart-topic",
			GroupID: "standart-cg",
		},
	})

	processor := NewProcessor()
	consumer.Consume(processor)

	fmt.Println("Consumer started...!")

	<-gracefulShutdown
}