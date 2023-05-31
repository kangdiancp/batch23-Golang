package main

import (
	"fmt"
	"strconv"
)

func main() {
	// evenOrOddNumber(5)
	// powNumber(2)
	// fibonanci(6)
	// fmt.Println(reverseNumber([]int{4, 3, 6, 7, 8, 1}))
	fmt.Println(nomorData(5))
}

// no 1
func evenOrOddNumber(n int) {
	tech := make([]int, n)
	tech = nil
	for i := 0; i < n; i++ {
		count := (i * 2) + 1
		tech = append(tech, count)
	}
	fmt.Println(tech)

	// tech := make([]int, n)
	// tech = nil
	// for i := 0; i < n; i++ {
	// 	if i%2 == 0 {
	// 		tech[i] = 1
	// 	} else {
	// 		tech[i] = 0
	// 	}
	// }
	// for _, v := range tech {
	// 	fmt.Println(v)
	// }
}

// no 2
func powNumber(n int) {
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] * 2
	}
	fmt.Println(arr)
}

// tech := make([]int, n)
// for i := 0; i < n; i++ {
// 	tech[i] = i * i
// 	}
// 	return tech
//

// no 3
func fibonanci(n int) {
	arr := make([]int, n)
	arr[0] = 2
	arr[1] = 3

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}

// a, b := 0, 1
// for i := 0; i < n; i++ {
// 	fmt.Println(a)
// 	a, b = b, a+b
// 	}

// no 4
func reverseNumber(arr []int) []int {
	for i := 1; i < (len(arr)/2)-1; i++ {
		saveArr := arr[0]
		arr[0] = arr[len(arr)-1]
		arr[len(arr)-1] = saveArr

		saveArr0 := arr[1]
		arr[1] = arr[len(arr)-2]
		arr[len(arr)-2] = saveArr0

		saveArr1 := arr[2]
		arr[2] = arr[len(arr)-3]
		arr[len(arr)-3] = saveArr1
	}
	return arr
}

// no 5
func nomorData(n int) []string {
	arr := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if isPrimeNumber(i) {
			arr[i] = strconv.Itoa(i) + "=True"
		} else {
			arr[i] = strconv.Itoa(i) + "=False"
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

// 	if n < 2 {
// 		return []int{0, 0}
// 	}
// 	if n == 2 {
// 		return []int{1, 1}
// 	}
// 	if n > 2 {
// 		var arr []int
// 		for i := 2; i < n; i++ {
// 			if n%i == 0 {
// 				arr = append(arr, 0)
// 			} else {
// 				arr = append(arr, 1)
// 			}
// 		}
// 		return arr
// 	}
