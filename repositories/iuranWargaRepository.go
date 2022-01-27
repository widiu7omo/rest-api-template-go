package repositories

import (
	"boilerplate/models"
)

func IuranWargaGet() ([]models.IuranWarga, error) {
	var iuranWargas []models.IuranWarga
	err := db.Find(&iuranWargas).Error
	return iuranWargas, err
}
func IuranWargaGetById(id string) (models.IuranWarga, error) {
	var iuranWarga models.IuranWarga
	err := db.Where("id = ?", id).First(&iuranWarga).Error
	return iuranWarga, err
}
func IuranWargaCreate(iuranWarga models.IuranWarga) error {
	result := db.Create(&iuranWarga)
	return result.Error
}
func IuranWargaUpdate(iuranWarga models.IuranWarga) error {
	return db.Save(&iuranWarga).Error
}
func IuranWargaDelete(iuranWarga models.IuranWarga) error {
	return db.Delete(&iuranWarga).Error

}
