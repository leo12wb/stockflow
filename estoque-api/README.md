# estoque-api

Backend REST da aplicação **Sistema de Controle de Estoque**.

> Documentação completa: veja o [README principal](../README.md)

![Preview](../img.png)

## Stack

- **Go 1.24** + **Gin** — API REST
- **GORM** + **PostgreSQL 16** — ORM com Soft Delete
- **JWT** (golang-jwt/jwt v5) — Autenticação
- **Swagger** (swaggo/swag) — Documentação em `/swagger/index.html`
- **Viper** — Configuração via `.env`

## Estrutura

```
estoque-api/
├── cmd/server/        # main.go — entrypoint + anotações Swagger
├── config/            # Conexão GORM + PostgreSQL
├── controllers/       # Handlers HTTP
├── services/          # Regras de negócio
├── repositories/      # Queries GORM (com paginação)
├── models/            # Produto, Fornecedor, Movimentacao, Usuario
├── dto/               # Structs de request
├── middlewares/       # JWT auth middleware
├── routes/            # Registro de rotas + Swagger
├── seeders/           # Dados iniciais (idempotente)
├── utils/             # Validação, Response, Paginação
└── docs/              # Swagger gerado (não editar manualmente)
```

## Executar com Docker

```bash
docker-compose up --build -d
```

- API: http://localhost:8080
- Swagger: http://localhost:8080/swagger/index.html

## Executar localmente

```bash
go mod tidy
swag init -g cmd/server/main.go --output docs
go run ./cmd/server
```

## Variáveis de ambiente (`.env`)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=estoque
DB_SSLMODE=disable
JWT_SECRET=sua-chave-secreta
PORT=8080
ENV=development
```
