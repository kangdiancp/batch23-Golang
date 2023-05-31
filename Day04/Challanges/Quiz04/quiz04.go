package main

import (
	"fmt"
	"strconv"
)

// Soal01 {1, 3, 5, 7, 9}
func evenOddNumber(count int) {
	Odd := []int{}

	for i := 0; i < count; i++ {
		counts := (i * 2) + 1
		Odd = append(Odd, counts)
	}
	fmt.Println(Odd)
}

// Soal02 {1, 2, 4, 8, 16, 32} = Bilang dipangkat 2
func powNumber(count int) {
	arr := make([]int, count)
	arr[0] = 1

	for i := 1; i < count; i++ {
		arr[i] = arr[i-1] * 2
	}
	fmt.Println(arr)
}

// Soal03 {2, 3, 5, 8, 13, 21}
func fibonnaci(count int) {
	arr := make([]int, count)
	arr[0] = 2
	arr[1] = 3

	for i := 2; i < count; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}

// Soal04 {4, 3, 6, 7, 8, 1} -Reverse Number -> {1, 8, 7, 6, 3, 4}
func reverseNumber(arr []int) []int {
	for i := 1; i <= len(arr)/2; i++ {
		temp := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = temp
	}
	return arr
}

// Soal05 Bilangan Prima
func isPrimeNumber(n int) []string {
	primeArray := make([]string, n+1) //Make([] string, length) = membuat array baru
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
	// evenOddNumber(5)
	// powNumber(6)
	// fibonnaci(6)
	// fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	fmt.Println(isPrimeNumber(10))
}
