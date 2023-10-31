package model

type User struct {
	Id int
	Username string
	Pwd string
}

func (u *User) Set(username,pwd string) {
	u.Username=username
	u.Pwd = pwd
}