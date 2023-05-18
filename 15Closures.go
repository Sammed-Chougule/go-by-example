package main

import "fmt"

func incre() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	called := incre()

	fmt.Println(called())
	fmt.Println(called())

}
