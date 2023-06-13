package main

import "fmt"

//copy by value
func copyByValue() {
	first := 100
	second := &first

	first++
	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second)
}

func copyByPointer() {
	first := 100
	second := &first

	first++
	*second++

	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second, "value :", *second)
}

func chainPointer() {
	first := 100
	second := &first

	first++

	*second++
	third := &second
	**third++
	fmt.Println(first, "address 1st: ", &first)
	fmt.Println(second, "address 2nd: ", &second, "value :", *second)
	fmt.Println(third, "address 3nd: ", &third, "value :", *third)
}

func pointerZerpValue() {
	first := 100
	var second *int
	second = &first

	fmt.Println("second: ", second)
	fmt.Println("first:", first)

}

func slicePointer() {
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("fruis addr:", &fruit)
	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}

}

func slicePointeAddrr() {
	//declare
	//nums := make([]string, 3, 6)

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
	//copyByValue()
	//copyByPointer()
	//chainPointer()
	//pointerZerpValue()
	//slicePointer()
	//slicePointeAddrr()
}
