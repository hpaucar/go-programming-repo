package main

import "fmt"

func main() {
	
	var array = [6]int {1, 2, 3, 4, 5, 6}

	fmt.Println(array[-42]) // invalid array index -1 (index must be non-negative)
	fmt.Println(array[-1]) // invalid array index -1 (index must be non-negative)
	fmt.Println(array[0]) // > 1
	fmt.Println(array[1]) // > 2
	fmt.Println(array[2]) // > 3
	fmt.Println(array[3]) // > 4
	fmt.Println(array[4]) // > 5
	fmt.Println(array[5]) // > 6
	fmt.Println(array[6]) // invalid array index 6 (out of bounds for 6-element array)
	fmt.Println(array[42]) // invalid array index 42 (out of bounds for 6-element array)

//	To set or modify a value in the array, the way is the same.
//Example:

	fmt.Println(array) // > [1 2 3 4 5 6]

	array[0] := 6
	fmt.Println(array) // > [6 2 3 4 5 6]

	array[1] := 5
	fmt.Println(array) // > [6 5 3 4 5 6]

	array[2] := 4
	fmt.Println(array) // > [6 5 4 4 5 6]

	array[3] := 3
	fmt.Println(array) // > [6 5 4 3 5 6]

	array[4] := 2
	fmt.Println(array) // > [6 5 4 3 2 6]

	array[5] := 1
	fmt.Println(array) // > [6 5 4 3 2 1]
}