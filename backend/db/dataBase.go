package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"WebSocketGo/backend/config"
)

var db *sql.DB

func InitDB(cfg *config.Config) {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com banco: %v", err)
	}

	log.Println("Conexão ao banco estabelecida com sucesso!")

	createTable()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		event_type VARCHAR(50) NOT NULL,
		data JSONB NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	log.Println("Tabela 'events' criada (ou já existia).")
}
