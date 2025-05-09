package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `json:"name" gorm:"not null; type:varchar(255);"`
	Username  string    `json:"username" gorm:"not null; type:varchar(255); uniqueIndex"`
	Email     string    `json:"email" gorm:"not null; type:varchar(255); uniqueIndex"`
	Password  string    `json:"-"`
	Role      string    `json:"role" gorm:"not null; type:varchar(255);"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}