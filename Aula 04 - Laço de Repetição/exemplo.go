package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 10; i <= j; i, j = i+1, j-1 {
		fmt.Println("i: ", i)
		fmt.Println("j: ", j)
	}
}
