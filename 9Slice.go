package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	a = append(a, 4)
	fmt.Println(a)

	b := []int{2, 4, 6}
	b = append(b, 8)
	fmt.Println(b)

	c := make([]int, 4)
	fmt.Println(c)
	c = append(c, 10, 20, 40)
	fmt.Println(c)

	copy(c, b)
	fmt.Println(c)

	d := c[:4]
	fmt.Println(d)

	e := c[2:6]
	fmt.Println(e)

}
