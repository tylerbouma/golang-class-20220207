package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Host string
		Mac  string
		Ram  int
		Cpu  int
	}{
		{"merlin", "2e549138a9e3", 1024, 1},
		{"prospero", "3c539188c9e3", 2048, 2},
		{"nestor", "1b166127a9e3", 2048, 1},
		{"odin", "4d545b88c9e3", 4096, 2},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Host < people[j].Host })
	fmt.Println("by hostname:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Ram < people[j].Ram })
	fmt.Println("by ram:", people)

	algo := func(i, j int) bool { return people[i].Mac > people[j].Mac }
	sort.Slice(people, algo)
	fmt.Println("Reverse sort by MAC:", people)

}
