package models

// Jokes model
type Jokes struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}
