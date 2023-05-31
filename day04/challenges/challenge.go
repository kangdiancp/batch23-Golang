package main

import (
	"fmt"
)

// function init array with one parameter & has return array
func initArray(n int) [] int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func displayArray(arr []int){
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ",")
	}
}

// soal 1 {1, 2, 3, 4, 5} sum all element -> 15
func soal1SumAllElement(arr []int) int{
	sum := 0
	for i := 0; i < len(arr); i++ {
		//sum = sum + arr[i]
		sum += arr[i]
	}
	return sum 

}

// soal 2 {89, 76, 10, 99, 10} Cari Max element --> 99
func soal2MaxElement(arr []int) int{
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}else{
			continue
		}
	}
	return max
}

// soal 3 {1, 1, 12, 10, 5, 12} Smallest index (mencari index terkecil dari nilai terbesar) --> 2
func soal3SmallestIndex(arr []int) int{
	max := 0
	min := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			min = i
		}
	}
	return min
}

// Soal 4 {2, 3, 5, 7, 11} shifting memindahkan nilai dalam array --> {3, 5, 7, 11, 2}
func soal4Shifting(arr []int, rotate int) []int{
	var shiftArray []int
	for i := 0; i < rotate; i++ {
		shiftArray = arr[1:len(arr)]
		shiftArray = append(shiftArray, arr[0])
		arr = shiftArray
	}
	return shiftArray
}

// Soal 5 {2, 3, 5, 7, 11} -> {3, 5, 7, 11, 2} Linear search mencari nilai didalam array dan menampilkan indexnya ---> 2
func soal5LinearSearch(n []int, num int)int{
	var linearArray []int
	found := 0
	linearArray = n[1:len(n)]
	linearArray = append(linearArray, n[0])
	for i := 0; i < len(linearArray); i++ {
		if linearArray[i] == num {
			found = i
		}else{
			continue
		}
	}
	return found
}



func main() {
	// list := initArray(5)
	// fmt.Println(list)
	// displayArray(list)
	//displayArray(initArray(10))
	// nums := []int {1,2,3,4,5}

	// fmt.Println(soal1SumAllElement(nums))
	//fmt.Println(soal1SumAllElement([]int {1,2,3,4,5}))
	//fmt.Println(soal2MaxElement([]int {89, 76, 10, 99, 10}))
	//fmt.Println(soal3SmallestIndex([]int{1, 1, 12, 10, 5, 12}))
	//fmt.Println(soal4Shifting([]int{2, 3, 5, 7, 11}, 4))
	//fmt.Println(soal5LinearSearch([]int{2, 3, 5, 7, 11}, 7))

}