package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	a[4] = 200
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println(len(b))

	twod := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 0}}
	fmt.Println(twod)

	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println((twod))
}
