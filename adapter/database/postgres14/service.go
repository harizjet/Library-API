package postgres14

import (
	"database/sql"
	"fmt"
	"library/Basic-Golang-Api/adapter/database"
	"library/Basic-Golang-Api/config"
	"log"

	_ "github.com/lib/pq"
)

func NewService(cfg config.Config) database.DatabaseService {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error argument to Postgres DB")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to Postgres DB")
	}

	return &postgres14Impl{
		DbEngine: db,
	}
}
