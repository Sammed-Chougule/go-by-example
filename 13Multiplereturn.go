package main

import "fmt"

func val() (int, int) {
	return 3, 7
}

func main() {
	a, b := val()
	fmt.Println(a, b)
	_, c := val()
	fmt.Println(c)
}
