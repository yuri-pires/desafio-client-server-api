package structs

type AwesomeApiResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Bid string `json:"bid"`
}

type MensagemDeErro struct {
	Contexto string `json:"contexto"`
	Erro     string `json:"erro"`
	Tipo     string `json:"tipo"`
}
