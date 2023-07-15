package main

import "fmt"

type User struct {
	UserName string
	Password string
}

func (user User) Login() {
	fmt.Println("User is logged In", user)
}

func (user User) Logout() {
	fmt.Println("User is logged Out")
}

type Admin struct {
	Role string
	User //  Parent: embedding User struct here so that we can reuse functionality from it
}

func (admin Admin) Login() { //overriding
	fmt.Println("User is logged In using Admin method", admin)
}



func main() {
	user1 := User{"Mukta", "1234"}
	//admin1 := Admin{user1, "Admin"}
	admin1 := Admin{Role: "Admin", User: user1}
	//	admin2 := Admin{Role: "Admin1"}
	admin1.Login()
	admin1.Logout()

	//admin2.Login()
}
