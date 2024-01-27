package main

import "fmt"

type User struct {
	name string
}

func (b User) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

func main() {
	user := User{name: "hoge"}
	fmt.Println(user.name)
	fmt.Println(user.cal(60, 170))
	foo(&user)
}

func foo(u *User) {
	fmt.Println(u.name)
}
