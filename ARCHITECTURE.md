# Arquitectura Hexagonal - Crispy Potato

## Descripción

Este proyecto ha sido refactorizado para implementar la **Arquitectura Hexagonal** (también conocida como Ports and Adapters), siguiendo los principios **SOLID**, **KISS** y **YAGNI**.

## Estructura del Proyecto

```
crispy-potato/
├── cmd/api/                     # Punto de entrada de la aplicación
│   └── main.go                  # Configuración e inicialización
├── internal/                    # Código interno de la aplicación
│   ├── core/                    # Núcleo de la aplicación (Domain + Application)
│   │   ├── domain/              # Entidades de dominio y reglas de negocio
│   │   │   ├── user.go          # Entidad Usuario
│   │   │   ├── tweet.go         # Entidad Tweet
│   │   │   ├── relation.go      # Entidad Relación
│   │   │   └── dto.go           # Objetos de transferencia de datos
│   │   ├── ports/               # Interfaces (Contratos)
│   │   │   ├── repositories.go  # Puertos secundarios (salida)
│   │   │   └── services.go      # Puertos primarios (entrada)
│   │   └── services/            # Casos de uso / Servicios de aplicación
│   │       ├── user_service.go      # Lógica de negocio para usuarios
│   │       ├── tweet_service.go     # Lógica de negocio para tweets
│   │       └── relation_service.go  # Lógica de negocio para relaciones
│   └── adapters/                # Adaptadores externos
│       ├── primary/             # Adaptadores primarios (driving)
│       │   └── http/            # Controladores HTTP
│       │       ├── user_handler.go     # Handler para usuarios
│       │       ├── tweet_handler.go    # Handler para tweets
│       │       ├── relation_handler.go # Handler para relaciones
│       │       └── router.go           # Configuración de rutas
│       └── secondary/           # Adaptadores secundarios (driven)
│           ├── database/mongodb/       # Implementación MongoDB
│           │   ├── connection.go       # Conexión a base de datos
│           │   ├── user_repository.go  # Repositorio de usuarios
│           │   ├── tweet_repository.go # Repositorio de tweets
│           │   └── relation_repository.go # Repositorio de relaciones
│           ├── auth/                   # Autenticación
│           │   ├── password_hasher.go  # Hash de contraseñas
│           │   └── jwt_generator.go    # Generación de JWT
│           └── storage/                # Almacenamiento de archivos
│               └── local_file_storage.go # Almacenamiento local
├── pkg/                         # Paquetes compartidos
│   └── middleware/              # Middlewares HTTP
│       ├── auth.go              # Middleware de autenticación
│       └── database.go          # Middleware de base de datos
├── config/                      # Configuración
│   └── config.go                # Gestión de configuración
└── [archivos legacy]            # Archivos del sistema anterior
```

## Principios Aplicados

### 1. Arquitectura Hexagonal

La aplicación está organizada en capas concéntricas:

- **Centro (Core)**: Contiene la lógica de negocio pura
  - **Domain**: Entidades y reglas de negocio
  - **Ports**: Interfaces que definen contratos
  - **Services**: Casos de uso de la aplicación

- **Adaptadores**: Implementaciones específicas de la infraestructura
  - **Primary**: Adaptadores que invocan la aplicación (HTTP handlers)
  - **Secondary**: Adaptadores invocados por la aplicación (DB, auth, storage)

### 2. Principios SOLID

#### Single Responsibility Principle (SRP)
- Cada entidad tiene una única responsabilidad
- Los servicios están separados por dominio (User, Tweet, Relation)
- Los handlers se encargan únicamente de HTTP
- Los repositorios se encargan únicamente de persistencia

#### Open/Closed Principle (OCP)
- Las interfaces permiten extender funcionalidad sin modificar código existente
- Nuevos adaptadores pueden agregarse implementando las interfaces

#### Liskov Substitution Principle (LSP)
- Cualquier implementación de los puertos puede reemplazar a otra
- Los mocks para testing pueden implementar las mismas interfaces

#### Interface Segregation Principle (ISP)
- Interfaces específicas por dominio (UserRepository, TweetRepository, etc.)
- No se fuerza a implementar métodos innecesarios

#### Dependency Inversion Principle (DIP)
- El core depende de abstracciones (interfaces), no de implementaciones
- Los adaptadores implementan las interfaces definidas en el core

### 3. KISS (Keep It Simple, Stupid)
- Estructura clara y fácil de entender
- Separación clara de responsabilidades
- Código simple y directo

### 4. YAGNI (You Aren't Gonna Need It)
- No se han agregado funcionalidades que no estaban en el sistema original
- Implementación mínima viable de la arquitectura
- Sin over-engineering

## Beneficios de la Refactorización

### 1. Testabilidad
- Fácil testing unitario con mocks
- Dependencias inyectadas
- Lógica de negocio aislada

### 2. Mantenibilidad
- Código organizado y estructurado
- Responsabilidades claras
- Fácil localización de bugs

### 3. Escalabilidad
- Fácil agregar nuevas funcionalidades
- Posibilidad de cambiar implementaciones sin afectar el core
- Estructura preparada para microservicios

### 4. Flexibilidad
- Cambio de base de datos sin afectar lógica de negocio
- Múltiples interfaces (HTTP, gRPC, CLI)
- Diferentes proveedores de almacenamiento

## Cómo Ejecutar

```bash
# Configurar variables de entorno (opcional)
export MONGODB_URI="mongodb://localhost:27017/twittor"
export JWT_SECRET="your-secret-key"
export PORT="8080"
export STORAGE_PATH="."

# Ejecutar la aplicación
go run cmd/api/main.go
```

## Variables de Entorno

| Variable | Descripción | Default |
|----------|-------------|---------|
| `MONGODB_URI` | URI de conexión a MongoDB | `mongodb://localhost:27017/twittor` |
| `JWT_SECRET` | Clave secreta para JWT | `MastersOfDevelopment_facebookGroup` |
| `PORT` | Puerto del servidor | `8080` |
| `STORAGE_PATH` | Ruta base para almacenamiento | `.` |

## Testing

La nueva arquitectura facilita enormemente el testing:

```go
// Ejemplo de test unitario para UserService
func TestUserService_Register(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{}
    mockHasher := &MockPasswordHasher{}
    mockTokenGen := &MockTokenGenerator{}
    mockStorage := &MockFileStorage{}
    
    service := services.NewUserService(mockRepo, mockHasher, mockTokenGen, mockStorage)
    
    // Act
    result, err := service.Register(context.Background(), "test@test.com", "password123")
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

## Migración desde el Sistema Anterior

Los archivos del sistema anterior se mantienen para referencia, pero la nueva aplicación debe usar:

- `cmd/api/main.go` en lugar de `main.go`
- Los nuevos handlers en lugar de `routers/`
- Los nuevos repositorios en lugar de `bd/`
- Los nuevos modelos en lugar de `models/`

## Próximos Pasos

1. Implementar tests unitarios completos
2. Agregar documentación OpenAPI/Swagger
3. Implementar logging estructurado
4. Agregar métricas y monitoring
5. Considerar implementación de eventos de dominio
6. Evaluar migración a microservicios
