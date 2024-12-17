package main

import (
	"log"
	"net/http"
	"WebSocketGo/backend/config"
	"WebSocketGo/backend/db"
	"WebSocketGo/backend/websocket"
	"os"            
	"path/filepath" 
	logrus "github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Erro ao obter o diretório de trabalho:", err)
	}

	envPath := filepath.Join(workingDir, ".env")
	cfg := config.LoadConfig(envPath)
	logger.Info("Configurações carregadas com sucesso.")

	db.InitDB(cfg)

	go websocket.HandleBroadcast()

	http.HandleFunc("/ws", websocket.HandleConnections)

	logger.Info("Servidor WebSocket rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}