# inventory-management-api-go

API REST para gestión de inventario en Go, aplicando Domain-Driven Design (DDD), Clean Architecture y buenas prácticas de backend profesional.

## Stack

- Go 1.22
- Gin (HTTP framework)
- PostgreSQL
- GORM
- godotenv
- Docker + Docker Compose

## Arquitectura

El proyecto está organizado por capas:

- `internal/domain`: Entidades y contratos de repositorio.
- `internal/application`: Casos de uso/servicios y reglas de negocio.
- `internal/infrastructure`: Implementaciones de base de datos, repositorios y migraciones.
- `internal/interfaces/http`: Handlers, DTOs, middleware y rutas.

Los handlers no contienen lógica de negocio; delegan completamente a los servicios de aplicación.

## Estructura

```text
inventory-management-api-go/
├── cmd/
│   └── server/
│       └── main.go
├── docs/
│   └── endpoints.http
├── internal/
│   ├── domain/
│   │   ├── category/
│   │   │   ├── entity.go
│   │   │   └── repository.go
│   │   ├── product/
│   │   │   ├── entity.go
│   │   │   └── repository.go
│   │   └── inventory/
│   │       ├── entity.go
│   │       └── repository.go
│   ├── application/
│   │   ├── category_service.go
│   │   ├── product_service.go
│   │   ├── inventory_service.go
│   │   ├── errors.go
│   │   └── tx.go
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── postgres.go
│   │   │   ├── migrations.go
│   │   │   └── tx_manager.go
│   │   └── persistence/
│   │       ├── category_repository_gorm.go
│   │       ├── product_repository_gorm.go
│   │       └── inventory_repository_gorm.go
│   └── interfaces/
│       └── http/
│           ├── handlers/
│           │   ├── category_handler.go
│           │   ├── product_handler.go
│           │   └── inventory_handler.go
│           ├── dto/
│           │   ├── category_dto.go
│           │   ├── product_dto.go
│           │   └── inventory_dto.go
│           ├── middleware/
│           │   └── error_middleware.go
│           └── routes.go
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## Reglas de negocio implementadas

- No permite SKU duplicado.
- No permite stock negativo.
- No permite cantidades menores o iguales a cero en movimientos.
- No permite eliminar una categoría con productos asociados.
- No permite eliminar un producto con movimientos registrados.
- Actualiza stock automáticamente al registrar movimientos `INBOUND` y `OUTBOUND`.

## Variables de entorno

Copiar `.env.example` a `.env` y ajustar valores:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=inventory_db
DB_SSLMODE=disable
```

## Ejecutar con Docker Compose

```bash
docker compose up --build
```

API disponible en `http://localhost:8080`.

## Ejecutar local sin Docker

1. Levantar PostgreSQL.
2. Configurar `.env`.
3. Instalar dependencias y ejecutar:

```bash
go mod tidy
go run ./cmd/server
```

## Endpoints mínimos

### Categorías

- `POST /api/v1/categories`
- `GET /api/v1/categories`
- `GET /api/v1/categories/:id`
- `PUT /api/v1/categories/:id`
- `DELETE /api/v1/categories/:id`

### Productos

- `POST /api/v1/products`
- `GET /api/v1/products`
- `GET /api/v1/products/:id`
- `PUT /api/v1/products/:id`
- `DELETE /api/v1/products/:id`

### Inventario

- `GET /api/v1/products/:id/stock`
- `POST /api/v1/inventory/inbound`
- `POST /api/v1/inventory/outbound`
- `GET /api/v1/products/:id/movements`

## Ejemplos curl

### 1. Crear una categoría

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Consumer electronics"
  }'
```

### 2. Crear un producto

```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Wireless Mouse",
    "description": "2.4Ghz ergonomic mouse",
    "sku": "MSE-001",
    "price": 29.90,
    "stock": 10,
    "category_id": 1
  }'
```

### 3. Registrar entrada de stock

```bash
curl -X POST http://localhost:8080/api/v1/inventory/inbound \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "quantity": 20,
    "description": "Initial stock load"
  }'
```

### 4. Registrar salida de stock

```bash
curl -X POST http://localhost:8080/api/v1/inventory/outbound \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "quantity": 3,
    "description": "Order #A-1024"
  }'
```

### 5. Consultar stock

```bash
curl -X GET http://localhost:8080/api/v1/products/1/stock
```

### 6. Listar movimientos de inventario

```bash
curl -X GET http://localhost:8080/api/v1/products/1/movements
```

## Formato de errores

Respuesta estándar:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "quantity must be greater than zero"
  }
}
```

## Colección de endpoints

Se incluye el archivo [`docs/endpoints.http`](docs/endpoints.http) para probar todos los endpoints desde VS Code REST Client o herramientas compatibles.
