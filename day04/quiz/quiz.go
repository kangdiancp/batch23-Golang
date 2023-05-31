package main

import "fmt"

func evenOrOddNumber(n int) {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = 2*i + 1
	}
	fmt.Println(slice)
}

func powNumber(n int) {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = 1 << i
	}
	fmt.Println(slice)
}

func fibonnaci(n int) {
	slice := make([]int, n)

	if n > 0 {
		slice[0] = 2
	}
	if n > 1 {
		slice[1] = 3
	}

	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	fmt.Println(slice)
}

func reverseNumber(arr []int) []int {
	for i := 1; i <= len(arr)/2; i++ {
		temp := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = temp

		// for i := 0; i < len(arr)/2; i++ {
		// 	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return arr
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeNumber(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		num := i + 1
		if isPrime(num) {
			slice[i] = fmt.Sprintf("%d = true", num)
		} else {
			slice[i] = fmt.Sprintf("%d = false", num)
		}
	}
	return slice
}

func main() {
	evenOrOddNumber(5)
	powNumber(6)
	fibonnaci(6)
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	fmt.Println(isPrimeNumber(8))
}
