package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func main()  {
	sum := plus(1, 2)
	fmt.Println("a + b =", sum) // a + b = 3
}
