package main

import "fmt"

//Copy by value
func copyByValue() {
	first := 100
	second := first

	first++

	fmt.Println(first, "Address 1st :", &first)
	fmt.Println(second, "Address 2nd :", &second)
}

func copyByPointer() {
	first := 100
	second := &first

	first++
	fmt.Println(first, "Address 1st :", &first)
	fmt.Println(second, "Address 2nd :", &second, "Value : ", *second)
}

func chainPointer() {
	first := 100
	second := &first

	first++
	*second++
	third := &second
	**third++

	fmt.Println(first, "Address 1st :", &first)
	fmt.Println(second, "Address 2nd :", &second, "Value : ", *second)
	fmt.Println(third, "Address 3th :", &third, "Value : ", *third)
}

func pointerZeroValue() {
	first := 100
	// Jika kita coba akses, muncul panic error : invalid memory address
	var second *int

	fmt.Println("first : ", first)
	fmt.Println("second : ", second)
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
