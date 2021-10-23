package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := int(1)
	y := int(2)
	// variable addresses
	swap(&x, &y)
	fmt.Println(x, y)
}
