package model

type Score struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}
