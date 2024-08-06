package main

import (
	"fmt"
)

func main() {
	// maps

	my_map := map[string]int{"sato": 1, "sati": 2}

	fmt.Println("map:", my_map)

	// define map using make, note: we are free to give it len and giving len makes it perform better in memory

	make_map := make(map[string]int, 2)
	fmt.Println("map:", make_map)

	// getting items, it actually copies the values

	v1, status := my_map["sato"]
	fmt.Printf("v1: %d and status is:%v \n", v1, status)

	// if we try to get a value that doesn't exist it will return false and create a key with initial value
	v2, status2 := my_map["saya"]
	fmt.Printf("v1: %d and status is:%v ", v2, status2)

}
