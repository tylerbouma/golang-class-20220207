package main

import "os"

func main() {
	_, err := os.Create("/carolDanvers/msmarvel")

	if err != nil {
		panic(err)
	}
}
