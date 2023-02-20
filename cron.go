package main

import (
    "time"
    "log"

    "github.com/robfig/cron/v3"
	"gogin/example/models"
)

func main() {
    log.Println("Starting...")

    c := cron.New()
    c.AddFunc("* * * * *", func() {
        log.Println("Run models.CleanAllArticle...")
        models.CleanAllArticle()
    })

    c.Start()

    //handle process
    time.Sleep(time.Minute * 5)
}