import { sendEvent } from "./websocket.js";

export function setupEventListeners() {
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
}
