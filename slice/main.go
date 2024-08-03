package main

func main() {
	// slices are dynamic arrays and it will create a reference system in memory which can be useful and dangerous at the same time

	// each slice in memory gets saved by three values, 1. pointer 2.length 3. capacity

	var numbers []int
	print(numbers)

	// as you can see it doesn't point to any where and it is null this way

	var numbers2 []int = []int{}
	print(numbers2)
	// but here it points to somewhere in memory, which means it is not null this way

	// another way to define slices giving length and capacity
	var numbers3 []int = make([]int, 3, 8)
	print(numbers3)

}
