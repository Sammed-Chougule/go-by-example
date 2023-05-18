package main

import "fmt"

func sum(num ...int) {
	fmt.Println(num)
	target := 0

	for _, v := range num {
		target += v
	}
	fmt.Println(target)

}

func main() {
	sum(2, 3)

	a := []int{5, 6, 7, 8}

	sum(a...)
}
