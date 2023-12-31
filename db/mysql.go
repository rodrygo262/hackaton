package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_NAME = "provisioning_connector"
	DB_HOST = "localhost"
	DB_USER = "root"
	DB_PASS = "rootpassword"
	DB_PORT = "3306"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))

	if err != nil {
		fmt.Print("erro ao conectar com database")
	}

	return db
}
