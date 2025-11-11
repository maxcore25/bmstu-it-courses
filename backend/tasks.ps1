param(
    [Parameter(Position=0)]
    [string]$Command = "help"
)

function Dev {
    Write-Host "Starting dev mode (Air hot reload)..." -ForegroundColor Cyan
    air
}

function Build {
    Write-Host "Building Go application..." -ForegroundColor Cyan
    if (-Not (Test-Path "bin")) {
        New-Item -ItemType Directory -Path "bin" | Out-Null
    }
    go build -o bin/app .
    Write-Host "Build complete: bin/app" -ForegroundColor Green
}

function Run {
    Write-Host "Running Go application..." -ForegroundColor Cyan
    go run .
}

function Help {
    Write-Host "Available commands:" -ForegroundColor Yellow
    Write-Host "  dev      - Start Air hot reload"
    Write-Host "  build    - Build Go app into ./bin/app"
    Write-Host "  run      - Run Go app (no reload)"
    Write-Host "  help     - Show this help"
}

switch ($Command.ToLower()) {
    "dev"   { Dev }
    "build" { Build }
    "run"   { Run }
    "help"  { Help }
    default {
        Write-Host "Unknown command: $Command" -ForegroundColor Red
        Help
    }
}
