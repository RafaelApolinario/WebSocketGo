## Dependências

### Requisitos para o Backend
Certifique-se de ter as seguintes ferramentas instaladas:

1. **Go (versão 1.23 ou superior)**  
   [Download e instalação do Go](https://go.dev/dl/).

2. **PostgreSQL (versão 9.6 ou superior)**  
   Banco de dados relacional para armazenamento dos eventos.  
   [Download e instalação do PostgreSQL](https://www.postgresql.org/download/).

3. **Bibliotecas do Go**
   No diretório do projeto, instale as seguintes dependências:
   - **`github.com/lib/pq`**: Driver PostgreSQL para Go.
     ```bash
     go get github.com/lib/pq
     ```
   - **`github.com/gorilla/websocket`**: Biblioteca para gerenciar conexões WebSocket.
     ```bash
     go get github.com/gorilla/websocket
     ```

### Requisitos para o Frontend
Certifique-se de ter um navegador moderno e um editor de texto para modificar o frontend:
- Google Chrome, Mozilla Firefox ou outro navegador compatível.
- Editor de texto como VSCode ou Sublime Text.
