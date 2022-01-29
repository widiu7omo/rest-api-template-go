package repositories

import (
	"boilerplate/helpers"
	"boilerplate/models"
	"fmt"
)

func WargaGet() ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Find(&wargas).Error
	return wargas, err
}

func WargaWithIuranWarga() ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Preload("IuranWargas.Iuran").Find(&wargas).Error
	return wargas, err
}

func WargaGetPagination(page string, pageSize string) ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Scopes(helpers.Paginate(page, pageSize)).Find(&wargas).Error
	return wargas, err
}
func WargaGetById(id string) (models.Warga, error) {
	var warga models.Warga
	err := db.Where("id = ?", id).First(&warga).Error
	return warga, err
}
func WargaCreate(user models.Warga) {
	mu.Lock()
	result := db.Create(&user)
	fmt.Println(result.Error)
	mu.Unlock()
}
