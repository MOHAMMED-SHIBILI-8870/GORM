package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn:="host=localhost user=postgres password=1234 dbname=products port=5432 sslmode=disable"

	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		panic("Database connection is droped")
	}

	DB=db
}