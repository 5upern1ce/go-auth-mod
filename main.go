package main

import (
	"fmt"
)

func getUserCredentials() (string, string) {
	var username string
	var password string
	fmt.Print("Enter your username: \n -> ")
	fmt.Scan(&username)
	fmt.Print("Enter your password: \n -> ")
	fmt.Scan(&password)

	return username, password
}

func main() {
	username, password := getUserCredentials()
	hashedpass, _ := hashPassword(password)
	store(username, hashedpass)
}
