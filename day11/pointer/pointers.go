package main

import "fmt"

// copy by value
func copyByValue() {
	first := 100
	second := first

	first++
	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second)
}

// copy by pointer
func copyByPointer() {
	first := 100
	second := &first

	first++
	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(
		second, "address 2nd: ", &second,
		"value :", *second)
}

// chain pointer
func chainPointer() {
	first := 100
	second := &first

	first++

	*second++
	third := &second
	**third++
	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second,
		"value :", *second)
	fmt.Println(third, "address 3rd: ", &third,
		"value :", *third)
}

// pointer zero value
func pointerZeroValue() {
	first := 100
	// jika kita coba akses, muncul panic error, :invalid memory address
	var second *int
	second = &first

	fmt.Println("second : ", *second)
	fmt.Println("first : ", first)
}

func slicePointer() {
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("fruits addr: ", &fruit)

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}
}

func slicePointerAddr() {
	fruit := []string{"apel", "mangga", "melon"}

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}

	fruit = append(fruit, "banana")

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}
}

func main() {
	// copyByValue()
	// copyByPointer()
	// chainPointer()
	// pointerZeroValue()
	slicePointer()
}
