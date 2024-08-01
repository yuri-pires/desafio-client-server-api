package repository

import (
	"context"
	"log"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Sessão contínua para todas as operações usarem o mesmo Context
	tx := db.WithContext(ctx)

	if err := tx.AutoMigrate(&structs.Bid{}); err != nil {
		log.Printf("Erro ao realizar AutoMigrate %v", err)
		return ErrInternalServerError
	}

	bid_model := structs.Bid{Bid: bid}

	if err := tx.Save(&bid_model).Error; err != nil {
		log.Printf("Ocorreu um erro ao salvar o registro %v \n", err)
		return ErrInternalServerError
	}

	return nil
}
