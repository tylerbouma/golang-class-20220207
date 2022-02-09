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
		name   string
		colors []string
		p      Person `json:"Person"`
	}

	per := Person{
		Fn: "RZ",
		Ln: "Feeser",
	}

	group := ColorGroup{
		ID:     24601,
		name:   "Greens",
		colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
		p:      per,
	}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}
