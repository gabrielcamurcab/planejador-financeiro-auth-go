package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/adapter/http"
	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/db/mongodb"
)

func main() {
	os.Setenv("DOTENV", "../../.env")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongodb.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer client.Disconnect(ctx)

	http.Init()
}
