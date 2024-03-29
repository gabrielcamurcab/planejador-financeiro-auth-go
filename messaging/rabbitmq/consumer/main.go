package consumer

import (
	"encoding/json"
	"log"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/db/mongodb"
	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/messaging/rabbitmq"
	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/model/user"
)

func InitConsumer() error {
	connOpts := rabbitmq.ConnectionOptions{
		URL:      "amqp://guest:guest@localhost:5672",
		Exchange: "auth",
		Queue:    "create-user",
	}

	consumer, err := rabbitmq.NewConsumer(connOpts)
	if err != nil {
		return err
	}

	consumer.RegisterHandler(CreateUser)

	if err := consumer.Start(); err != nil {
		return err
	}

	log.Println("Consumidor RabbitMQ iniciado...")

	return nil
}

func CreateUser(message []byte) {
	user := user.User{}
	err := json.Unmarshal(message, &user)
	if err != nil {
		log.Printf("Erro ao decodificar o JSON: %v", err)
		return
	}

	err = mongodb.InsertUser(&user)
	if err != nil {
		log.Printf("Erro ao inserir usuário no MongoDB: %v", err)
		return
	}

	log.Println("Usuário inserido com sucesso:", user)
}
