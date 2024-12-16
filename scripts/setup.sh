#!/bin/bash

echo "Iniciando a configuração do ambiente..."

# Navegar para o diretório raiz do projeto
cd ..

# 1. Verifica se o Go está instalado
if ! command -v go &> /dev/null
then
    echo "Erro: Go não está instalado. Instale o Go antes de continuar."
    exit 1
fi

# 2. Verifica se o Node.js está instalado
if ! command -v node &> /dev/null
then
    echo "Erro: Node.js não está instalado. Instale o Node.js antes de continuar."
    exit 1
fi

# 3. Configuração do backend
echo "Instalando dependências do backend..."
cd backend
go mod tidy
go mod download
cd ..

# 4. Configuração do frontend
echo "Instalando dependências do frontend..."
cd frontend
npm install http-server --save-dev
npm install -g lite-server
cd ..

echo "Configuração concluída com sucesso!"
echo "Para rodar o projeto:"
echo "1. Abra dois terminais."
echo "2. No primeiro, execute: cd backend && go run main.go"
echo "3. No segundo, execute: cd frontend && lite-server --config bs-config.json"
