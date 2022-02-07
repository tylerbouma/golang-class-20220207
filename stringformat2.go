package main

import (
	"fmt"
)

func main() {
	n2 := 2
	n3 := 3
	n4 := 4

	res := fmt.Sprintf("there are %d oranges %d appes %d plums", n2, n3, n4)
	fmt.Println(res)

	res2 := fmt.Sprintf("There are %[2]d oranges %[3]d apples %[1]d plums", n2, n3, n4)
	fmt.Println(res2)
}
