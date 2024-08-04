package main

import "fmt"

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

	// we can have a part (slice) of an array as a separate slice

	var numbers4 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice_number = numbers4[2:8]

	fmt.Printf("%v \n", slice_number)
	fmt.Printf("%d and %d \n", len(slice_number), cap(slice_number))

	// check out the cap, as you can see it determines the cap based on the the first segment of the slice to the last element of the original array.
	// any change in the slice can change the original array up to its capacity.

	slice_number[3] = 888
	fmt.Printf("slice is %v and the original array is %v ", slice_number, numbers4)

	slice_number = append(slice_number, 777)
	fmt.Printf("slice is %v and the original array is %v ", slice_number, numbers4)
	print(slice_number)
	print(numbers4)

	slice_number = append(slice_number, 666)
	fmt.Printf("slice is %v and the original array is %v ", slice_number, numbers4)
	// you could see it no longer changed the original array because it surpassed its original capacity.
	print(slice_number)
	print(numbers4)

	// difference between iteration in the conventional way and using range

	// range indeed just evaluates the slice on time and then iterate over it and any change in the slice after the loop wont affect the loop.
	for i, v := range slice_number {
		fmt.Printf("%v, %v \n", i, v)
		slice_number = append(slice_number, 8)

	}
	fmt.Println(slice_number)

	// but in the conventional loops, it will iterate over the appended values.

}
