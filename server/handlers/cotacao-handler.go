package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yuri-pires/desafio-client-server-api/server/webservices"
)

func ConsultarCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	resposta, erro := webservices.ConsultarCotacaoDolar()
	if erro != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(erro)
		return
	}

	//Set header devem vir antes do write??
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(resposta)
}
