//This File used for setup connection to database
//gorm is a dep for coonecting database to golang project
package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {			   //username:password(if exist)@tcp(database URL)/database_name
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/db_project_tranparansi_publik"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Transparasi{})
	database.AutoMigrate(&User{})

	DB = database
}