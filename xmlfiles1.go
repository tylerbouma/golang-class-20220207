package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Subject string `xml:"subject"`
	Body    string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("avengers.xml")

	note := &Notes{}

	_ = xml.Unmarshal(data, &note)

	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Subject)
	fmt.Println(note.Body)
}
