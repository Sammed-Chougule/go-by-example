package main

import "fmt"

func twosum(a int, b int) int {
	return a + b
}

func threesum(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := twosum(3, 4)
	fmt.Println(res)

	res = threesum(5, 6, 7)
	fmt.Println(res)

}
