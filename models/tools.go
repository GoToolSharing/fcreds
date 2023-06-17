package models

type Tools struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
