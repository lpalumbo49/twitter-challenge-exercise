<div align="center">
<h3 align="center">twitter-challenge-exercise</h3>

  <p align="center">
    El objetivo de este proyecto es crear una versión simplificada de una plataforma de microblogging similar a twitter que
permita a los usuarios publicar, seguir y ver el timeline de tweets.
  </p>
</div>

### Construido con

* [Go](https://go.dev/) Lenguaje de programación
* [Gin-Gonic](https://github.com/gin-gonic/gin) Framework HTTP para la creación de la API REST
* [JWT](https://jwt.io/) Tokens de autenticación de usuarios 
* [MySQL](https://www.mysql.com/) Base de datos relacional

## 🚀  Como correr el proyecto

Para ejecutar localmente esta aplicación, hay que seguir los siguientes pasos:

### Prerequisitos

* [Docker](https://www.docker.com/) Necesario para ejecutar el container que contiene la aplicación, y una base de datos in memory

### Instalación

1. Clonar el repositorio
   ```sh
   git clone https://github.com/lpalumbo49/twitter-challenge-exercise.git
   ```
2. Ejecutar el contenedor mediante Docker
   ```sh
   docker compose up -d --build
   ```
3. Las credenciales de conexión a la base de datos y para la autenticación de usuarios, se encuentran en `./docker-compose.yml`
   ```yml
   # These should be handled as secrets, outside the scope of this application
   environment:
      - DATABASE_HOST=localhost
      - DATABASE_PORT=3306
      - DATABASE_USER=root
      - DATABASE_PASSWORD=password
      - DATABASE_NAME=twitter_db
      - JWT_TOKEN_SECRET=this_is_a_secret
   ```
4. Ejecutar en la base de datos (solo por única vez) los scripts presentes en `./scripts/db_creation_queries.sql`, para crear las tablas de la aplicación

5. La aplicación se ejecutará en el puerto 8080: `http://localhost:8080`


## 📡 Uso de la API

### Endpoints públicos

### Crear usuario
Crea un nuevo usuario.
```
POST /api/v1/user
```
#### Request
  ``` json
  {
    "name": "Nombre",
    "surname": "Apellido",
    "email": "email@test.com",
    "password": "super-safe-password",
    "username": "nombre_de_usuario"
  }
  ```
#### Response
  ``` json
  {
    "id": 49,
    "name": "Nombre",
    "surname": "Apellido",
    "email": "email@test.com",
    "username": "nombre_de_usuario",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-09T00:00:00Z"
  }
  ```

### Login de usuario
Realiza el login de un usuario existente. Devuelve el token JWT que luego deberá ser utilizado para poder acceder a los métodos privados de la API.
```
POST /api/v1/login
```
#### Request
  ``` json
  {
    "email": "email@test.com",
    "password": "super-safe-password"
  }
  ```
#### Response
  ``` json
  {
    "token": "dgsfkjhgdfsjkh534kj543kjhfdskjh5234bnfdskj5432jk"
  }
  ```
### Endpoints privados
A partir de aquí, los siguientes métodos requieren todos ser autenticados con el header `Authorization: Bearer {token}` para poder ser utilizados, con el token generado en el método `/login`.

### Get de usuario por ID
Obtiene un usuario por su ID.
```
GET /api/v1/user/:id
```
#### Response
  ``` json
  {
    "id": 49,
    "name": "Nombre",
    "surname": "Apellido",
    "email": "email@test.com",
    "username": "nombre_de_usuario",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-09T00:00:00Z"
  }
  ```

### Modificación de usuario
Modifica los campos de un usuario por ID. Por simplicidad de diseño, deben enviarse todos los campos del body.
```
PUT /api/v1/user/:id
```
#### Request
  ``` json
  {
    "id": 49,
    "username": "nombre_de_usuario",
    "name": "Nombre",
    "surname": "Apellido"
  }
  ```

#### Response
  ``` json
  {
    "id": 49,
    "name": "Nombre",
    "surname": "Apellido",
    "email": "email@test.com",
    "username": "nombre_de_usuario",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-09T00:00:00Z"
  }
  ```

### Search de usuarios
Realiza el search de todos los usuarios. Este método podría ampliarse para implementar funcionalidades de full-text search.
```
GET /api/v1/users
```
#### Response
  ``` json
  {
    "users": [
      {
        "id": 49,
        "name": "Nombre",
        "surname": "Apellido",
        "email": "email@test.com",
        "username": "nombre_de_usuario",
        "created_at": "2025-07-09T00:00:00Z",
        "updated_at": "2025-07-09T00:00:00Z"
      },
      {
        "id": 50,
        "name": "Otro Nombre",
        "surname": "Otro Apellido",
        "email": "otro_email@test.com",
        "username": "otro_nombre_de_usuario",
        "created_at": "2025-06-10T00:00:00Z",
        "updated_at": "2025-06-10T00:00:00Z"
      }
    ]
  }
  ```

### Creación de tweet
Crea un nuevo tweet.
```
POST /api/v1/tweet
```
#### Request
  ``` json
  {
    "user_id": 49,
    "text": "Este es el texto de un tweet. Tiene un máximo de 280 caracteres."
  }
  ```
#### Response
  ``` json
  {
    "id": 126,
    "user_id": 49,
    "text": "Este es el texto de un tweet. Tiene un máximo de 280 caracteres",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-09T00:00:00Z"    
  }
  ```

### Get de tweet por ID
Obtiene un tweet por su ID.
```
GET /api/v1/tweet/:id
```
#### Response
  ``` json
  {
    "id": 126,
    "user_id": 49,
    "text": "Este es el texto de un tweet. Tiene un máximo de 280 caracteres",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-09T00:00:00Z"       
  }
  ```

### Modificación de tweet
Modifica el texto de un tweet. Solo puede realizar esta operación el usuario que haya creado al tweet.
```
PUT /api/v1/tweet/:id
```
#### Request
  ``` json
  {
    "id": 126,
    "user_id": 49,
    "text": "Este es un texto modificado"
  }
  ```
#### Response
  ``` json
  {
    "id": 126,
    "user_id": 49,
    "text": "Este es un texto modificado",
    "created_at": "2025-07-09T00:00:00Z",
    "updated_at": "2025-07-10T00:00:00Z"         
  }
  ```

### Obtener el timeline de un usuario
Obtiene el timeline del usuario, que consiste en los tweets de los usuarios que él sigue, ordenados decrecientemente por fecha de creación.
El ID de usuario corresponderá al que esté autenticado.
```
GET /api/v1/timeline
```
#### Response
  ``` json
  {
    "timeline": [
      {
        "id": 126,
        "user_id": 49,
        "text": "Este es el texto de un tweet. Tiene un máximo de 280 caracteres",
        "created_at": "2025-07-09T00:00:00Z",
        "updated_at": "2025-07-09T00:00:00Z",
        "user": {
          "id": 49,
          "name": "Nombre",
          "surname": "Apellido",
          "email": "email@test.com",
          "username": "nombre_de_usuario",
          "created_at": "2025-07-09T00:00:00Z",
          "updated_at": "2025-07-09T00:00:00Z"        
        }
      },
      {
        "id": 96,
        "user_id": 50,
        "text": "Este es un un tweet de otro usuario que me interesa",
        "created_at": "2025-07-01T00:00:00Z",
        "updated_at": "2025-07-01T00:00:00Z",
        "user": {
        "id": 50,
          "name": "Otro Nombre",
          "surname": "Otro Apellido",
          "email": "otro_email@test.com",
          "username": "otro_nombre_de_usuario",
          "created_at": "2025-06-10T00:00:00Z",
          "updated_at": "2025-06-10T00:00:00Z"        
        }    
      }
    ]
  }
  ```

### Crear follower de usuario
Crea un follower de usuario. Es decir, una asociación entre un usuario y el usuario que le interesa seguir.
```
POST /api/v1/follower
```
#### Request
  ``` json
  {
    "user_id": 50,
    "followed_by_user_id": 49
  }
  ```
#### Response
  ``` json
  {
    "user_id": 50,
    "followed_by_user_id": 49.
    "created_at": "2025-07-09T00:00:00Z"    
  }
  ```
## 📁 Estructura del proyecto

   ```sh
.
├── Dockerfile
├── docker-compose.yml
├── go.mod                      # Módulos y dependencias en libraries de la aplicación
├── go.sum
├── cmd/
│   └── api/
│       └── main.go             # Punto de entrada y ejecución de la aplicación
├── docs/                       # Imágenes que se utilizan en esta documentación
├── internal/                   # Código privado de la aplicación
│   ├── adapter/                # Implementaciones de los métodos genéricamente definidos en los ports
│   │   ├── handler/            # Métodos expuestos para comunicación entrante y saliente con clientes 
│   │   │   └── http/           # Comunicación REST
│   │   │       └── dto/        # Métodos de transformación de entidades de comunicación a entidades de negocio
│   │   │       └── middleware/ # Implementación de middleware de autenticación
│   │   └── repository/         # Métodos relacionados a implementaciones de almacenamiento (base de datos)
│   ├── config/                 # Lectura de configuraciones de entorno
│   ├── container.go            # Wiring (inyección) de dependencias
│   └── core/                   # Código relacionado al dominio de la aplicación
│       ├── domain/             # Entidades y modelos de negocio
│       ├── port/               # Definición de interfaces para la comunicación entre el mundo exterior y el core
│       └── service/            # Casos de uso propiamente dichos
├── pkg/                        # Libraries y gestión de errores. Podría implementarse en un repositorio toolkit
├── scripts/                    # Scripts SQL de creación de tablas
└── README.md

   ```

Este proyecto sigue una organización basada en los principios de la _Arquitectura Hexagonal_ (también llamada _Ports y Adapters_), permitiendo una alta separación de responsabilidades, facilidad de testeo, y escalabilidad.
Resumidamente, los tres actores principales en esta arquitectura son:

* Core: El core es la lógica principal de la aplicación. Contiene todas las reglas de negocio y funcionamiento propiamente dicho de la misma.
* Port: Un port es un contrato que define como la aplicación se comunicará con sistemas o servicios externos.
* Adapter: El adapter es una implementación concreta de un port. Posee el detalle técnico necesario para cumplir un contrato en particular.

Esto permite que las `interfaces` de golang puedan ser aprovechadas para inyección de dependencias, que permite el fácil desacople y cambio de módulos de la aplicación (por ejemplo motor de persistencia, punto de entrada de HTTP a Kafka, etc) muy fácilmente. Además, cada componente del core puede testearse de forma aislada, utilizando mocks de sus puertos.

Los `handlers` HTTP analizan las solicitudes HTTP entrantes, validan su formato y autenticación, y responden a los clientes con los códigos de estado definidos en el estándar REST.
Los handlers se comunican con el `core` de la aplicación mediante la implementación de los ports (`services`), que manejan la lógica de negocio.
Los services acceden a la capa de persistencia mediante los `repositories`.

## 🧪 Testing

Este proyecto incluye algunos ejemplos de tests unitarios para los servicios de usuarios y tweets, y tests de integración para probar la implementación del servicio de tweets. 

### Ejecutar todos los tests:
   ```sh
   go test -v ./...
   ```
### Ejecutar tests unitarios:
   ```sh
   go test -v ./internal/core/service
   ```
### Ejecutar tests de integración:
   ```sh
   go test -v ./internal/adapter/handler/http
   ```

## 🏗️ Arquitectura

TODO: explicación de la arquitectura

## 🤝 Contacto

Link del proyecto: [https://github.com/lpalumbo49/twitter-challenge-exercise](https://github.com/lpalumbo49/twitter-challenge-exercise)

