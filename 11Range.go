package main

import "fmt"

func main() {

	slices := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range slices { // "_" which is used to nullify index
		sum += num
	}
	fmt.Println(sum)

	for i, num := range slices {
		if i == 2 {
			fmt.Println(num)
		}
	}

	mp := map[int]string{1: "med", 2: "sam"}

	for k, v := range mp {
		fmt.Println(k, v)
	}

}
