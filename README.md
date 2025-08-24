# UserData API

<a id="english"></a>
## English

### Table of Contents
1. [About the Project](#about-en)
2. [Prerequisites](#prerequisites-en)
3. [Installation](#installation-en)
4. [Configuration](#configuration-en)
5. [Running the Application](#running-en)
6. [API Endpoints](#api-endpoints-en)
7. [Examples](#examples-en)
8. [Project Structure](#structure-en)

### <a id="about-en"></a>About the Project
RESTful API for user data management built with Golang and PostgreSQL. Provides CRUD operations for user information including login, personal data, contacts, and activity status.

### <a id="prerequisites-en"></a>Prerequisites
- Docker and Docker Compose
- Go 1.25+ (if running locally)
- PostgreSQL (if running without Docker)

### <a id="installation-en"></a>Installation
1. Clone the repository:
```bash
git clone <repository-url>
cd userdata
```

2. Set up environment variables:
```bash
echo "DB_PASSWORD=your_password" > .env 
```

### <a id="configuration-en"></a>Configuration

#### Environment variables:

- DB_HOST - Database host (default: localhost)
- DB_PORT - Database port (default: 5432)
- DB_USER - Database user (default: postgres)
- DB_PASSWORD - Database password
- DB_NAME - Database name (default: userdata)
- DB_SSLMODE - SSL mode (default: disable)

### <a id="running-en"></a>Running the Application

#### Using Docker Compose (recommended):

```bash
docker-compose up --build
```

#### Using Makefile:

```bash 
make docker-compose-up
```

#### Manual start:

```bash 
go run ./cmd/userdata
```

The API will be available at: http://localhost:8080/users

### <a id="api-endpoints-en"></a>API Endpoints

|Method |Endpoint|Description |
|-------|--------|-------------|
|GET	|/users |	Get all users |
|GET	|/users/:id |	Get user by ID |
|POST	|/users |	Create new user |
|PUT	|/users/:id |	Update user |
|DELETE |	/users/:id |	Delete user |

### <a id="examples-en"></a>Examples

Create a new user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "login": "masha_12",
    "fcs": "Петрова Мария Сергеевна",
    "sex": "female",
    "age": 22,
    "contacts": [
      "masha.pet@mail.ru",
      "+7 999 123 45 67"
    ],
    "avatar": "/avatars/masha.jpg",
    "status": false
  }'
```

### Get all users:
```bash
curl http://localhost:8080/users
```

### <a id="structure-en"></a>Project Structure

```text
userdata/
├── cmd/
│   └── userdata/
│       └── main.go
├── internal/
│   ├── createU/
│   │   └── create.go
│   ├── deleteU/
│   │   └── delete.go
│   ├── handler/
│   │   └── handler.go
│   ├── readU/
│   │   └── read.go
│   ├── storage/
│   │   ├── migrate.go
│   │   ├── postgres.go
│   │   └── user.go
│   └── updateU/
│       └── update.go
├── .dockerignore
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

