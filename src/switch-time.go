package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 2

	fmt.Print("write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	TimeNow := time.Now().Weekday()
	if TimeNow == time.Saturday || TimeNow == time.Sunday {
		fmt.Println("今天是周末")
	} else {
		fmt.Println("今天是工作日, ", TimeNow)
	}

	Time := time.Now()

	fmt.Println("now time: ", Time.String())
}
