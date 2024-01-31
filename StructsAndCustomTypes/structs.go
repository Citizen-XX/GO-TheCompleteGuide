package main

import (
	"fmt"
	"structs/StructsAndCustomTypes/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser,err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("text@example.com","test123")
	
	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
	// ... do something awesome with that gathered data!
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
