package repositories

import (
	"boilerplate/models"
)

func IuranWargaGet() ([]models.IuranWarga, error) {
	var iuranWargas []models.IuranWarga
	err := db.Preload("Warga").Preload("Iuran").Find(&iuranWargas).Error
	return iuranWargas, err
}
func IuranWargaGetById(id string) (models.IuranWarga, error) {
	var iuranWarga models.IuranWarga
	err := db.Preload("Warga").Preload("Iuran").Where("id = ?", id).First(&iuranWarga).Error
	return iuranWarga, err
}
func IuranWargaCreate(iuranWarga models.IuranWarga) (models.IuranWarga, error) {
	err := db.Preload("Warga").Preload("Iuran").Create(&iuranWarga).Error
	return iuranWarga, err
}
func IuranWargaUpdate(iuranWarga models.IuranWarga) (models.IuranWarga, error) {
	err := db.Preload("Warga").Preload("Iuran").Save(&iuranWarga).Error
	return iuranWarga, err
}
func IuranWargaDelete(iuranWarga models.IuranWarga) error {
	return db.Preload("Warga").Preload("Iuran").Delete(&iuranWarga).Error
}
