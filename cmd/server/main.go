package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	
    amqp "github.com/rabbitmq/amqp091-go"
	"github.com/EdoardoPanzeri1/learn-pub-sub-starter/internal/routing"
)

func main() {
	const rabbitConnString = "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}

	chs, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create a nre channel: %v", err)
	}
	defer conn.Close()

	fmt.Println("Starting Peril server...")

	state := routing.PlayingState{
		IsPaused: true,
	}

	jsonData, err := json.Marhsal(state)
	if err != nil {
		log.Fatalf("could not marshal JSON: %v", err)
	}

	err = chs.PublishWithContext(
		context.Background()
		routing.ExchangePerilDirect,
		routing.PauseKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: jsonData,
		},
	)
	if err != nil {
		log.Fatalf("could not publish message: %v", err)
	}

	fmt.Println("Message published successfully")

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
