package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/db/mongodb"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	http.Handle("/", r)

	log.Println("Servidor HTTP iniciado na porta 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
}
