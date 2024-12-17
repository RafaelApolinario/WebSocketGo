import { connectWebSocket } from "./websocket.js";
import { updateLogs } from "./logUpdater.js";
import { setupEventListeners } from "./eventHandler.js";

// Inicializar conexão WebSocket
connectWebSocket((event) => {
  const log = JSON.parse(event.data);
  updateLogs(log);
});

// Configurar listeners de eventos
setupEventListeners();
