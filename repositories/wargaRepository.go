package repositories

import (
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

func WargaGetPagination() {
	db.Find(&models.Warga{})
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
