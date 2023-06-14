package main

import "fmt"

// funciton initArray with 1 parameter & has return array
func initArray(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ", ")
	}
}

func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		//sum = sum + arr[i]
		sum += arr[i]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAllElement(nums))
	fmt.Println(sumAllElement([]int{12, 13, 14, 45, 32}))

	/* 	list := initArray(5)
	   	displayArray(list)
	   	fmt.Println()
	   	displayArray(initArray(10)) */
	//fmt.Println(list)
}
