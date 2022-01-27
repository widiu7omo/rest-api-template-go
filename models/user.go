package models

// User model
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	PushToken string `json:"pushToken" gorm:"default:''"`
	ApiToken  string `json:"apiToken" gorm:"default:''"`
}
