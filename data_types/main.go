package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// when you initiate a variable in this case a string, a part of memory gets allocated to this variable
	// and if you try to add another string to it, go normally tries to allocate the same location in memory to it
	// as an optimization method for memory.

	first_name := "Sato"
	// print location in memory
	fmt.Println(&first_name)

	first_name = first_name + "Karashi"
	// print location in memory
	fmt.Println(&first_name)

	// string concatination the normal way and its performance

	start := time.Now()
	for i := 0; i < 1000; i++ {
		first_name += "A"
	}

	fmt.Println(time.Since(start))

	// string concatination using the strings from standard library
	var sb_string strings.Builder
	sb_string.WriteString(first_name)

	start = time.Now()
	for i := 0; i < 1000; i++ {
		sb_string.WriteString("A")
	}

	fmt.Println(time.Since(start))

	// String package made a huge difference in performance and it happens to string comparison too.

	// integer overflow occurs in compile time but in the runtime it doesn't panic

	var first_int int8 = 127
	fmt.Println(&first_int)
	b := first_int + 1
	fmt.Println(&b)

	// run data type: it is identified with single quotation and if you print it, it will show you its asci and if try to get its length,
	// it will give an unexpected length, because length gives you result based on number of bytes.

	firt_run := 'A'
	print(firt_run)
	fmt.Println(firt_run)

}
