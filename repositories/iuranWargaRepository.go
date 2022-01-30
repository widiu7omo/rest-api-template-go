package repositories

import (
	"boilerplate/models"
	"fmt"
	"strconv"
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
func IuranWargaCreate(iuranWarga models.IuranWarga) (models.WargaWithIuran, error) {
	//err := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "iuran_id"}, {Name: "warga_id"}},
	//	UpdateAll: true,
	//}).Create(&iuranWarga).Error
	var count int64
	_ = db.Model(&iuranWarga).Where("iuran_id = ? and warga_id = ?", iuranWarga.IuranID, iuranWarga.WargaID).Count(&count)
	if count == 0 {
		db.Create(&iuranWarga) // create new record from newUser
	} else {
		fmt.Println(iuranWarga)
		db.Where("iuran_id = ? and warga_id = ?", iuranWarga.IuranID, iuranWarga.WargaID).Updates(&iuranWarga)
	}
	wargaWithIuran, err := WargaWithIuranWargaGetById(strconv.FormatInt(iuranWarga.IuranID, 10), strconv.FormatInt(iuranWarga.WargaID, 10))
	//err := db.Preload("Warga").Preload("Iuran").Create(&iuranWarga).Error
	return wargaWithIuran, err
}
func IuranWargaUpdate(iuranWarga models.IuranWarga) (models.IuranWarga, error) {
	err := db.Preload("Warga").Preload("Iuran").Save(&iuranWarga).Error
	return iuranWarga, err
}
func IuranWargaDelete(iuranWarga models.IuranWarga) error {
	return db.Preload("Warga").Preload("Iuran").Delete(&iuranWarga).Error
}
