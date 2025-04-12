package models

import "database/sql"

type User struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
