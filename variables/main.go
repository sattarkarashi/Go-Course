package main

func main() {
	// Explicit var definition
	var a int = 10
	var b int

	print(a)
	print(b)

	c := 6
	print(c)

	d := &a
	print(d)
	print(*d)

	*d = 30
	print(d)
	print(a)

}
