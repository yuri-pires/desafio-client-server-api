package repository

import (
	"log"

	"github.com/yuri-pires/desafio-client-server-api/server/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	CONTEXTO_ERRO = "Inserir cotação no Sqlite"
	ERRO          = "Ocorreu um erro ao inserir a cotação atual na base dados."
	TIPO_ERRO     = "INTERNAL SERVER ERROR"
)

var (
	ErrInternalServerError = &structs.MensagemDeErro{CONTEXTO_ERRO, ERRO, TIPO_ERRO}
)

func CriarConexaoSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro criar conexão com o Sqlite %v", err)
		return nil
	}

	return db
}

func SalvarCotacao(bid string) *structs.MensagemDeErro {
	db := CriarConexaoSqlite()
	if db == nil {
		return ErrInternalServerError
	}

	db.AutoMigrate(&structs.Bid{})

	bid_model := structs.Bid{Bid: bid}

	result := db.Save(&bid_model)
	if result.RowsAffected != 1 {
		return ErrInternalServerError
	}

	return nil
}
