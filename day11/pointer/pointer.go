package main

import "fmt"

func copyByValue() {
	first := 100
	second := first

	first++

	fmt.Println(first, "address 1st :", &first)
	fmt.Println(second, "address 2nd :", &second)
}

func copyByPointer() {
	first := 100
	second := &first

	first++
	fmt.Println(first, "address 1st :", &first)
	fmt.Println(second, "address 2nd :", &second, "value :", *second)
}

func chainPointer() {
	first := 100
	second := &first
	third := &second
	first++
	*second++
	**third++
	fmt.Println(first, "address 1st :", &first)
	fmt.Println(second, "address 2nd :", &second, "value :", *second)
	fmt.Println(third, "address 3nd :", &third, "value :", *third)
}

func pointerZero() {
	first := 100
	var second *int
	second = &first

	fmt.Println("Second :", *second)
	fmt.Println("First :", first)
}

func slicePointer() {
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("Fruit addr :", &fruit)

	for key := range fruit {
		fmt.Printf("[%d]%p %s \n", key, &fruit[key], fruit[key])
	}
}

func slicePointerAddr() {
	//nums := make([]string, 3, 6)
	fruit := []string{"apel", "mangga", "melon"}

	for key := range fruit {
		fmt.Printf("[%d]%p %s \n", key, &fruit[key], fruit[key])
	}

	fruit = append(fruit, "banana")

	for key := range fruit {
		fmt.Printf("[%d]%p %s \n", key, &fruit[key], fruit[key])
	}
}

func main() {
	//copyByValue()
	//copyByPointer()
	// chainPointer()
	// pointerZero()
	// slicePointer()
	slicePointerAddr()

}
