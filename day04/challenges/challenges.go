package main

import "fmt"

func main() {
	// initArray(5)
	// fmt.Println(initArray(5))
	// displayArray(initArray(5))

	// fmt.Println(sumAllElement([]int{1, 2, 3, 4, 5}))  --> Contoh input parameter dengan tipe array

	println("Nomor 1----")
	fmt.Println(sumAllElement(initArray(5)))
	println()
	println("Nomor 2----")
	fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))
	println()
	println("Nomor 3----")
	fmt.Println(smallestIndex([]int{1, 1, 12, 10, 5, 12}))
	println()
	println("Nomor 4----")
	var rotasi int
	fmt.Print("Masukkan Rotasi yang anda inginkan : ")
	fmt.Scan(&rotasi)
	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, rotasi))
	println()
	println("Nomor 5----")
	fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}, 7))
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ", ")
	}
	println()
}

func initArray(n int) []int { //[]int mengharapkan nilai return dari hasil proses yg terjadi
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

//Nomor1
//1,2,3,4,5
func sumAllElement(arr []int) int {
	hasil := 0

	for _, value := range arr {
		hasil += value
	}

	// fmt.Println(hasil)

	return hasil
}

// func nomor1() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	hasil := 0

// 	for _, value := range nums {
// 		hasil += value
// 	}

// 	fmt.Println(hasil)
// }

//Nomor 2
//{89, 76, 10, 99, 10}
func maxElement(arr []int) int {
	hasil := 0

	for _, value := range arr {
		if hasil < value {
			hasil = value
		}
	}
	// fmt.Println(hasil)
	return hasil
}

//Nomor 3
//1,1,12,10,5,12
func smallestIndex(arr []int) int {
	hasil := 0
	// listMax := make([]int, 2)
	min := 0
	// listMin := make([]int, 2)
	for i, value := range arr {
		if hasil < value {
			hasil = value
			min = i
		}
		// listMax = []int{hasil}
		// listMin = []int{min}
	}
	// fmt.Println(hasil)
	// fmt.Println(min)
	return min
}

//Nomor 4
//2,3,5,7,11
func shifting(arr []int, rotate int) []int {
	// Case 1
	// firstDecalre := 0
	// lastIndex := len(arr) - 1
	// for i := 0; i < rotate; i++ {
	// 	simpanArr0 := arr[0]
	// 	arr[0] = arr[1]
	// 	arr[1] = arr[2]
	// 	arr[2] = arr[3]
	// 	arr[3] = arr[lastIndex]
	// 	arr[lastIndex] = simpanArr0
	// }

	// Case 2
	for i := 0; i < rotate; i++ {
		simpanArr0 := arr[0]
		replaceElmnt := arr[1:len(arr)]
		copy(arr, replaceElmnt)
		arr[len(arr)-1] = simpanArr0
	}
	//fmt.Println(arr)

	return arr

}

//Nomor 5
//3,5,7,11,2 --> cari index dengan value 7
func linearSearch(arr []int, cari int) int {
	// Case 1
	hasil := 0
	// lastIndex := len(arr) - 1
	// simpanArr0 := arr[0]
	// arr[0] = arr[1]
	// arr[1] = arr[2]
	// arr[2] = arr[3]
	// arr[3] = arr[lastIndex]
	// arr[lastIndex] = simpanArr0

	// Case 2
	simpanArr0 := arr[0]
	replaceElmnt := arr[1:len(arr)]
	copy(arr, replaceElmnt)
	arr[len(arr)-1] = simpanArr0

	for i, value := range arr {
		if cari == value {
			hasil = i
		}
	}

	//fmt.Println(hasil)

	return hasil
}
