package main

import "fmt"

func main() {
	//print even number upto 10

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
