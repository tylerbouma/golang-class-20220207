package main

import (
	"fmt"
)

func main() {
	var mychannel chan int
	fmt.Println("value of the channel: ", mychannel)
	fmt.Printf("type of the channel: %T", mychannel)

	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}
