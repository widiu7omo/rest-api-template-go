package models

type Warga struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Phone       string       `json:"phone"`
	IuranWargas []IuranWarga `json:"iuran_wargas" gorm:"foreignKey:WargaID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type WargaWithIuran struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Date   string `json:"date"`
	Amount int64  `json:"amount"`
}
