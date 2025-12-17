package main

import(
	"fmt"
	"example.com/structs/user"
)



func main() {
	var firstName, lastName string
	var age int
	fmt.Print("Enter your first name, last name and age:")
	fmt.Scanln(&firstName, &lastName, &age)

	user1, err := user.New(firstName, lastName, age)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	user1.GetUserInfo()
	user1.ClearUserInfo()

	admin := user.NewAdmin("test@example.com", "password123")
	admin.GetUserInfo()
	admin.ClearUserInfo()
}

