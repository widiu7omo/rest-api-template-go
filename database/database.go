package database

import (
	"boilerplate/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlDb  *gorm.DB
	mysqlErr error
)

// Connect with database
func Connect(dbUrl string) {
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
