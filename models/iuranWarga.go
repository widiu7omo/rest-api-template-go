package models

type IuranWarga struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	WargaID     int64  `json:"warga_id"`
	IuranID     int64  `json:"iuran_id"`
	Warga       *Warga `json:"warga"`
	Iuran       *Iuran `json:"iuran"`
}
