package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	var numData int

	println("Nomor 1 ----")
	fmt.Print("Masukkan input yang anda inginkan : ")
	fmt.Scan(&input)
	evenOrOddNumber(input)

	println("Nomor 2 ----")
	fmt.Print("Masukkan input yang anda inginkan : ")
	fmt.Scan(&input)
	powNumber(input)

	println("Nomor 3 ----")
	fmt.Print("Masukkan input yang anda inginkan : ")
	fmt.Scan(&input)
	fibonnaci(input)

	println("Nomor 4 ----")
	fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))

	println("Nomor 5 ----")
	fmt.Print("Masukkan Data yang anda inginkan : ")
	fmt.Scan(&numData)

	fmt.Println(nomor5Data(numData))
}

// Nomor 1
func evenOrOddNumber(n int) {
	tech := make([]int, n)
	tech = nil
	for i := 0; i < n; i++ {
		hitung := (i * 2) + 1
		tech = append(tech, hitung)
	}

	// }
	// Salah karna tidak sesuai dengan soal
	// return arr
	// arr := []int{}
	// for i := 0; i < n; i++ {
	// 	if i%2 != 0 {
	// 		arr = append(arr, i)
	// 	}
	// }
	fmt.Println(tech)
}

// Nomor 2
func powNumber(n int) {
	arrPow := make([]int, n)
	arrPow[0] = 1
	for i := 1; i < n; i++ {
		arrPow[i] = arrPow[i-1] * 2
	}

	fmt.Println(arrPow)
}

// Nomor 3
func fibonnaci(n int) {
	arrFibo := make([]int, n)
	arrFibo[0] = 2
	arrFibo[1] = 3

	for i := 2; i < n; i++ {
		arrFibo[i] = arrFibo[i-1] + arrFibo[i-2]
	}
	fmt.Println(arrFibo)
}

// Nomor 4
func reverseNumber(arrRev []int) []int {

	// firstDecalre := 0
	for i := 1; i < (len(arrRev)/2)-1; i++ {
		simpanArr0 := arrRev[0]
		arrRev[0] = arrRev[len(arrRev)-1]
		arrRev[len(arrRev)-1] = simpanArr0

		simpanArr1 := arrRev[1]
		arrRev[1] = arrRev[len(arrRev)-2]
		arrRev[len(arrRev)-2] = simpanArr1

		simpanArr2 := arrRev[2]
		arrRev[2] = arrRev[len(arrRev)-3]
		arrRev[len(arrRev)-3] = simpanArr2

	}
	return arrRev
}

// Nomor 5
func nomor5Data(n int) []string {
	arr := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if isPrimeNumber(i) {
			arr[i] = strconv.Itoa(i) + " = True"
		} else {
			arr[i] = strconv.Itoa(i) + " = False"
		}
	}

	return arr
}

func isPrimeNumber(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
