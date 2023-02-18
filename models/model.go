package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gogin/example/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var sqlDB *sql.DB

type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn time.Time `json:"created_on"`
    ModifiedOn time.Time `json:"modified_on"`
}


func init() {
    var (
        err error
        dbName, user, password, host, tablePrefix string
    )

    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    dbName = sec.Key("NAME").String()
    user = sec.Key("USER").String()
    password = sec.Key("PASSWORD").String()
    host = sec.Key("HOST").String()
    tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
	user, 
	password, 
	host, 
	dbName)

    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
		  TablePrefix: tablePrefix,   // table name prefix, table for `User` would be `t_users`
		  SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	  })

    if err != nil {
        log.Println(err)
    }

	sqlDB, err := db.DB()

	if err != nil {
        log.Println(err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
}

func CloseDB() {
    defer sqlDB.Close()
}