package main

import "fmt"

func main() {

	map1 := make(map[string]int)

	fmt.Println("empty map: ", map1)

	map1["k1"] = 1
	map1["k2"] = 2

	fmt.Println("set map1: ", map1)

	delete(map1, "k2")
	fmt.Println("deleted map: ", map1)

	_, exist := map1["k2"]

	if !exist {
		fmt.Println("no key 'k2' ")
	}
}
