package main

import "fmt"

func main() {
	uri := "https://example.org:6001/v2/snacks?"
	r := "req=snickers"
	q := "quantity=1"
	s := "size=king"

	fmt.Printf("%s%s&%s&%s", uri, r, q, s)

}
