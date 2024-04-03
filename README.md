# User Management API in Go

## Packages used

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
```

## Pull PostgreSQL Docker Image

```bash
docker pull postgres
```

## Setting up PostgreSQL database Container

```bash
 docker run --name {ContainerName} -e POSTGRES_USER={username} -e POSTGRES_PASSWORD={password} -e POSTGRES_DB={Database} -p 5432:5432 -d postgres
```

## Check that the PostgreSQL container is running

```bash
docker ps
```

## Running application thorugh Compile Daemon

```bash
compiledaemon --command="./user-management"
```
