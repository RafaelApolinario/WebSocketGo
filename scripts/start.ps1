Write-Host "Iniciando os servidores do projeto..." -ForegroundColor Green

# Define o caminho absoluto do diretório raiz
$projectRoot = (Get-Item $PSScriptRoot).Parent.FullName

# Inicia o backend
Write-Host "Iniciando o backend..." -ForegroundColor Yellow
Start-Process -NoNewWindow -FilePath "powershell.exe" `
    -ArgumentList "-Command", "`"cd '$projectRoot\backend'; go run main.go`""

# Inicia o frontend usando lite-server instalado localmente
Write-Host "Iniciando o frontend..." -ForegroundColor Yellow
Start-Process -NoNewWindow -FilePath "powershell.exe" `
    -ArgumentList "-Command", "`"cd '$projectRoot'; npx lite-server --config bs-config.json`""

Write-Host "Os servidores foram iniciados com sucesso!" -ForegroundColor Green
Write-Host "Acesse o frontend em: http://127.0.0.1:8081/frontend/"
