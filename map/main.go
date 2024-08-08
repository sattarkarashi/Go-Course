package main

import (
	"fmt"
	"maps"
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
	fmt.Printf("v1: %d and status is:%v \n", v2, status2)

	// delete items
	delete(my_map, "sati")
	fmt.Printf("%v", my_map)

	// clear all the map
	clear(my_map)
	fmt.Printf("%v \n", my_map)

	// using maps standard library
	my_map_2 := map[string]int{"x": 4, "y": 8}

	if maps.Equal(my_map, my_map_2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// iteration over maps doesn't guarantee the order
	for k, v := range my_map_2 {
		fmt.Printf("%s -> %d \n", k, v)
	}

	// in GO like other languages maps are being hashed using hash functions. keys go through the hash function which lets the access to variable faster.
	// Giving size to maps will make processing maps faster.

}
