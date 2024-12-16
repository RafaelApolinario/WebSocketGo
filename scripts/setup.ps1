Write-Host "Iniciando a configuração do ambiente..." -ForegroundColor Green

# Navegar para o diretório raiz do projeto
cd .. # Volta para a raiz do projeto (WebSocketGo)

# 1. Verifica se o Go está instalado
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Erro: Go não está instalado. Instale o Go antes de continuar." -ForegroundColor Red
    exit 1
}

# 2. Verifica se o Node.js está instalado
if (-not (Get-Command node -ErrorAction SilentlyContinue)) {
    Write-Host "Erro: Node.js não está instalado. Instale o Node.js antes de continuar." -ForegroundColor Red
    exit 1
}

# 3. Configuração do backend
Write-Host "Instalando dependências do backend..." -ForegroundColor Yellow
Push-Location ./backend
go mod tidy
go mod download
Pop-Location

# 4. Configuração do frontend
Write-Host "Instalando dependências do frontend..." -ForegroundColor Yellow
Push-Location ./frontend
npm install http-server --save-dev
npm install -g lite-server
Pop-Location

Write-Host "Configuração concluída com sucesso!" -ForegroundColor Green
Write-Host "Para rodar o projeto:"
Write-Host "1. Abra dois terminais."
Write-Host "2. No primeiro, execute: cd backend && go run main.go"
Write-Host "3. No segundo, execute: cd frontend && lite-server --config bs-config.json"
