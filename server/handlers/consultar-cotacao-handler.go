package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yuri-pires/desafio-client-server-api/server/webservices"
)

func ConsultarCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	resposta, err := webservices.ConsultarCotacaoDolar()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(err)
		return
	}

	//Set header deve vir antes do write??
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(resposta)
}
