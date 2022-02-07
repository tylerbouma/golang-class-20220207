package main

import "fmt"

func main() {
	var myFloat32 float32 = 4.5
	var myComplex64 complex64 = 4.5
	fmt.Println(myFloat32)
	fmt.Println(myComplex64)

	type AltaString string

	const myUntypedString = "Alta3 Research"
	var uts AltaString = myUntypedString

	fmt.Println(uts)

	var zString AltaString = "A string"
	fmt.Println(zString)

	// This won't work because we use a typed const
	// const typedInt int = 1
	// var myFloated64 float64 = typedInt
	// fmt.Println(myFloated64)

	// This now works because of using an untyped int
	const untypeInt = 1
	var myFloat64 float64 = untypeInt
	fmt.Println(myFloat64)
}
