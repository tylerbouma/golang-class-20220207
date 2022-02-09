package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Person struct {
		Fn string
		Ln string
	}

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		P      Person `json:"Person"`
	}

	per := Person{
		Fn: "RZ",
		Ln: "Feeser",
	}

	group := ColorGroup{
		ID:     24601,
		Name:   "Greens",
		Colors: []string{"moss", "shamrock", "lime", "hunter"},
		P:      per,
	}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}
