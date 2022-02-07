package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("bool: %t\n", true)

	fmt.Printf("int : %d\n", 123)

	fmt.Printf("bin: %b\n", 14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3 :%x\n", "hex this")

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	fmt.Printf("wdith2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("left-justify: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("left-justify2: |%6s|%6s|\n", "foo", "bar")

	fmt.Printf("left-justify3: |%-6s|%-6s|\n", "foo", "bar")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
