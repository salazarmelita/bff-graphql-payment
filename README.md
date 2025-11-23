# 🚀 GraphQL Payment BFF

Backend for Frontend (BFF) implementando **Clean Architecture** + **Arquitectura Hexagonal** para servicios de pago con GraphQL y gRPC.

## 📋 Características

- ✅ **Clean Architecture** con separación clara de capas
- ✅ **Arquitectura Hexagonal** con puertos e interfaces bien definidos
- ✅ **GraphQL API** con gqlgen v0.17.78+
- ✅ **gRPC Client** para comunicación con microservicios
- ✅ **Dependency Injection** con contenedor personalizado
- ✅ **Health Check** endpoint `/ping`
- ✅ **GraphQL Playground** en ruta raíz `/`
- ✅ **CORS** configurado para desarrollo
- ✅ **Docker** ready con multi-stage build
- ✅ **Graceful Shutdown** con contexto y timeout

## 🏗️ Arquitectura

```
Clean Architecture + Hexagonal Architecture
├── Domain (Core) - Sin dependencias externas
│   ├── model/       # Entidades y Value Objects
│   ├── ports/       # Interfaces de casos de uso
│   ├── service/     # Servicios de dominio
│   └── exception/   # Excepciones de dominio
├── Application - Casos de uso y puertos
│   ├── service/     # Casos de uso (use cases)
│   ├── ports/       # Puertos de salida (repositories)
│   └── exception/   # Excepciones de aplicación
└── Infrastructure - Adaptadores y frameworks
    ├── inbound/     # Adaptadores de entrada (GraphQL)
    └── outbound/    # Adaptadores de salida (gRPC, Cache)
```

## 🚀 Inicio Rápido

### Prerrequisitos

- Go 1.21+
- Git

### Instalación

1. **Setup inicial del proyecto:**
```bash
scripts\setup.bat
```

2. **Ejecutar en modo desarrollo:**
```bash
scripts\run_dev.bat
```

3. **O ejecutar manualmente:**
```bash
go run cmd/server/main.go
```

### URLs Importantes

- **GraphQL Playground**: http://localhost:8080/
- **GraphQL Endpoint**: http://localhost:8080/query
- **Health Check**: http://localhost:8080/ping

## 📊 Funcionalidades Disponibles

### GraphQL Queries

#### 1. Health Check
```graphql
query {
  ping
}
```

#### 2. Obtener Información de Infraestructura de Pago
```graphql
query GetPaymentInfra {
  getPaymentInfraByID(input: { paymentRackId: "rack-001" }) {
    transactionId
    message
    status
    paymentRack {
      id
      description
      address
    }
    installation {
      id
      name
      region
      city
      address
      imageUrl
    }
    bookingTimes {
      id
      name
      unitMeasurement
      amount
    }
  }
}
```

## 🛠️ Desarrollo

### Estructura del Proyecto

```
graphql-payment-bff/
├── cmd/server/              # Entry point
├── config/                  # Configuración e inyección dependencias
├── graph/                   # GraphQL schemas y generados
├── internal/
│   ├── domain/             # CAPA DOMINIO (CORE)
│   ├── application/        # CAPA APLICACIÓN
│   └── infrastructure/     # CAPA INFRAESTRUCTURA
├── proto/payment/          # Archivos .proto
├── scripts/                # Scripts de automatización
├── docker-compose.yml      # Para desarrollo local
├── Dockerfile              # Para contenerización
└── README.md
```

### Scripts Disponibles

- `scripts\setup.bat` - Setup inicial del proyecto
- `scripts\run_dev.bat` - Ejecutar en modo desarrollo
- `scripts\gen_graphql.bat` - Regenerar código GraphQL
- `scripts\gen_proto.bat` - Regenerar código protobuf

### Regenerar Código

**GraphQL:**
```bash
scripts\gen_graphql.bat
```

**Protobuf:**
```bash
scripts\gen_proto.bat
```

## 🐳 Docker

### Desarrollo con Docker Compose

```bash
docker-compose up --build
```

### Build individual

```bash
docker build -t graphql-payment-bff .
docker run -p 8080:8080 graphql-payment-bff
```

## ⚙️ Configuración

### Variables de Entorno

Copia `.env.example` a `.env` y configura:

```env
# Server Configuration
SERVER_PORT=8080

# gRPC Services Configuration
PAYMENT_SERVICE_GRPC_ADDRESS=localhost:50051

# JWT Configuration (for future use)
JWT_SECRET=your-jwt-secret-key
JWKS_URL=https://your-auth-provider.com/.well-known/jwks.json

# Development Configuration
ENVIRONMENT=development
LOG_LEVEL=debug
```

## 🧪 Testing

### Probar la API

1. **Health Check:**
```bash
curl http://localhost:8080/ping
```

2. **GraphQL Query en Playground:**
   - Ir a http://localhost:8080/
   - Ejecutar queries de ejemplo

3. **GraphQL Query con curl:**
```bash
curl -X POST \
  http://localhost:8080/query \
  -H 'Content-Type: application/json' \
  -d '{
    "query": "query { ping }"
  }'
```

## 📋 Próximos Pasos

- [ ] Implementar autenticación JWT
- [ ] Agregar más servicios de pago
- [ ] Implementar WebSocket subscriptions
- [ ] Agregar logging estructurado
- [ ] Implementar métricas y observabilidad
- [ ] Agregar tests unitarios e integración
- [ ] Configurar CI/CD pipeline

## 🤝 Contribución

1. Fork el proyecto
2. Crear branch para feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push al branch (`git push origin feature/nueva-funcionalidad`)
5. Crear Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver archivo [LICENSE](LICENSE) para detalles.

---

**🚀 Happy Coding!** 

Para más información, revisa la documentación en [PROMPT_INICIALIZAR_NUEVO_BFF.md](PROMPT_INICIALIZAR_NUEVO_BFF.md)# bff-graphql-payment
