package main

import (
	"fmt"
	"strconv"
)

func evenOddNumber(n int) []int {
	index := make([]int, n)
	for i := 1; i < n; i++ {
		if n%2 == 1 {
			index[0] = 1
			index[i] = index[i-1] + 2
		}
	}
	return index
}

func powNumber(n int) []int {
	index := make([]int, n)
	index[0] = 1
	for i := 1; i < n; i++ {
		index[i] = index[i-1] * 2
	}
	return index
}

func fibonacci(n int) []int {
	index := make([]int, n)
	index[0] = 2
	index[1] = 3
	for i := 2; i < n; i++ {
		index[i] = index[i-1] + index[i-2]
	}
	return index
}

func reverseNumber(arr []int) []int {
	var index []int
	for i := len(arr) - 1; i >= 0; i-- {
		index = append(index, arr[i])
	}
	return index
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeNumber(n int) []string {
	primeArray := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			primeArray[i] = strconv.Itoa(i) + " = true,"
		} else {
			primeArray[i] = strconv.Itoa(i) + " = False,"
		}
	}
	return primeArray
}

func main() {
	// fmt.Println(evenOddNumber(5))
	// fmt.Println(powNumber(6))
	// fmt.Println(fibonacci(5))
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	//fmt.Println(isPrimeNumber(15))

}
