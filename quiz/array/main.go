package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("nomor 1v2 : ", evenOrOddNumberCopy(5))

	fmt.Println("")
	fmt.Println("nomor 2 : ", powNumber(6))

	fmt.Println(" ")
	fmt.Println("nomor 3: ", fibonacci(6))

	fmt.Println(" ")
	nums4 := []int{4, 3, 6, 7, 8, 1}
	fmt.Println("nomor 4: ", reverseNumber(nums4))

	fmt.Println(" ")
	fmt.Println("nomor 5: ", isPrimeNumber(100))
}

func evenOrOddNumberCopy(n int) []int {
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

// NO 2
func powNumber(n int) []int {
	arr := make([]int, n)
	pow := 1

	for i := 0; i < n; i++ {
		arr[i] = pow
		pow *= 2
	}
	return arr
}

// NO 3
func fibonacci(n int) []int {
	arr := make([]int, n)

	if n >= 1 {
		arr[0] = 2
	}
	if n >= 2 {
		arr[1] = 3
	}

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr
}

// NO 4
func reverseNumber(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	return arr
}

// NO 5
func isPrimeNumber(n int) []string {
	arr := make([]string, n+1)

	for num := 1; num <= n; num++ {
		if isPrime(num) {
			arr[num] = strconv.Itoa(num) + " = true, "
		} else {
			arr[num] = strconv.Itoa(num) + " = false, "
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
