package main

import "fmt"

type Role int

const (
	SuperAdmin Role = iota + 1
	Admin
	User
)

func (r Role) String() string {
	return [...]string{"super_admin", "admin", "user"}[r-1]
}

func main() {
	var r Role = Admin
	fmt.Println(r.String())
	fmt.Println(SuperAdmin.String())
}
