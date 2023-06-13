package main

import "fmt"

//copy by value
func copyByValue() {
	first := 100
	second := first

	first++

	fmt.Println(first, "address 1st:", &first)
	fmt.Println(second, "address 2nd:", &second)
}

//copy by pointer
func copyByPointer() {
	first := 100
	second := &first

	first++

	fmt.Println(first, "address 1st:", &first)
	fmt.Println(second, "address 2nd:", &second, "value :", *second)
}

func chainPointer() {
	first := 100
	second := &first
	third := &second

	first++
	*second++
	**third++

	fmt.Println(first, "address 1st:", &first)
	fmt.Println(second, "address 2nd:", &second, "value :", *second)
	fmt.Println(third, "address 3th:", &third, "value :", *third)
}

func pointerZeroValue() {
	first := 100
	var second *int

	fmt.Println("first:", first)
	fmt.Println("second:", second)
}

func slicePointer() {
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("fruit addr:", &fruit)

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])

	}
}

func slicePointerAddr() {
	//declare
	//nums := make([]string, 3, 6)

	fruit := []string{"apel", "mangga", "melon"}

	fruit = append(fruit, "banana")

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}
}

func main() {
	slicePointerAddr()
	//slicePointer()
	// pointerZeroValue()
	//chainPointer()
	// copyByPointer()
	// copyByValue()
}
