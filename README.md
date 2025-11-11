# МГТУ им. Н.Э. Баумана — Курсовая работа на тему "Информационная система компьютерных курсов"

# Start

1. start

```sh
go mod init github.com/user/bmstu-it-courses/backend
```

2. Install globally

```sh
# hot reload dev server
go install github.com/air-verse/air@latest

# Swagger
go install github.com/swaggo/swag/cmd/swag@latest
```

# Commands

## Only for the first time (if app does not work)

```sh
# Init air for hot reload dev server
air init

# Init Swagger API docs
swag init --parseDependency --parseInternal
```

## Windows

```pwsh
# if you want to see all commands
.\tasks.ps1 help

# run dev server with hot reaload
.\tasks.ps1 dev
```

## Linux

```sh
make help

make dev
```
