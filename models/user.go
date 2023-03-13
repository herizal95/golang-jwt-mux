package models

import (
	"time"

	"github.com/satori/uuid"
)

type User struct {
	Uid       uuid.UUID `gorm:"primary_key" json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}
