# StockFlow

> Full-stack inventory management system built with Go, Gin, PostgreSQL and Vue 3.

**[🇧🇷 Versão em Português](README.pt.md)**

![StockFlow Preview](img.png)

---

## Tech Stack

### Backend — `estoque-api`

| Technology | Version | Role |
|---|---|---|
| **Go** | 1.24 | Main language |
| **Gin** | v1.10 | HTTP framework / routing |
| **GORM** | v1.25 | ORM with Soft Delete |
| **PostgreSQL** | 16 | Relational database |
| **JWT** | golang-jwt/jwt v5 | Bearer token authentication |
| **Swagger** | swaggo/swag | Auto-generated API docs |
| **Viper** | v1.19 | `.env` configuration |
| **Zap** | uber-go/zap | Structured logging |

### Frontend — `estoque-front`

| Technology | Version | Role |
|---|---|---|
| **Vue 3** | v3.4 | Reactive framework (Composition API) |
| **TypeScript** | v5.4 | Static typing |
| **Vite** | v5.3 | Build tool & dev server |
| **PrimeVue 4** | v4.0 | UI component library |
| **Pinia** | v2.1 | State management |
| **Axios** | v1.7 | HTTP client with JWT interceptor |
| **Tailwind CSS** | v4.0 | CSS utilities (utilities only) |
| **Vue Router** | v4.3 | SPA routing with navigation guards |

### Infrastructure

| Technology | Role |
|---|---|
| **Docker** | Containerization of all services |
| **Docker Compose** | Multi-container orchestration |
| **Nginx** | Serves the frontend static build |
| **Adminer** | Web-based database UI |

---

## Project Structure

```
stockflow/
├── estoque-api/           # Go backend
│   ├── cmd/server/        # Entrypoint + Swagger annotations
│   ├── config/            # Database connection
│   ├── controllers/       # HTTP handlers (Gin)
│   ├── services/          # Business logic
│   ├── repositories/      # GORM queries (with pagination)
│   ├── models/            # Entities: Product, Supplier, Movement, User
│   ├── dto/               # Request structs
│   ├── middlewares/       # JWT auth middleware
│   ├── routes/            # Route registration + Swagger
│   ├── seeders/           # Idempotent seed data
│   ├── utils/             # Validation, response helpers, pagination
│   └── docs/              # Auto-generated Swagger files
│
└── estoque-front/         # Vue 3 frontend
    └── src/
        ├── api/           # Axios instance + all API calls
        ├── stores/        # Pinia store (auth)
        ├── router/        # Routes + beforeEach guard
        ├── components/    # AppLayout (sidebar)
        └── views/         # Pages: Login, Dashboard, Products, Suppliers, Movements, Reports
```

---

## Features

### Authentication
- Email and password login
- JWT token with configurable expiration
- Frontend route guards (redirects to `/login` when unauthenticated)
- Access roles: `administrador`, `gerente`, `estoquista`

### Products
- Full CRUD with Laravel-style inline validation
- Soft Delete + restore
- Search by name, barcode or category
- Visual low-stock indicator (`quantity <= minimum_stock`)
- Server-side pagination

### Suppliers
- Full CRUD with unique CNPJ validation
- Soft Delete + restore
- Server-side pagination

### Stock Movements
- Record entries, exits and adjustments
- Automatic stock update via SQL transaction
- Sufficient stock validation for exits
- Server-side pagination

### Reports
- Full inventory
- Low-stock products
- Filtered entries and exits
- Best-selling products
- Complete movement history

### Dashboard
- Summary cards (products, suppliers, movements, low stock)
- Critical stock product list
- Latest movements

---

## Getting Started

### Prerequisites
- [Docker](https://www.docker.com/) + Docker Compose

### 1. Clone and start

```bash
cd estoque-api
docker-compose up --build -d
```

### 2. Access the services

| Service | URL |
|---|---|
| **Frontend** | http://localhost:3000 |
| **REST API** | http://localhost:8080 |
| **Swagger UI** | http://localhost:8080/swagger/index.html |
| **Adminer** | http://localhost:8081 |

### 3. Default credentials (seeded)

| Email | Password | Role |
|---|---|---|
| admin@estoque.com | admin123 | Administrator |
| gerente@estoque.com | admin123 | Manager |
| estoquista@estoque.com | admin123 | Stock clerk |

---

## API Reference

### Auth
```
POST /auth/login       → Authenticate user
POST /auth/register    → Register new user
GET  /auth/me          → Current user data
```

### Products
```
GET    /produtos                    → List (paginated: ?page=1&per_page=10&search=)
POST   /produtos                    → Create
PUT    /produtos/:id                → Update
DELETE /produtos/:id                → Soft delete
DELETE /produtos/:id/force          → Permanent delete
PATCH  /produtos/:id/restaurar      → Restore
```

### Suppliers
```
GET    /fornecedores                → List (paginated: ?page=1&per_page=10)
POST   /fornecedores                → Create
PUT    /fornecedores/:id            → Update
DELETE /fornecedores/:id            → Soft delete
PATCH  /fornecedores/:id/restaurar  → Restore
```

### Movements
```
GET  /movimentacoes    → List (paginated: ?page=1&per_page=10)
POST /movimentacoes    → Register (ENTRADA | SAIDA | AJUSTE)
```

### Reports
```
GET /relatorios/inventario
GET /relatorios/estoque-baixo
GET /relatorios/entradas
GET /relatorios/saidas
GET /relatorios/mais-vendidos
GET /relatorios/movimentacoes
```

> All endpoints (except login) require the header: `Authorization: Bearer {token}`

---

## Pagination

Paginated endpoints return:

```json
{
  "success": true,
  "data": [ "..." ],
  "total": 50,
  "page": 1,
  "per_page": 10,
  "last_page": 5
}
```

---

## Validation

The backend uses a Laravel-inspired inline validation system:

```go
utils.Validate(c, &req, utils.Rules{
    "nome":          "required",
    "email":         "required|email|unique:usuarios,email",
    "preco_compra":  "required|numeric",
    "fornecedor_id": "nullable|exists:fornecedores,id",
})
```

Available rules: `required`, `nullable`, `string`, `numeric`, `boolean`, `email`, `min:N`, `max:N`, `in:a,b,c`, `unique:table,column`, `exists:table,column`

---

## Local Development (without Docker)

### Backend
```bash
cd estoque-api
cp .env.example .env
go mod tidy
swag init -g cmd/server/main.go --output docs
go run ./cmd/server
```

### Frontend
```bash
cd estoque-front
npm install
npm run dev   # http://localhost:5173
```

---

## License

MIT
