package main

import "fmt"

func main() {
	// Implicit variable definition is not wise always and mostly in cases where you are not sure
	// about the type, you use them. In implicit conversion, the type actually gets defined after the first use
	// not after initiating it.

	const a = 12
	var y float32 = a + 13.2
	fmt.Println(y)

	const b = '3'
	var x int = b + 15
	fmt.Print(x)

}
