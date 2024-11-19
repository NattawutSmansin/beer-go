# Set up project
```bash
$ go mod init beer
```

```bash base lib
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/joho/godotenv
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/swag
```

```bash database
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get go.mongodb.org/mongo-driver/mongo
```

``` run host & swagger
go run main.go
swag init
```