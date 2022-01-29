package database

import (
	"boilerplate/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbUrl = "u6945793_erte5:JbAADXqobwWQ@tcp(127.0.0.1:3306)/u6945793_erte5?charset=utf8mb4"

//const dbUrl = "root:@tcp(127.0.0.1:3306)/posrt5?charset=utf8mb4"

var (
	mysqlDb  *gorm.DB
	mysqlErr error
)

// Connect with database
func Connect() {
	mysqlDb, mysqlErr = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if mysqlErr != nil {
		fmt.Println("Failed to reach Database")
		return
	}
	fmt.Println("Connected with Database")
	if Migration() {
		fmt.Println("Migration Successfully")
	} else {
		fmt.Println("Migration failed")
	}
}

func Migration() bool {
	err := mysqlDb.AutoMigrate(&models.User{}, &models.Warga{}, &models.Iuran{}, &models.IuranWarga{})
	fmt.Println(err)
	return !(err != nil)
}
func DBInstance() *gorm.DB {
	return mysqlDb
}
