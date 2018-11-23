package main

import "fmt"

func main() {

	sli := make([]string, 3)

	fmt.Println("empty slice: ", sli, ", len: ", len(sli))

	sli[0] = "a"
	sli[1] = "b"
	sli[2] = "c"

	fmt.Println("set slice: ", sli)
	fmt.Println("get slice[2]: ", sli[2])

	sli = append(sli, "d")
	sli = append(sli, "e")
	sli = append(sli, "f")

	fmt.Println("append slice: ", sli, ", len: ", len(sli))

	fmt.Println("sli1: ", sli[2:5])
	fmt.Println("sli2: ", sli[:5])
	fmt.Println("sli3: ", sli[2:])

	newSli := make([]string, len(sli))
	copy(newSli, sli)

	fmt.Println("new slice: ", newSli)

	arraySli := make([][]int, 3)

	for i := 0; i < 3; i++ {
		arraySli[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			arraySli[i][j] = i + j
		}
	}

	fmt.Println("array slice: ", arraySli)

}
