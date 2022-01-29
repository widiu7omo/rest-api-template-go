package helpers

import (
	"gorm.io/gorm"
	"strconv"
)

func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(pageSize)
		switch {
		case pageSize > 50:
			pageSize = 50
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 5
		case pageSize <= 5:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
