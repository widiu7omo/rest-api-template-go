package models

type Iuran struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	MinAmount   int64        `json:"min_amount"`
	IuranWargas []IuranWarga `json:"iuran_wargas" gorm:"foreignKey:IuranID"`
}
