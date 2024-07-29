package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yuri-pires/desafio-client-server-api/client/arquivos"
	"github.com/yuri-pires/desafio-client-server-api/server/structs"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Ocorreu um erro ao criar a requisição com contexto %v \n", err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Ocorreu um erro ao realizar a requisição ao servidor %v \n", err)
	}
	defer res.Body.Close()

	var awesomeApiResponse structs.AwesomeApiResponse
	responseBody, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(responseBody, &awesomeApiResponse); err != nil {
		log.Fatalf("Ocorreu um erro ao processar a resposta: %v", err)
	}

	fmt.Println(string(responseBody))

	if err = arquivos.SalvarCotacao(awesomeApiResponse.USDBRL.Bid); err != nil {
		log.Fatalf("Ocorreu um erro ao salvar a cotação em arquivo de texto: %v", err)
	}
}
