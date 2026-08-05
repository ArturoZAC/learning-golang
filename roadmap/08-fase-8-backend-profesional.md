# Fase 8 — Backend profesional 🏗️

> **Objetivo:** Armar el proyecto final: API con estructura profesional, base de datos, tests y buenas prácticas. El cierre del roadmap.

## Lo que aprenderemos

### 1. Estructura de proyecto profesional
La separación de responsabilidades que usan los equipos reales:

```
mi-backend/
├── cmd/api/main.go        → punto de entrada, wiring
├── internal/
│   ├── handler/           → recibe HTTP, valida, responde
│   ├── service/           → lógica de negocio
│   ├── repository/        → acceso a datos (DB)
│   └── model/             → structs de dominio
└── go.mod
```

- Por qué `internal/` es una convención de Go (no exportable fuera del módulo)
- El patrón **handler → service → repository** (capas, como en NestJS pero a la Go)

### 2. Router: Go chi (o Echo/Gin)
- `github.com/go-chi/chi/v5` — router idiomatico, ligero, compatible con `net/http`
- Sub-routers, middlewares de chi (`chi.Logger`, CORS)
- La filosofía de Go: **la stdlib primero, frameworks solo cuando aportan**

### 3. Base de datos con pgx + SQLC
- **`pgx`**: el driver PostgreSQL más performante
- **SQLC**: genera código Go tipado a partir de tus queries SQL (nada de ORM mágico — SQL real + tipos reales)
- Pool de conexiones y config desde env vars

### 4. Configuración
- Variables de entorno con `os.Getenv` o paquete `envconfig`
- `.env` en desarrollo (cuidado: nunca commitear secretos)

### 5. Tests
- **`testing`** (stdlib): `func TestXxx(t *testing.T)`
- **`httptest`**: probar handlers sin levantar servidor
  ```go
  req := httptest.NewRequest("GET", "/items", nil)
  rec := httptest.NewRecorder()
  ```
- Tests de tabla (table-driven tests) — el patrón de tests favorito de Go
- Cobertura con `go test -cover`

### 6. Buenas prácticas finales
- Contexto (`context.Context`) propagado en toda la cadena
- Errores envueltos con `fmt.Errorf("...: %w", err)` y sentinel errors
- Graceful shutdown (cerrar el server sin cortar requests a medias)
- Logs estructurados (`log/slog`)
- Swagger/OpenAPI para documentar la API

## Ejercicio 🏗️ — PROYECTO FINAL: API de Tareas (TODO)

Requisitos completos:

1. **Tabla `tareas`**: id, título, descripción, completada, creada_en (PostgreSQL)
2. **Endpoints**: GET /tareas, GET /tareas/{id}, POST /tareas, PUT /tareas/{id}, DELETE /tareas/{id}
3. **Estructura de capas**: handler → service → repository
4. **Config por env vars**: DATABASE_URL, PUERTO
5. **Tests**: mínimo 5 tests table-driven de los handlers (con httptest)
6. **Middleware**: logging + CORS
7. **Validación**: POST con body vacío → 400; ID inexistente → 404
8. **Graceful shutdown** en el main

Checklist de "backend bonito" ✅:
- [ ] Compila con `go build` sin warnings
- [ ] `go vet ./...` y `go test ./...` pasan limpios
- [ ] `go run -race` sin race conditions
- [ ] La API responde JSON consistente (mismo formato de error siempre)
- [ ] Código formateado con `gofmt` (un solo `gofmt` y queda igual)

## 🏆 Al terminar esta fase: SÍ sos un aventurero de Go

Con esto cerramos el roadmap. De acá en adelante: microservicios, gRPC, Docker/K8s, observabilidad... pero eso ya es otra historia (y otra ruta).
