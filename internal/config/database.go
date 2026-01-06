package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := "host=localhost user=postgres password=1234 port=5432 dbname=products sslmode=disable"

	db,err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err!=nil{
		panic("DataBase is Disconnect")
	}

	DB=db
}