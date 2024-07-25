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
	CONTEXTO_ERRO             = "Consultar taxa de câmbio no WebService AwesomeApi"
	ERRO                      = "Ocorreu um erro ao consultar a cotação no serviço AwesomeApi"
	TIPO_ERRO                 = "INTERNAL SERVER ERROR"
)

var (
	ErrInternalServerError = &structs.MensagemDeErro{CONTEXTO_ERRO, ERRO, TIPO_ERRO}
)

func ConsultarCotacaoDolar() (*structs.AwesomeApiResponse, *structs.MensagemDeErro) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", COTACAO_USDBRL_WEBSERVICE, nil)
	if err != nil {
		return nil, ErrInternalServerError
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer response.Body.Close()

	resposta, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, ErrInternalServerError
	}

	var awesomeApiResponse structs.AwesomeApiResponse
	err = json.Unmarshal(resposta, &awesomeApiResponse)
	if err != nil {
		return nil, ErrInternalServerError
	}

	io.Copy(os.Stdout, response.Body)
	return &awesomeApiResponse, nil
}
