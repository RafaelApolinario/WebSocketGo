export let socket;

export function connectWebSocket(onMessageCallback) {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = () => console.log("Conexão WebSocket estabelecida");
  socket.onerror = (error) => console.error("Erro no WebSocket:", error);
  socket.onclose = () => {
    console.log("Conexão WebSocket encerrada. Tentando reconectar em 5 segundos...");
    setTimeout(() => connectWebSocket(onMessageCallback), 5000);
  };

  socket.onmessage = onMessageCallback;
}

export function sendEvent(eventType, eventData) {
  const message = {
    event_type: eventType,
    data: eventData,
    timestamp: new Date().toISOString(),
  };

  if (socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message));
    console.log("Evento enviado:", message);
  } else {
    console.warn("WebSocket não está conectado. Evento ignorado:", message);
  }
}
