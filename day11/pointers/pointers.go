package main

import "fmt"

func main() {
	// copyByValue()
	// println()
	// copyByPointer()
	// println()
	// chainPointer()
	// pointerZeroValue()
	// slicePointer()
	slicePointerAddr()
}

//Fungsi Pointers
//menghindari copy data, variable kedua tidak akan ngecopy keseluruhan data di variable 1, variable ke 2 hanya ngepoint data variable 1

func copyByValue() {
	first := 100
	second := first
	// second := &first

	// first++

	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second)
}

func copyByPointer() {
	first := 100
	second := &first

	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second)
	//*(asterik) mengcopy nilai utuh dari variable yg diambil valuenya
	fmt.Println(*second, "address 2nd: ", &second)

}

func chainPointer() {
	first := 100
	second := &first

	*second++
	first++
	third := &second
	**third++

	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second)
	fmt.Println(third, "address 3rd: ", &third)

}

func pointerZeroValue() {
	first := 100
	var second *int

	second = &first

	fmt.Println("Second : ", *second)
	fmt.Println("First : ", first)
}

func slicePointer() {
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("fruits addr : ", &fruit)

	for key := range fruit {
		fmt.Printf("[%d], %p, %s\n", key, &fruit[key], fruit[key])
	}
}

func slicePointerAddr() {
	//declare slice if any capacity max for make a new slice
	//nums := make([]string, 3, 6)

	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("fruits addr : ", &fruit)

	for key := range fruit {
		fmt.Printf("[%d], %p, %s\n", key, &fruit[key], fruit[key])
	}
	println()
	fruit = append(fruit, "banana")
	for key := range fruit {
		fmt.Printf("[%d], %p, %s\n", key, &fruit[key], fruit[key])
	}
}
