package models

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type:varchar(225);unique;not null"`
}
