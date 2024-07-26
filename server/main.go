package main

import (
	"github.com/yuri-pires/desafio-client-server-api/server/repository"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /cotacao", handlers.ConsultarCotacaoHandler)

	// if err := http.ListenAndServe(":8081", mux); err != nil {
	// 	log.Fatalf("Erro ao iniciar o servidor: \n %v", err)
	// }

	repository.SalvarCotacao("12.33")
}
