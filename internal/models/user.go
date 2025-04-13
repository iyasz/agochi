package models

import "database/sql"

type User struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `json:"name" gorm:"not null; type:varchar(255);"`
	Username  string       `json:"username" gorm:"not null; type:varchar(255); uniqueIndex"`
	Email     string       `json:"email" gorm:"not null; type:varchar(255); uniqueIndex"`
	Password  string       `json:"-"`
	Role      string       `json:"role" gorm:"not null; type:varchar(255);"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
