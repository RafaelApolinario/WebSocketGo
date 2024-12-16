package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	logrus "github.com/sirupsen/logrus"
	"encoding/json"
	"github.com/gorilla/websocket"
	"WebSocketGo/backend/db"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var logger = logrus.New()
var clients = make(map[*websocket.Conn]bool) // Lista de clientes conectados
var broadcast = make(chan map[string]interface{}) // Canal de transmissão de mensagens

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.WithError(err).Error("Erro ao estabelecer conexão WebSocket")
		return
	}
	defer conn.Close()

	clients[conn] = true
	logger.Info("Nova conexão WebSocket estabelecida")

	for {
		var msg struct {
			EventType string          `json:"event_type"`
			Data      json.RawMessage `json:"data"`
			Timestamp string          `json:"timestamp"`
		}

		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				logger.Info("Conexão encerrada pelo cliente.")
			} else {
				logger.WithError(err).Warn("Erro ao ler mensagem")
			}
			delete(clients, conn)
			break
		}

		if msg.EventType == "" || len(msg.Data) == 0 {
			logger.Warn("Mensagem inválida recebida")
			continue
		}

		_, err = db.Exec(`INSERT INTO events (event_type, data, timestamp) VALUES ($1, $2, $3)`,
			msg.EventType, msg.Data, msg.Timestamp)
		if err != nil {
			logger.WithError(err).Error("Erro ao salvar evento no banco de dados")
			continue
		}

		broadcast <- map[string]interface{}{
			"event_type": msg.EventType,
			"data":       string(msg.Data),
			"timestamp":  msg.Timestamp,
		}
	}
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				logger.WithError(err).Warn("Erro ao enviar mensagem ao cliente")
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	err := godotenv.Load("C:\\Users\\Teknisa Software\\go\\src\\WebSocketGo\\.env")
	if err != nil {
		logger.Warn("Não foi possível carregar .env. Usando variáveis de ambiente do sistema.")
	} else {
		logger.Info("Arquivo .env carregado com sucesso.")
	}

	db.InitDB()

	go handleBroadcast()

	http.HandleFunc("/ws", handleConnections)
	logger.Info("Servidor WebSocket rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}