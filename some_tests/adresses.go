package main

import "fmt"

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Set(username, pwd string) {
	fmt.Printf("%p\n", u)
	u.Username = username
	u.Password = pwd
}

func (u User) Setnoptr(username, pwd string) {
	fmt.Printf("%p\n", &u)
	u.Username = username
	u.Password = pwd
}

func main() {
	a := User{
		Id:       2,
		Username: "qwe",
		Password: "zxc",
	}
	fmt.Printf("%p\n", &a)
	fmt.Println(a)
	a.Set("q","w")
	a.Setnoptr("r","w")
	fmt.Println(a)
}
