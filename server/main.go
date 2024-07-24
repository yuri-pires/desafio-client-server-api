package main

import (
	"log"
	"net/http"

	"github.com/yuri-pires/desafio-client-server-api/server/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao", handlers.ConsultarCotacaoHandler)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: \n %v", err)
	}
}
