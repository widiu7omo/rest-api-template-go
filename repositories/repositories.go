package repositories

import (
	"boilerplate/database"
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB
	mu sync.Mutex
)

func Init() {
	db = database.DBInstance()
}
