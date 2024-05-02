package main

import (
	"dsacourse/Search"
	"fmt"
)

func main() {
	fmt.Print("Hola!")

	var haystack []int
	haystack = []int{1, 3, 4, 6}
	fmt.Print(Search.LinearSearch(haystack, 5))
	fmt.Print(Search.BinarySearch(haystack, 3))

}
