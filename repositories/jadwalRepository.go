package repositories

import (
	"boilerplate/helpers"
	"boilerplate/models"
)

func JadwalGetPagination(page string, pageSize string) ([]models.Jadwal, error) {
	var jadwals []models.Jadwal
	err := db.Scopes(helpers.Paginate(page, pageSize)).Find(&jadwals).Error
	return jadwals, err
}
func JadwalGet() ([]models.Jadwal, error) {
	var jadwals []models.Jadwal
	err := db.Find(&jadwals).Error
	return jadwals, err
}
func JadwalGetById(id string) (models.Jadwal, error) {
	var jadwal models.Jadwal
	err := db.Where("id = ?", id).First(&jadwal).Error
	return jadwal, err
}
func JadwalCreate(jadwal models.Jadwal) (models.Jadwal, error) {
	err := db.Create(&jadwal).Error
	return jadwal, err
}
func JadwalUpdate(jadwal models.Jadwal) (models.Jadwal, error) {
	err := db.Save(&jadwal).Error
	return jadwal, err
}
func JadwalDelete(jadwal models.Jadwal) error {
	return db.Delete(&jadwal).Error

}
