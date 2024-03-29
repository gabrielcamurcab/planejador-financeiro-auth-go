package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/db/mongodb"
	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/messaging/rabbitmq/consumer"
)

func main() {
	client, err := mongodb.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Erro ao desconectar do banco de dados: %v", err)
		}
	}()

	err = consumer.InitConsumer()
	if err != nil {
		log.Fatalf("Erro ao iniciar o consumidor de mensagens RabbitMQ: %v", err)
	}

	log.Println("Serviços de mensagens incializado...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Desligando o serviço de mensagens...")

	time.Sleep(2 * time.Second)

	log.Println("Serviço de menesagens desligado")
}
