package websocket

import (
	"encoding/json"
	"net/http"
	"WebSocketGo/backend/db"

	logrus "github.com/sirupsen/logrus"
	"github.com/gorilla/websocket"
)

var (
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	logger    = logrus.New()
	clients   = make(map[*websocket.Conn]bool)
	Broadcast = make(chan map[string]interface{})
)

const authToken = "seu_token_secreto" // Substitua por algo mais seguro em produção

type Message struct {
	EventType string          `json:"event_type"`
	Data      json.RawMessage `json:"data"`
	Timestamp string          `json:"timestamp"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != authToken {
		http.Error(w, "Token inválido ou ausente", http.StatusUnauthorized)
		logger.Warn("Tentativa de conexão com token inválido.")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.WithError(err).Error("Erro ao estabelecer conexão WebSocket")
		return
	}
	defer conn.Close()

	clients[conn] = true
	logger.Info("Nova conexão WebSocket estabelecida")

	for {
		var msg Message
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

		Broadcast <- map[string]interface{}{
			"event_type": msg.EventType,
			"data":       string(msg.Data),
			"timestamp":  msg.Timestamp,
		}
	}
}

func HandleBroadcast() {
	for {
		msg := <-Broadcast
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
