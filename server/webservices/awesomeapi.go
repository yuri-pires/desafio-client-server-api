package webservices

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/yuri-pires/desafio-client-server-api/server/structs"
)

const (
	COTACAO_USDBRL_WEBSERVICE = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	CONTEXTO                  = "Consultar taxa de câmbio no WebService AwesomeApi"
	TIPO                      = "INTERNAL SERVER ERROR"
)

func ConsultarCotacaoDolar() (*structs.AwesomeApiResponse, *structs.MensagemDeErro) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	mensagem := structs.CriarMensagemDeErro(CONTEXTO, TIPO)

	request, err := http.NewRequestWithContext(ctx, "GET", COTACAO_USDBRL_WEBSERVICE, nil)
	if err != nil {
		return nil, mensagem
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, mensagem
	}
	defer response.Body.Close()

	resposta, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, mensagem
	}

	var awesomeApiResponse structs.AwesomeApiResponse
	err = json.Unmarshal(resposta, &awesomeApiResponse)
	if err != nil {
		return nil, mensagem
	}

	io.Copy(os.Stdout, response.Body)
	return &awesomeApiResponse, nil
}
