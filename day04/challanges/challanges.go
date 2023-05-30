package main

import (
	"fmt"
)

// function initArray with 1 parameter & has return array
func initArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ", ")
	}
}

func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		// sum = sum + arr[i]
		sum += arr[i]
	}

	return sum
}

func maxElement(arr []int) int {
	max := 0
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

func smallestIndex(arr []int) int {
	max := 0
	minIndex := 0
	for i, num := range arr {
		if num > max {
			max = num
			minIndex = i
		} else if num == max {
			if i < minIndex {
				minIndex = i
			}
		}
	}

	return minIndex
}

func shifting(arr []int, rotate int) []int {
	length := len(arr)

	shiftArr := make([]int, length)

	startIndex := rotate % length

	for i := 0; i < length; i++ {
		shiftArr[i] = arr[startIndex]
		startIndex = (startIndex + 1) % length
	}

	return shiftArr
}

func linearSearch(arr []int, search int) int {
	value := 0
	for i, num := range arr {
		if num == search {
			value = i
		}
	}

	return value

}

func main() {
	list := initArray(5)
	displayArray(list)
	fmt.Println()
	displayArray(initArray(10))
	fmt.Println(list)

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAllElement(nums))
	fmt.Println(sumAllElement([]int{1, 2, 3, 4, 5}))

	fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))

	fmt.Println(smallestIndex([]int{1, 1, 12, 10, 12}))

	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 1))

	fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}, 7))
}
