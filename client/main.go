package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/cotacao", nil)
	if err != nil {
		log.Fatalf("Ocorreu um erro ao criar a requisição com contexto %v \n", err)
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Ocorreu um erro ao realizar a requisição ao servidor %v \n", err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
