package structs

type AwesomeApiResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Bid string `json:"bid"`
}

type MensagemDeErro struct {
	Contexto string `json:"contexto"`
	Tipo     string `json:"tipo"`
}

func CriarMensagemDeErro(contexto string, tipo string) *MensagemDeErro {
	return &MensagemDeErro{
		Contexto: contexto,
		Tipo:     tipo,
	}
}
