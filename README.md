# Go Gin Example In 2022

fork from this [project](https://github.com/eddycjy/go-gin-example) and rewrite it to comply With Golang 1.19 in 2022.

## How to Run 

### Required 

- Mysql 

### Ready

Create a **blog database**

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_
```

### Run

1. Clone this project
2. go build migrate.go
3. ./migrate -c up or ./migrate.exe -c up
4. go run main.go

## Features

- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Cli command