package entities

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Salt      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (u *User) IsExist() bool {
	return u.ID != 0
}
