package repositories

import (
	"boilerplate/models"
)

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
func JadwalCreate(jadwal models.Jadwal) error {
	result := db.Create(&jadwal)
	return result.Error
}
func JadwalUpdate(jadwal models.Jadwal) error {
	return db.Save(&jadwal).Error
}
func JadwalDelete(jadwal models.Jadwal) error {
	return db.Delete(&jadwal).Error

}
