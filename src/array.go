package main

import "fmt"

func main() {
	var a [10]int

	fmt.Println("empty array: ", a)

	a[9] = 100

	fmt.Println("after set array: ", a)
	fmt.Println("array len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("new array: ", b)

	var c [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}

	fmt.Println("mutil array: ", c)
}
