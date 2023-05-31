package main

import (
	"fmt"
	"strconv"
)

// Nomor 1:
func evenOrOddNumber(n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 2*i + 1
	}
	fmt.Println(arr)
}

// Nomor 2 :
func powNumber(n int) {
	arr := make([]int, n+1)
	power := 1
	for i := 0; i <= n; i++ {
		arr[i] = power
		power *= 2
	}
	fmt.Println(arr)
}

// Nomor 3 :
func fibonacci(n int) {
	fibo := make([]int, n)

	if n >= 1 {
		fibo[0] = 2
	}

	if n >= 2 {
		fibo[1] = 3
	}

	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}
	fmt.Println(fibo)
}

// Nomor 4 :
func reverseNumber(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

// Nomor 5 :
func isPrimeNumber(n int) []string {
	arr := make([]string, n+1)

	for i := 1; i <= n; i++ {
		if isPrime(i) {
			arr[i] = strconv.Itoa(i) + " = true, "
		} else {
			arr[i] = strconv.Itoa(i) + " = false, "
		}
	}

	return arr
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	evenOrOddNumber(5)
	powNumber(5)
	fibonacci(6)
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	fmt.Println(isPrimeNumber(10))
}
