package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	a, b := vals()
	fmt.Printf("a = %d b = %d\n",a, b) //a = 1 b = 2

	_, b2 := vals()
	fmt.Printf("b = %d\n", b2) // b = 2

	total := add(1, 2)
	fmt.Println(total) // 3

}
