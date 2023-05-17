package models

import (
	setting "back/pkg"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Fail to open 'database' :%v in 'models/models.go' ", err)
	}
	InitModels()
}

func InitModels() {
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&Faculty{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&System{})
	db.AutoMigrate(&FC{})
}
