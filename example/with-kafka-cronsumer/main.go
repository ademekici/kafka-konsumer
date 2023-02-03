package main

import (
	"github.com/Abdulsametileri/kafka-template/pkg/config"
	"github.com/Abdulsametileri/kafka-template/pkg/kafka"
	"github.com/Abdulsametileri/kafka-template/pkg/listener"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	kafkaCfg := &config.Kafka{
		Servers: "localhost:29092",
	}
	kafkaConsumer := &config.Consumer{
		Concurrency:   10,
		Topics:        []string{"standart-topic"},
		ConsumerGroup: "standart-cg",
		Exception:     config.Exception{},
	}

	consumer, err := kafka.NewConsumer(kafkaCfg, kafkaConsumer)
	if err != nil {
		log.Fatal(err.Error())
	}

	processor := NewProcessor()
	processor.Cronsumer.Start()

	activeListenerManager := listener.NewManager()
	activeListenerManager.RegisterAndStart(consumer, processor, kafkaConsumer.Concurrency)

	<-gracefulShutdown

	activeListenerManager.Stop()
	processor.Cronsumer.Stop()
}
