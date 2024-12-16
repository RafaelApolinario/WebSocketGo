let socket;

function connectWebSocket() {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = () => console.log("Conexão WebSocket estabelecida");
  socket.onerror = (error) => console.error("Erro no WebSocket:", error);
  socket.onclose = () => {
    console.log("Conexão WebSocket encerrada. Tentando reconectar em 5 segundos...");
    setTimeout(connectWebSocket, 5000);
  };

  socket.onmessage = (event) => {
    const log = JSON.parse(event.data);
    updateLogs(log);
  };
}

function updateLogs(log) {
  const container = document.getElementById("logs-container");

  const logEntry = document.createElement("div");
  logEntry.className = "log-entry";

  logEntry.innerHTML = `
    <strong>Timestamp:</strong> ${log.timestamp}<br>
    <strong>Tipo de Evento:</strong> ${log.event_type}<br>
    <strong>Dados:</strong> ${JSON.stringify(log.data, null, 2)}
  `;

  container.appendChild(logEntry);

  container.scrollTop = container.scrollHeight;
}


connectWebSocket();


function sendEvent(eventType, eventData) {
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

document.addEventListener("click", (event) => {
  sendEvent("click", { x: event.clientX, y: event.clientY });
});

window.addEventListener("scroll", () => {
  sendEvent("scroll", { scrollY: window.scrollY });
  console.log(`Scroll detectado! Posição Y: ${window.scrollY}`);
});


document.addEventListener("keydown", (event) => {
  sendEvent("key_press", { key: event.key });
});
