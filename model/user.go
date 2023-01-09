package model

type User struct {
	nickname string
	email    string
	password string
}

func (u *User) Nickname() string {
	return u.nickname
}

func (u *User) SetNickname(nickname string) {
	u.nickname = nickname
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}
