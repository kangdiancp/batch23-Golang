package main

import "fmt"

// Copy by value
func copyByValue() {
	first := 100
	second := first

	first++

	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second)
}

// Copy by pointers
func copyByPointers() {
	first := 100
	second := &first

	first++
	*second++

	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second, ", Value : ", *second)
}

// Chain pointers
func chainPointers() {
	first := 100
	second := &first
	third := &second

	first++
	*second++
	**third++

	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second, ", Value : ", *second)
	fmt.Println(third, "Address 3rd : ", &third, ", Value : ", *third)
}

func pointerZeroValue() {
	first := 100

	// Panic Error : Invalid Memory Address
	var second *int

	second = &first

	fmt.Println("Second : ", second)
	fmt.Println("First : ", first)
}

func slicePointer() {
	fruits := []string{"mangga", "pokat", "melon"}

	fmt.Println("Fruits addr : ", &fruits)

	for i := range fruits {
		fmt.Printf("[%d] %p %s \n", i, &fruits[i], fruits[i])
	}
}

func slicePointerAddr() {
	// Declare Slice
	// nums := make([]string, 3, 6)

	fruits := []string{"mangga", "pokat", "melon"}

	for i := range fruits {
		fmt.Printf("[%d] %p %s \n", i, &fruits[i], fruits[i])
	}

	fruits = append(fruits, "naga")

	for i := range fruits {
		fmt.Printf("[%d] %p %s \n", i, &fruits[i], fruits[i])
	}
}

func main() {
	copyByValue()

	fmt.Println()

	copyByPointers()

	fmt.Println()

	chainPointers()

	fmt.Println()

	pointerZeroValue()

	fmt.Println()

	slicePointer()

	fmt.Println()

	slicePointerAddr()
}
