package main

import (
	"fmt"
)

func evenOrOddNumber(n int) []int {
	oddNumber := make([]int, n)
	if n%2 == 1 {
		oddNumber[0] = 1
	}
	for i := 1; i < n; i++ {
		oddNumber[i] = oddNumber[i-1] + 2
	}
	return oddNumber
}

func powNumber(n int) []int {
	powNum := make([]int, n)
	powNum[0] = 1
	for i := 1; i < n; i++ {
		powNum[i] = powNum[i-1] * 2
	}
	return powNum
}

func fibonnaci(n int) []int {
	fb := make([]int, n)
	fb[0] = 2
	fb[1] = 3
	for i := 2; i < n; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}

func reverseNumber(arr []int) []int {
	for i := 1; i < len(arr)/2; i++ {
		index := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = index
	}
	return arr
}

func main() {
	//fmt.Println(evenOrOddNumber(5))
	//fmt.Println(powNumber(6))
	//fmt.Println(fibonnaci(6))
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
}
