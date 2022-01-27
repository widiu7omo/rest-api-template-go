package repositories

import (
	"boilerplate/models"
	"fmt"
)

func IuranGet() ([]models.Iuran, error) {
	var iurans []models.Iuran
	err := db.Model(&models.Iuran{}).Find(&iurans).Error
	return iurans, err
}

func IuranWithIuranWargaGet() ([]models.Iuran, error) {
	var iurans []models.Iuran
	err := db.Model(&models.Iuran{}).Preload("IuranWargas.Warga").Find(&iurans).Error
	return iurans, err
}

func IuranGetPagination() {
	db.Find(&models.Iuran{})
}
func IuranGetById(id string) (models.Iuran, error) {
	var iuran models.Iuran
	err := db.Where("id = ?", id).First(&iuran).Error
	return iuran, err
}
func IuranCreate(iuran models.Iuran) {
	mu.Lock()
	result := db.Create(&iuran)
	fmt.Println(result.Error)
	mu.Unlock()
}
