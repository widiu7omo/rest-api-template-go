package repositories

import (
	"boilerplate/helpers"
	"boilerplate/models"
)

func WargaGetPagination(page string, pageSize string) ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Scopes(helpers.Paginate(page, pageSize)).Find(&wargas).Error
	return wargas, err
}
func WargaGet() ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Find(&wargas).Error
	return wargas, err
}
func WargaGetById(id string) (models.Warga, error) {
	var warga models.Warga
	err := db.Where("id = ?", id).First(&warga).Error
	return warga, err
}
func WargaCreate(warga models.Warga) (models.Warga, error) {
	err := db.Create(&warga).Error
	return warga, err
}
func WargaUpdate(warga models.Warga) (models.Warga, error) {
	err := db.Save(&warga).Error
	return warga, err
}
func WargaDelete(warga models.Warga) error {
	return db.Delete(&warga).Error

}
