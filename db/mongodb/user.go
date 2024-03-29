package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/model/user"
)

func InsertUser(user *user.User) error {
	client, err := Connect()
	if err != nil {
		return fmt.Errorf("erro ao conectar ao MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("go").Collection("users")

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário no MongoDB: %v", err)
	}

	log.Printf("Usuário inserido no MongoDB: %+v", user)
	return nil
}
