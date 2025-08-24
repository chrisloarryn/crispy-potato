# Crispy Potato - Twitter-like API

Una API REST similar a Twitter construida con Go, implementando Arquitectura Hexagonal y siguiendo los principios SOLID, KISS y YAGNI.

## 🏗️ Arquitectura

Este proyecto implementa **Arquitectura Hexagonal** (Ports and Adapters) para lograr:

- ✅ Separación clara de responsabilidades
- ✅ Alta testabilidad
- ✅ Fácil mantenimiento
- ✅ Flexibilidad para cambios futuros

Ver [ARCHITECTURE.md](./ARCHITECTURE.md) para detalles completos.

## 🚀 Características

- **Gestión de Usuarios**: Registro, login, perfiles
- **Sistema de Tweets**: Crear, leer, eliminar tweets
- **Relaciones**: Seguir/dejar de seguir usuarios
- **Autenticación JWT**: Seguridad basada en tokens
- **Upload de Archivos**: Avatares y banners
- **Feed de Seguidos**: Ver tweets de usuarios seguidos

## 🛠️ Tecnologías

- **Go 1.24+**
- **MongoDB** - Base de datos NoSQL
- **JWT** - Autenticación
- **Gorilla Mux** - Router HTTP
- **bcrypt** - Hash de contraseñas
- **CORS** - Cross-Origin Resource Sharing

## 📦 Instalación

### Prerrequisitos

- Go 1.24 o superior
- MongoDB (local o remoto)

### Configuración

1. **Clonar el repositorio**
```bash
git clone https://github.com/chrisloarryn/crispy-potato.git
cd crispy-potato
```

2. **Instalar dependencias**
```bash
go mod tidy
```

3. **Configurar variables de entorno** (opcional)
```bash
# Crear archivo .env o exportar variables
export MONGODB_URI="mongodb://localhost:27017/twittor"
export JWT_SECRET="your-secret-key"
export PORT="8080"
export STORAGE_PATH="."
```

4. **Ejecutar la aplicación**
```bash
go run cmd/api/main.go
```

El servidor estará disponible en `http://localhost:8080`

## 📡 API Endpoints

### Autenticación

```http
POST /signUp
POST /signIn
```

### Usuarios

```http
GET    /{id}/{idBox}/me      # Obtener perfil
PUT    /me                   # Actualizar perfil
GET    /usersFollow          # Listar usuarios
GET    /avatars              # Obtener avatar
POST   /avatars              # Subir avatar
GET    /banners              # Obtener banner
POST   /banners              # Subir banner
```

### Tweets

```http
POST   /tweets               # Crear tweet
GET    /tweets               # Obtener tweets
DELETE /tweets               # Eliminar tweet
GET    /tweetsFollowers      # Feed de seguidos
```

### Relaciones

```http
POST   /relations            # Seguir usuario
DELETE /relations            # Dejar de seguir
GET    /relations            # Estado de relación
```

## 🧪 Testing

```bash
# Ejecutar todos los tests
go test ./...

# Test con coverage
go test -cover ./...

# Test verbose
go test -v ./...
```

## 📝 Ejemplos de Uso

### Registro de Usuario

```bash
curl -X POST http://localhost:8080/signUp \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/signIn \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### Crear Tweet

```bash
curl -X POST http://localhost:8080/tweets \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "message": "Mi primer tweet!"
  }'
```

## 🐳 Docker (Próximamente)

```bash
# Construcción
docker build -t crispy-potato .

# Ejecución
docker run -p 8080:8080 crispy-potato
```

## 🤝 Contribución

1. Fork el proyecto
2. Crear rama feature (`git checkout -b feature/AmazingFeature`)
3. Commit cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.

## 📞 Contacto

Cristobal Contreras - [@chrisloarryn](https://github.com/chrisloarryn)

Proyecto: [https://github.com/chrisloarryn/crispy-potato](https://github.com/chrisloarryn/crispy-potato)

## 🙏 Agradecimientos

- Inspirado en las mejores prácticas de arquitectura de software
- Basado en los principios de Clean Architecture
- Implementación de patrones de diseño modernos