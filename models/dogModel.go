package models

// Dogs Model
type Dogs struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	Message string `json:"message"`
}
