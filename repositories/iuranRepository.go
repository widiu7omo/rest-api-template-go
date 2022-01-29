package repositories

import (
	"boilerplate/models"
	"fmt"
)

func IuranWithTotal() ([]models.IuranWithTotal, error) {
	var iurans []models.IuranWithTotal
	err := db.Model(&models.IuranWithTotal{}).Table("iurans i").Select("i.*,SUM(iw.amount) as current_total").Joins("left join iuran_wargas iw on iw.iuran_id = i.id").Group("i.id").Find(&iurans).Error
	return iurans, err
}
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
