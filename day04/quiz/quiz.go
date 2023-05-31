package main

import (
	"fmt"
	"strconv"
)

func evenOrOddNumber(n int) {
	angka := []int{}

	for i := 0; i < n*2; i++ {
		if i%2 != 0 {
			angka = append(angka, i)
		}
	}
	fmt.Print(angka)
}

func powNumber(n int) {
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] * 2
	}
	fmt.Print(arr)
}

func fibonacci(n int) {
	hasil := make([]int, n)
	hasil[0] = 2
	hasil[1] = 3
	for i := 2; i < n; i++ {
		hasil[i] = hasil[i-1] + hasil[i-2]
	}
	fmt.Print(hasil)
}

func reverseNumber(arr []int) []int {
	for i := 1; i <= len(arr)/2; i++ {
		x := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = x
	}
	return arr
}

func isPrimeNumber(n int) []string {
	primeArr := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			primeArr[i] = strconv.Itoa(i) + " = true|"
		} else {
			primeArr[i] = strconv.Itoa(i) + " = False|"
		}
	}
	return primeArr
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
	powNumber(6)
	fibonacci(6)
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	fmt.Println(isPrimeNumber(10))
}
