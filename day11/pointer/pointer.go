package main

import "fmt"

//copy by value
func copyByValue(){
	first := 100
	second := first

	first++

	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second)
}

func copyByPointer(){
	first := 100
	second := &first

	first++
	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second, "\nValue 2nd : ", *second)
	
}

func chainPointer(){
	first := 100
	second := &first
	
	first++
	*second++
	third := &second
	**third++
	fmt.Println(first, "Address 1st : ", &first)
	fmt.Println(second, "Address 2nd : ", &second, "\nValue 2nd : ", *second)
	fmt.Println(third, "Address 3rd : ", &third, "\nValue 3rd : ", **third)
}

func pointerZeroValue(){
	first := 100
	var second *int
	second = &first

	fmt.Println("Second : ", *second)
	fmt.Println("First : ", first)
}

func slicePointer(){
	fruit := []string{"apel", "mangga", "melon"}
	fmt.Println("Fruit Address : ", &fruit)

	for key := range fruit {
		fmt.Printf("[%d] %p %s\n", key, &fruit[key], fruit[key])
	}
}

func slicePointerAddr(){
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
	// copyByValue()
	// copyByPointer()
	// chainPointer()
	// pointerZeroValue()
	// slicePointer()
	slicePointerAddr()
}