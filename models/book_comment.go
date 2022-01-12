package models

import "time"

type BookComment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"index"`
	BookID    int       `json:"book_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
