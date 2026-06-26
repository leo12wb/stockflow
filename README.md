# Sistema de Controle de Estoque

Sistema web completo para gerenciamento de estoque, produtos, fornecedores e movimentações, com autenticação JWT, paginação server-side e documentação Swagger.

![Preview do Sistema](img.png)

---

## Stacks Utilizadas

### Backend — `estoque-api`

| Tecnologia | Versão | Função |
|---|---|---|
| **Go** | 1.24 | Linguagem principal |
| **Gin** | v1.10 | Framework HTTP / roteamento |
| **GORM** | v1.25 | ORM + Soft Delete |
| **PostgreSQL** | 16 | Banco de dados relacional |
| **JWT** | golang-jwt/jwt v5 | Autenticação via Bearer Token |
| **Swagger** | swaggo/swag | Documentação automática da API |
| **Viper** | v1.19 | Configuração via `.env` |
| **Zap** | uber-go/zap | Logs estruturados |

### Frontend — `estoque-front`

| Tecnologia | Versão | Função |
|---|---|---|
| **Vue 3** | v3.4 | Framework reativo (Composition API) |
| **TypeScript** | v5.4 | Tipagem estática |
| **Vite** | v5.3 | Build tool e dev server |
| **PrimeVue 4** | v4.0 | Biblioteca de componentes UI |
| **Pinia** | v2.1 | Gerenciamento de estado |
| **Axios** | v1.7 | Cliente HTTP com interceptor JWT |
| **Tailwind CSS** | v4.0 | Utilitários CSS (utilities only) |
| **Vue Router** | v4.3 | Roteamento SPA com guards |

### Infraestrutura

| Tecnologia | Função |
|---|---|
| **Docker** | Containerização de todos os serviços |
| **Docker Compose** | Orquestração multi-container |
| **Nginx** | Serve o build estático do frontend |
| **Adminer** | Interface web para o banco de dados |

---

## Arquitetura

```
Sistema de Estoque/
├── estoque-api/          # Backend Go
│   ├── cmd/server/       # Entrypoint + anotações Swagger
│   ├── config/           # Conexão com banco de dados
│   ├── controllers/      # Handlers HTTP (Gin)
│   ├── services/         # Regras de negócio
│   ├── repositories/     # Acesso ao banco (GORM)
│   ├── models/           # Entidades (Produto, Fornecedor, Movimentação, Usuário)
│   ├── dto/              # Data Transfer Objects
│   ├── middlewares/      # Autenticação JWT
│   ├── routes/           # Definição de rotas
│   ├── seeders/          # Dados iniciais
│   ├── utils/            # Validação, respostas, paginação
│   └── docs/             # Swagger gerado automaticamente
│
└── estoque-front/        # Frontend Vue 3
    └── src/
        ├── api/          # Axios + chamadas à API
        ├── stores/       # Pinia (auth)
        ├── router/       # Rotas + guards
        ├── components/   # AppLayout (sidebar)
        └── views/        # Páginas: Login, Dashboard, Produtos, Fornecedores, Movimentações, Relatórios
```

---

## Funcionalidades

### Autenticação
- Login com e-mail e senha
- Token JWT com validade configurável
- Guard de rotas no frontend (redireciona para `/login` se não autenticado)
- Perfis de acesso: `administrador`, `gerente`, `estoquista`

### Produtos
- CRUD completo com validação inline (estilo Laravel)
- Soft Delete + restauração
- Busca por nome, código de barras ou categoria
- Indicador visual de estoque baixo (quando `quantidade <= estoque_minimo`)
- Paginação server-side

### Fornecedores
- CRUD completo com validação de CNPJ único
- Soft Delete + restauração
- Paginação server-side

### Movimentações
- Registro de entradas, saídas e ajustes de estoque
- Atualização automática do estoque do produto via transação SQL
- Validação de estoque suficiente para saídas
- Paginação server-side

### Relatórios
- Inventário geral
- Produtos com estoque baixo
- Entradas e saídas filtradas
- Produtos mais vendidos
- Histórico completo de movimentações

### Dashboard
- Cards com totais (produtos, fornecedores, movimentações, estoque baixo)
- Lista de produtos com estoque crítico
- Últimas movimentações registradas

---

## Como executar

### Pré-requisitos
- [Docker](https://www.docker.com/) + Docker Compose instalados

### 1. Clonar e subir os containers

```bash
# Na pasta estoque-api
docker-compose up --build -d
```

### 2. Acessar os serviços

| Serviço | URL |
|---|---|
| **Frontend** | http://localhost:3000 |
| **API REST** | http://localhost:8080 |
| **Swagger UI** | http://localhost:8080/swagger/index.html |
| **Adminer** (banco) | http://localhost:8081 |

### 3. Credenciais padrão (seeders)

| E-mail | Senha | Perfil |
|---|---|---|
| admin@estoque.com | admin123 | Administrador |
| gerente@estoque.com | admin123 | Gerente |
| estoquista@estoque.com | admin123 | Estoquista |

---

## API — Principais Endpoints

### Autenticação
```
POST /auth/login       → Autenticar usuário
POST /auth/register    → Registrar novo usuário
GET  /auth/me          → Dados do usuário logado
```

### Produtos
```
GET    /produtos                    → Listar (paginado: ?page=1&per_page=10&search=)
POST   /produtos                    → Criar
PUT    /produtos/:id                → Atualizar
DELETE /produtos/:id                → Soft delete
DELETE /produtos/:id/force          → Excluir permanentemente
PATCH  /produtos/:id/restaurar      → Restaurar
```

### Fornecedores
```
GET    /fornecedores                → Listar (paginado: ?page=1&per_page=10)
POST   /fornecedores                → Criar
PUT    /fornecedores/:id            → Atualizar
DELETE /fornecedores/:id            → Soft delete
PATCH  /fornecedores/:id/restaurar  → Restaurar
```

### Movimentações
```
GET  /movimentacoes    → Listar (paginado: ?page=1&per_page=10)
POST /movimentacoes    → Registrar (ENTRADA | SAIDA | AJUSTE)
```

### Relatórios
```
GET /relatorios/inventario
GET /relatorios/estoque-baixo
GET /relatorios/entradas
GET /relatorios/saidas
GET /relatorios/mais-vendidos
GET /relatorios/movimentacoes
```

> Todos os endpoints (exceto login) exigem header: `Authorization: Bearer {token}`

---

## Paginação

A API retorna o seguinte formato para listagens paginadas:

```json
{
  "success": true,
  "data": [ ... ],
  "total": 50,
  "page": 1,
  "per_page": 10,
  "last_page": 5
}
```

---

## Validação

O backend usa um sistema de validação inline inspirado no Laravel:

```go
utils.Validate(c, &req, utils.Rules{
    "nome":          "required",
    "email":         "required|email|unique:usuarios,email",
    "preco_compra":  "required|numeric",
    "fornecedor_id": "nullable|exists:fornecedores,id",
})
```

Regras disponíveis: `required`, `nullable`, `string`, `numeric`, `boolean`, `email`, `min:N`, `max:N`, `in:a,b,c`, `unique:tabela,coluna`, `exists:tabela,coluna`

---

## Desenvolvimento local (sem Docker)

### Backend
```bash
cd estoque-api
cp .env.example .env   # configure as variáveis
go mod tidy
swag init -g cmd/server/main.go --output docs
go run ./cmd/server
```

### Frontend
```bash
cd estoque-front
npm install
npm run dev            # http://localhost:5173
```
