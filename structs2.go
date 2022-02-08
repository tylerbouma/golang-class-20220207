package main

import (
	"fmt"
	"os/user"
)

type Userdata struct {
	na string
	un string
	hd string
	em string
}

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	currentUser := Userdata{user.Uid, user.Username, user.HomeDir, "tylerbouma1@gmail.com"}

	fmt.Println("User id:", currentUser.na)        // recall current user id
	fmt.Println("Username:", currentUser.un)       // recall username
	fmt.Println("Home Directory:", currentUser.hd) // recall homedirectory
	fmt.Println("Email:", currentUser.em)
}
