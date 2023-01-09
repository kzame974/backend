package models

import (
	"time"
)

type UserEntity struct {
	ID        string `gorm:"primaryKey;"`
	nickname  string `gorm:"type:varchar(255);unique;not null"`
	email     string `gorm:"type:varchar(255);unique;not null"`
	password  string `gorm:"type:varchar(255);not null"`
	isActive  bool   `gorm:"type:bool;default:false"`
	createdAt time.Time
	updatedAt time.Time
}

func (u *UserEntity) Nickname() string {
	return u.nickname
}

func (u *UserEntity) SetNickname(nickname string) {
	u.nickname = nickname
}

func (u *UserEntity) Email() string {
	return u.email
}

func (u *UserEntity) SetEmail(email string) {
	u.email = email
}

func (u *UserEntity) Password() string {
	return u.password
}

func (u *UserEntity) SetPassword(password string) {
	u.password = password
}
