# Guía de Migración a Arquitectura Hexagonal

## Resumen de Cambios

Se ha refactorizado completamente el proyecto para implementar **Arquitectura Hexagonal** siguiendo los principios **SOLID**, **KISS** y **YAGNI**.

## Cambios Realizados

### 1. Nueva Estructura de Directorios

```
Antes:                          Después:
├── main.go                     ├── cmd/api/main.go
├── handlers/handlers.go        ├── internal/adapters/primary/http/
├── routers/                    │   ├── user_handler.go
├── bd/                         │   ├── tweet_handler.go
├── models/                     │   ├── relation_handler.go
├── middlew/                    │   └── router.go
├── jwt/                        ├── internal/adapters/secondary/
                                │   ├── database/mongodb/
                                │   ├── auth/
                                │   └── storage/
                                ├── internal/core/
                                │   ├── domain/
                                │   ├── ports/
                                │   └── services/
                                ├── pkg/middleware/
                                └── config/
```

### 2. Separación de Responsabilidades

#### Antes (Monolítico):
- Lógica de negocio mezclada con HTTP
- Acceso directo a base de datos desde handlers
- Dependencias hardcodeadas
- Testing difícil

#### Después (Hexagonal):
- **Core**: Lógica de negocio pura
- **Adapters**: Infraestructura (HTTP, DB, Auth)
- **Ports**: Interfaces que definen contratos
- **Dependency Injection**: Flexibilidad y testabilidad

### 3. Mapeo de Archivos Antiguos vs Nuevos

| Archivo Antiguo | Nuevo Equivalente | Descripción |
|----------------|-------------------|-------------|
| `main.go` | `cmd/api/main.go` | Punto de entrada refactorizado |
| `handlers/handlers.go` | `internal/adapters/primary/http/router.go` | Router HTTP |
| `routers/*.go` | `internal/adapters/primary/http/*_handler.go` | Handlers HTTP |
| `bd/*.go` | `internal/adapters/secondary/database/mongodb/*` | Repositorios |
| `models/*.go` | `internal/core/domain/*.go` | Entidades de dominio |
| `middlew/*.go` | `pkg/middleware/*.go` | Middlewares HTTP |
| `jwt/jwt.go` | `internal/adapters/secondary/auth/jwt_generator.go` | Generación JWT |

### 4. Principales Beneficios

#### ✅ Testabilidad
```go
// Antes: Testing difícil
func TestSomething(t *testing.T) {
    // Necesita base de datos real
    // Dependencias hardcodeadas
}

// Después: Testing fácil con mocks
func TestUserService_Register(t *testing.T) {
    userRepo := NewMockUserRepository()
    service := services.NewUserService(userRepo, ...)
    // Test aislado y rápido
}
```

#### ✅ Flexibilidad
```go
// Cambiar de MongoDB a PostgreSQL
userRepo := postgresql.NewUserRepository(conn)
// Solo cambias la implementación, no la lógica
```

#### ✅ Mantenibilidad
- Código organizado por dominios
- Responsabilidades claras
- Fácil localización de bugs

### 5. Nuevas Características

#### Configuration Management
```go
// Configuración centralizada con variables de entorno
cfg := config.Load()
```

#### Dependency Injection
```go
// Inyección de dependencias clara
service := services.NewUserService(repo, hasher, tokenGen, storage)
```

#### Interfaces para Testing
```go
// Interfaces facilitan testing y mocking
type UserRepository interface {
    Create(ctx context.Context, user *domain.User) (*domain.UserCreated, error)
    // ...
}
```

## Cómo Usar la Nueva Versión

### 1. Ejecutar con Docker
```bash
make docker-run
```

### 2. Ejecutar Localmente
```bash
# Configurar variables de entorno
export MONGODB_URI="mongodb://localhost:27017/twittor"
export JWT_SECRET="your-secret-key"

# Ejecutar
make run
# o
go run cmd/api/main.go
```

### 3. Testing
```bash
make test
# o
go test ./...
```

## Validación de Endpoints

Todos los endpoints originales siguen funcionando:

```bash
# Registro
curl -X POST http://localhost:8080/signUp \
  -H "Content-Type: application/json" \
  -d '{"email": "test@test.com", "password": "password123"}'

# Login
curl -X POST http://localhost:8080/signIn \
  -H "Content-Type: application/json" \
  -d '{"email": "test@test.com", "password": "password123"}'

# Crear Tweet (requiere token)
curl -X POST http://localhost:8080/tweets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"message": "Hello world!"}'
```

## Archivos Legacy

Los archivos originales se mantienen para referencia:
- `main.go` (original)
- `handlers/`
- `routers/`
- `bd/`
- `models/`
- `middlew/`
- `jwt/`

**⚠️ Importante**: Para usar la nueva arquitectura, ejecutar `cmd/api/main.go` en lugar de `main.go`.

## Próximos Pasos Recomendados

1. **Eliminar archivos legacy** después de validar que todo funciona
2. **Implementar tests completos** usando los mocks proporcionados
3. **Configurar CI/CD** con los nuevos comandos
4. **Documentar API** con OpenAPI/Swagger
5. **Agregar logging estructurado**
6. **Implementar métricas y monitoring**

## Rollback (Si es Necesario)

Si necesitas volver a la versión anterior temporalmente:

```bash
# Usar el main.go original
go run main.go

# Los archivos originales siguen intactos
```

## Soporte

Para dudas sobre la nueva arquitectura:
1. Consultar [ARCHITECTURE.md](./ARCHITECTURE.md)
2. Revisar los tests de ejemplo en `internal/core/services/tests/`
3. Verificar la documentación en README.md

---

**🎉 ¡Felicitaciones!** Has migrado exitosamente a una arquitectura hexagonal moderna que facilitará el mantenimiento y escalabilidad de tu aplicación.
