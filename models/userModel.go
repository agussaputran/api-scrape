package models

// Users model
type Users struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
