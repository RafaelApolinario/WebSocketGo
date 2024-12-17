
export function updateLogs(log) {
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
  