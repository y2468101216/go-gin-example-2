package main

import (
	"fmt"
	"gogin/example/pkg/setting"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli"

    _ "github.com/golang-migrate/migrate/v4/source/file"
    _ "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/go-sql-driver/mysql"
)

type (
    // Config information.
    Config struct {
        command string
    }
)

var config Config

func main() {
    app := cli.NewApp()
    app.Name = "Migreate"
    app.Usage = "migrate table"
    app.Action = run
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "command,c",
            Usage: "command for migrate, only up and down",
        },
    }

    app.Run(os.Args)
}

func run(c *cli.Context) {
    config = Config{
        command: c.String("command"),
    }

    exec()
}

func exec() {
    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal("Fail to get section database: ", err)
    }

    dbName := sec.Key("NAME").String()
    user := sec.Key("USER").String()
    password := sec.Key("PASSWORD").String()
    host := sec.Key("HOST").String()

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	user, 
	password, 
	host, 
	dbName)
	
    m, err := migrate.New(
        "file://migrations",
        dsn)
    if err != nil {
        log.Fatal("Fail to connection database: ", err)
    }
    switch config.command {
        case "up" :
            log.Println("run migrate up")
            err = m.Up()
        case "down" :
            log.Println("run migrate down")
            err = m.Down()
        default :
            log.Fatal("No equal command")
    }
    

    if err != nil {
        log.Fatal("Migrate fail database: ", err)
        return
    }

    log.Println("migrate success")
}