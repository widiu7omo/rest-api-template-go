package repositories

import (
	"boilerplate/helpers"
	"boilerplate/models"
)

func WargaGetPagination(page string, pageSize string) ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Preload("IuranWargas").Scopes(helpers.Paginate(page, pageSize)).Find(&wargas).Error
	return wargas, err
}
func WargaWithIuranWargaGetPagination(page string, pageSize string, iuranId string) ([]models.WargaWithIuran, error) {
	var wargas []models.WargaWithIuran
	err := db.Table("wargas").Select("wargas.*,(select amount from iuran_wargas iw where iw.warga_id = wargas.id and iuran_id = ?) as amount,(select date from iuran_wargas iw where iw.warga_id = wargas.id and iuran_id = ?) as date", iuranId, iuranId).Scopes(helpers.Paginate(page, pageSize)).Find(&wargas).Error
	return wargas, err
}
func WargaWithIuranWargaGetById(iuranId string, wargaId string) (models.WargaWithIuran, error) {
	var warga models.WargaWithIuran
	err := db.Table("wargas").Select("wargas.*,(select amount from iuran_wargas iw where iw.warga_id = wargas.id and iuran_id = ?) as amount,(select date from iuran_wargas iw where iw.warga_id = wargas.id and iuran_id = ?) as date", iuranId, iuranId).Where("id = ?", wargaId).Find(&warga).Error
	return warga, err
}
func WargaGet() ([]models.Warga, error) {
	var wargas []models.Warga
	err := db.Preload("IuranWargas").Find(&wargas).Error
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
