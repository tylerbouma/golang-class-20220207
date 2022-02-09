package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating file...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing employee list to file...")
	fmt.Fprintln(f, "Cloud Strife\nAerith Gainsborough\nRed XIII\nVincent Valentine")
}

func closeFile(f *os.File) {
	fmt.Println("closing file...")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("./employeeList.txt")

	defer closeFile(f)

	writeFile(f)
}
