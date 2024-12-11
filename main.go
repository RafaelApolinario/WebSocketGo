package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Erro ao atualizar para WebSocket:", err)
        return
    }
    defer conn.Close()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Erro ao ler mensagem:", err)
            break
        }
        log.Printf("Mensagem recebida: %s", message)

        err = conn.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            log.Println("Erro ao enviar mensagem:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    log.Println("Servidor WebSocket iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
