package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || pass == "" || name == "" {
		log.Fatalf("Erro: Uma ou mais variáveis de ambiente do banco de dados estão ausentes")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

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
