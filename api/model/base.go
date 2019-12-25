package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
)

const (
	kDBUsername string = "TODO_APP_DB_USERNAME"
	kDBPassword string = "TODO_APP_DB_PASSWORD"
	kDBHost     string = "TODO_APP_DB_HOST"
	kDBName     string = "TODO_APP_DB_NAME"
)

var db *gorm.DB

func init() {
	if e := godotenv.Load(); e != nil {
		panic(e)
	}

	username := os.Getenv(kDBUsername)
	password := os.Getenv(kDBPassword)
	host := os.Getenv(kDBHost)
	dbName := os.Getenv(kDBName)

	dbUri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Todo{})
}

func GetDB() *gorm.DB {
	return db
}
