package http

import (
	"log"
	"net/http"

	"github.com/gabrielcamurcab/planejador-financeiro-auth-go/adapter/http/actuator"
	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/health", actuator.Health).Methods("GET")

	http.Handle("/", r)

	log.Println("Servidor HTTP iniciado na porta 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
}
