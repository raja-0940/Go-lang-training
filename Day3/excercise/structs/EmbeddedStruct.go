package main

import "fmt"

type User struct {
	UserName string
	Passsword string
}

func (user User) Login(){
	fmt.Println("User logged in", user)
}

func (user User) Logout(){
	fmt.Println("User logged out")
}

type  Admin struct {
	User //Parent : embedding User struct here so that we can reuse the functionality from it
	Role string
}

func (admin Admin) Login(){ //overrriding
	fmt.Println("Admin logged in", admin)
}

func main() {
	user1 := User{"Raja","myPassword"}
	admin1 := Admin{Role: "Admin", User: user1}// admin1 := Admin{user1,"Admin"}
	admin2 := Admin{Role: "Dev"}

	admin1.Login()
	admin1.Logout()

	admin2.Login()
}