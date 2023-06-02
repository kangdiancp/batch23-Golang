package main

import (
	"fmt"
	"strconv"
)

// numer 1
func evenOrOddNumber(n int) []int {
	arr := make([]int, n)
	if n%2 == 1 {
		arr[0] = 1
	} else {
		arr[0] = 2

	}

	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] + 2
	}
	return arr
}

//number2

func powNumber(n int) []int {
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] * 2
	}
	return arr
}

//number3

func fibonaci(n int) []int {
	arr := make([]int, n)
	arr[0] = 2
	arr[1] = 3
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

//number4

func reverseNumber(arr []int) []int {
	for i := 1; i <= len(arr)/2; i++ {
		temp := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = temp

	}
	return arr
}

//number5

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

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(evenOrOddNumber(5))
	fmt.Println(fibonaci(6))
	fmt.Println(powNumber(6))
	num4 := []int{4, 3, 6, 7, 8, 1, 2, 5}
	fmt.Println(reverseNumber(num4[0:]))
	fmt.Println(isPrimeNumber(50))

}
