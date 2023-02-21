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

[redis]
Host = 127.0.0.1:6379
Password = secret_redis
DB = 0
```

### Run

1. Clone this project
2. go run migrate.go -c up
3. go run main.go

## Features

- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Cli command
- Redis