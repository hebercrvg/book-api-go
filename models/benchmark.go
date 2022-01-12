package models

type Benchmark struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"index"`
	Description string `json:"description"`
}
