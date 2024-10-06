package models

import "time"

type User struct {
	Id       uint64    `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	BirthDay time.Time `json:"birthDay"`
}
