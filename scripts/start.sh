#!/bin/bash

# Este script inicia os servidores do backend e do frontend.
# Certifique-se de que o Go e o Node.js estão instalados no sistema.

# INSTRUÇÕES:
# 1. Dê permissão de execução ao arquivo:
#    chmod +x start.sh
# 2. Para executar o script, rode:
#    ./start.sh

# Calcula o caminho para o diretório raiz do projeto
PROJECT_ROOT=$(dirname "$(pwd)") # Move um nível acima do diretório atual

echo "Iniciando os servidores do projeto..."

# Inicia o backend
echo "Iniciando o backend..."
(cd "$PROJECT_ROOT/backend" && go run main.go) &

# Inicia o frontend
echo "Iniciando o frontend..."
(cd "$PROJECT_ROOT/frontend" && npx http-server .)

echo "Os servidores foram iniciados com sucesso!"
echo "Pressione Ctrl+C para encerrar."
