package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("it is one")
	case 2:
		fmt.Println("it is two")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")

	default:
		fmt.Println("it is weekeday")
	}
}
