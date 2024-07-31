package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	// One way to initialize arrays. Arrays are fixed size and the size and the type together identify the array type
	var numbers [5]int
	fmt.Println(numbers)

	// another way

	var numbers2 = [5]int{4: 22}
	fmt.Println(numbers2)

	// assigning a an array to another value, will copy it by value and it will have another space in memory and they don't point to the same space like Python

	numbers3 := numbers2
	// as you can see they point to different parts of the memory

	fmt.Printf("%p \n", &numbers2)
	fmt.Printf("%p \n", &numbers3)

	// in go we only have for loops for iteration

	// conventional way
	for i := 0; i < 10; i++ {
		print(i)
	}

	// conditional way

	// for number == 5 {
	// 	fmt.Println("Number is still 5")
	// }

	// iterate with range

	for range numbers {
		fmt.Printf("numbers \n")
	}

	for index, value := range numbers {
		fmt.Printf("%d %d \n", index, value)
	}

	for _, value := range numbers {
		print(value)
	}

	// multi dimensional arrays, a small project working with image to show the turn the multi dim array into image pixels and turning that into an image

	data := [][]uint8{
		{35, 130, 45},
		{150, 42, 150},
		{44, 140, 39},
	}

	img := image.NewGray(image.Rect(0, 0, len(data[0]), len(data)))

	for y, row := range data {
		for x, pixel := range row {
			img.Set(x, y, color.Gray{Y: pixel})
		}
	}

	f, err := os.Create("image.png")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

}
