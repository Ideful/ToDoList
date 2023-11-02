package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Set(username, pwd string) {
	u.Username = username
	u.Password = pwd
}
