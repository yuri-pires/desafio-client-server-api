package structs

import "gorm.io/gorm"

type AwesomeApiResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type USDBRL struct {
	Bid string `json:"bid"`
}

type Bid struct {
	ID      int `gorm:"primaryKey"`
	Bid string
	gorm.Model
}

type MensagemDeErro struct {
	Contexto string `json:"contexto"`
	Erro     string `json:"erro"`
	Tipo     string `json:"tipo"`
}
