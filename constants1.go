package main

import "fmt"

const Pi = 3.14

func main() {
	const MyURI = "http://example.com:5000/v2/"
	fmt.Println("The protocol, authority, and path /v2/ cannout be modified", MyURI)

	fmt.Println("Happy", Pi, "Day")

	const Xfiles = true
	fmt.Println("The truth is out there?", Xfiles)
}
