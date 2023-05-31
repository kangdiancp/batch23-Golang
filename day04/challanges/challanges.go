package main

import "fmt"

// func sumArray()  {
// 	sumA := []int{1,2,3,4,5}

// }

//function dengan parameter array dengan retur array bertipe data integer
func initArray(n int) []int {
	arr := make([]int, 10)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ",")
	}
}

func sumAllElement(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func maxElement(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}

	}
	return max
}

func smallestIndex(arr []int) int {
	small := 0
	max := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			small = i
		}
	}
	return small
}

func shifting(arr []int, rotate int) []int {
	temp := 0
	panjangArr := len(arr)

	for i := 0; i < rotate; i++ {
		temp = arr[panjangArr-1]

		arr[panjangArr-1] = arr[0]

		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = temp
	}

	return arr

}

func linierSearch(arr []int, search int) int {
	temp := 0
	indexSearch := 0

	panjangArr := len(arr)

	temp = arr[panjangArr-1]

	arr[panjangArr-1] = arr[0]

	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = arr[3]
	arr[3] = temp
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			indexSearch = i
		}
	}

	return indexSearch
}

func main() {
	//list := initArray(5)
	//displayArray(list)
	//displayArray(initArray(10))
	//fmt.Println(list)
	num1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sumAllElement(num1[0:]))

	num2 := [5]int{97, 76, 10, 99, 10}
	fmt.Println(maxElement(num2[0:]))

	num3 := []int{1, 1, 12, 10, 5, 12}
	fmt.Println(smallestIndex(num3[0:]))

	num4 := []int{2, 3, 5, 7, 11}
	fmt.Println(shifting(num4[0:], 3))

	num5 := []int{2, 3, 5, 7, 11}
	fmt.Println(linierSearch(num5[0:], 7))

}
