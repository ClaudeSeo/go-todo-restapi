# GO Todo REST API
A Simple todo app written in Go Language.

This project use *go-martini*, *gorm*, *envconfig*

## Environment Setup
```shell
$ export DB_DIALECT="MySQL"
$ export DB_HOST="127.0.0.1"
$ export DB_PORT="3306"
$ export DB_USERNAME="username"
$ export DB_PASSWORD="password"
$ export DB_NAME="dbname"
$ export WEB_PORT="8080"
$ export WEB_HOST="127.0.0.1"
```

## Installation & Run
```shell
$ go get github.com/claudeseo/go-todo-restapi

$ go build
$ ./go-todo-restapi
[martini] listening on 127.0.0.1:8080 (development)

```


