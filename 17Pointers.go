package main

import "fmt"

func val(n int) {
	n = 0
}

func point(n *int) {
	*n = 0
}

func main() {
	i := 1
	fmt.Println(i)
	val(i)
	fmt.Println(i)
	point(&i)
	fmt.Println(i)
	fmt.Println(&i)

}
