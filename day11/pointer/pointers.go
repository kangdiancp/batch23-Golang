package main

import "fmt"

//copy by value
func copybyValue() {
	first := 100
	second := first
	first++
	fmt.Println(first, "Address 1st: ", &first)
	fmt.Println(second, "Address 2nd: ", &second)
}

//copy by references
func copybyPointer() {
	first := 100
	second := &first
	first++

	fmt.Println(first, "Address 1st: ", &first)
	fmt.Println(second, "Value: ", *second)
}

func chainPonter() {
	first := 100
	second := &first
	first++
	*second++
	third := &second
	**third++
	fmt.Println(first, "Address 1st: ", &first)
	fmt.Println(second, "Address 2nd: ", &second, "value: ", *second)
	fmt.Println(third, "Address 3rd: ", &third, "value: ", **third)

}

func pointerZeroValue() {
	first := 100
	var second *int
	fmt.Println("Second: ", second)
	fmt.Println("First: ", first)
}

// array pointer
func slicePointer() {
	fruit := []string{"Apel", "Mangga", "Melon"}
	fmt.Println("Fruits addr: ", &fruit)

	for key := range fruit {
		fmt.Printf("[%d] %p %s \n", key, &fruit[key], fruit[key])
	}
}

//append pointer array
func slicePointerAddr() {
	//declare extends
	//nums := make([]string, 3, 6)
	fruit := []string{"Apel", "Mangga", "Melon"}

	for key := range fruit {
		fmt.Printf("[%d] %p %s \n", key, &fruit[key], fruit[key])
	}

	fruit = append(fruit, "banana")

	for key := range fruit {
		fmt.Printf("[%d] %p %s \n", key, &fruit[key], fruit[key])
	}
}

func main() {
	slicePointerAddr()
	//copybyValue()
	//copybyPointer()
	//chainPonter()
	//pointerZeroValue()
	//slicePointer()
}
