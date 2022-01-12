package models

import "time"

type Book struct {
	ID            int           `json:"id" gorm:"primaryKey"`
	Title         string        `json:"title" gorm:"index"`
	Author        string        `json:"author"`
	Description   string        `json:"description"`
	NumberOfPages int           `json:"number_of_pages"`
	ReleaseDate   time.Time     `json:"release_date"`
	Comments      []BookComment `json:"comments" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
