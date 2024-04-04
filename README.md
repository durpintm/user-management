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

## Test the API

- Assumption: At least one user should exist in the Users table to be able to generate invitation code

```bash
### User Login
curl -X POST http://localhost:4000/login -H "Content-Type: application/json" -d "{\"email\":\"youremail@gmail.com\", \"password\":\"yourpassword\"}"

### Get Invitation Code
curl -X GET 'http://localhost:4000/invitationCode' -H 'Cookie: Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1cnBpbnRoYXBhQGdtYWlsLmNvbSIsImV4cCI6MTcxNDc3MDk2MCwic3ViIjoxfQ.X69s4PfObKsoibvxudVHQ10btFtrgKFS5A5r012caC0'

### City Weather using GET request
curl  -X POST 'http://localhost:4000/signup' -H 'Content-Type: application/json' -H 'Cookie: Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImR1cnBpbnRoYXBhQGdtYWlsLmNvbSIsImV4cCI6MTcxNDc3MTg5Miwic3ViIjoxfQ.TIyMdcoQ-UZnBTpXESLzFv3MWKHFfcOUkZizdOsULv4' --data-raw '{
    "email": "newemail@gmail.com",
    "password":"newpassword",
    "invitationcode": "invite_code"
}'
```

## ER Diagram

![ER Diagram](https://github.com/durpintm/user-management/blob/main/images/er-diagram.png)
