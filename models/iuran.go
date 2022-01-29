package models

type Iuran struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	MinAmount   int64        `json:"min_amount"`
	IuranWargas []IuranWarga `json:"iuran_wargas" gorm:"foreignKey:IuranID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type IuranWithTotal struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	MinAmount    int64        `json:"min_amount"`
	CurrentTotal int64        `json:"current_total"`
	IuranWargas  []IuranWarga `json:"iuran_wargas" gorm:"foreignKey:IuranID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
