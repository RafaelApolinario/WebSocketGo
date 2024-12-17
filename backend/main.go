package main

import (
	"log"
	"net/http"
	"WebSocketGo/backend/config"
	"WebSocketGo/backend/db"
	"WebSocketGo/backend/websocket"

	logrus "github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	cfg := config.LoadConfig("C:\\Users\\Teknisa Software\\go\\src\\WebSocketGo\\.env")
	logger.Info("Configurações carregadas com sucesso.")

	db.InitDB(cfg)

	go websocket.HandleBroadcast()

	http.HandleFunc("/ws", websocket.HandleConnections)

	logger.Info("Servidor WebSocket rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}