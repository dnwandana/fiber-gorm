# Fiber Gorm

Simple CRUD Rest Api, created using [GoFiber](https://gofiber.io/ "An Express-inspired web framework written in Go") and [GORM](https://gorm.io/ "The fantastic ORM library for Golang")

## Download Depedencies

```bash
go mod download
```

## Running the Application

```bash
go run main.go
```

## Compile the Application

```bash
go build main.go
```

## Dockerize the Application

```bash
# build
docker build -t my-golang-app:1.0 .

# run
docker run --name my-golang-app -p 5000:5000 -e DB_HOST="host" -e DB_PORT="5432" -e DB_NAME="fiber_gorm" -e DB_USER="superuser" -e DB_PASS="secret" -e DB_TZ="Asia/Jakarta" -d my-golang-app:1.0
```

## Reference

- [GORM Documentation](https://gorm.io/docs/ "GORM Documentation Website")
- [GoFiber Documentation](https://docs.gofiber.io/ "GoFiber Documentation Website")
