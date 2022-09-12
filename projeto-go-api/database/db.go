package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=postgres dbname=postgres port=63903 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados", err)
	}
}
