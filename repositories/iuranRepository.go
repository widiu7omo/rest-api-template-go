package repositories

import (
	"boilerplate/models"
)

func IuranWithTotal() ([]models.IuranWithTotal, error) {
	var iurans []models.IuranWithTotal
	err := db.Model(&models.IuranWithTotal{}).Table("iurans i").Select("i.*,SUM(iw.amount) as current_total").Joins("left join iuran_wargas iw on iw.iuran_id = i.id").Group("i.id").Find(&iurans).Error
	return iurans, err
}
func IuranGet() ([]models.Iuran, error) {
	var iurans []models.Iuran
	err := db.Preload("IuranWargas.Warga").Find(&iurans).Error
	return iurans, err
}
func IuranGetById(id string) (models.Iuran, error) {
	var iuran models.Iuran
	err := db.Where("id = ?", id).First(&iuran).Error
	return iuran, err
}
func IuranCreate(iuran models.Iuran) (models.Iuran, error) {
	err := db.Create(&iuran).Error
	return iuran, err
}
func IuranUpdate(iuran models.Iuran) (models.Iuran, error) {
	err := db.Save(&iuran).Error
	return iuran, err
}
func IuranDelete(iuran models.Iuran) error {
	return db.Delete(&iuran).Error

}
