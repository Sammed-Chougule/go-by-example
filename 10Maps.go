package main

import "fmt"

func main() {
	a := make(map[int]int)

	a[1] = 6
	a[2] = 5

	fmt.Println(a, len(a))

	delete(a, 1)
	fmt.Println(a)

	b := map[int]int{1: 1, 2: 2}
	fmt.Println(b)
}
