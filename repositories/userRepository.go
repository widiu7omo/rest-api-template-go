package repositories

import (
	"boilerplate/models"
	"fmt"
)

func UserGet() ([]models.User, error) {
	var users []models.User
	err := db.Select([]string{"name", "phone", "email"}).Find(&users).Error
	return users, err
}

func UserGetPagination() {
	db.Find(&models.User{})
}
func UserGetById(id string) (models.User, error) {
	var user models.User
	err := db.Select("name", "phone").Where("id = ?", id).First(&user).Error
	return user, err
}
func UserCreate(user models.User) {
	mu.Lock()
	result := db.Create(&user)
	fmt.Println(result.Error)
	mu.Unlock()
}
