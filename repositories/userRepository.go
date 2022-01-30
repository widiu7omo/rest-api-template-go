package repositories

import (
	"boilerplate/helpers"
	"boilerplate/models"
)

func UserGetPagination(page string, pageSize string) ([]models.User, error) {
	var users []models.User
	err := db.Scopes(helpers.Paginate(page, pageSize)).Find(&users).Error
	return users, err
}
func UserGet() ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}
func UserGetById(id string) (models.User, error) {
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}
func UserCreate(user models.User) (models.User, error) {
	err := db.Create(&user).Error
	return user, err
}
func UserUpdate(user models.User) (models.User, error) {
	err := db.Save(&user).Error
	return user, err
}
func UserDelete(user models.User) error {
	return db.Delete(&user).Error

}
