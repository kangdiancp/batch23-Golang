package main

import (
	"fmt"
	"sort"
	"strings"
)

// nomor 2

func soal2looping(){
	var matrix [7][7] int
	nilai := 1
	matrix[0][0] = nilai
	for row := 1; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == row {
				matrix[row][col] = matrix[row-1][col-1] + 2
			}else if row > col{
				matrix[row][col] = (nilai + row) + col
			}else if col == 0 && row != 0{
				matrix[col][row] = matrix[row-1][col-1] + 1
			}

		}
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			if Colom > row{
				fmt.Print(" ")
			}else{
				fmt.Printf("%d ", matrix[row][Colom])
			}
		}
		fmt.Println()
	}
}

func soal3FindTarget(arr []int, n int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n{
			fmt.Print(i)
		}
	}
}

func soal4AddSum(arr []int, n int){
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i] + arr[j] == n{
				fmt.Printf("%d, %d \n", i, j)
			}
		}
	}
}

func soal5SumZero(arr[] int)[] int{
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			for k := 2; k < len(arr); k++ {
				if arr[i] + arr[j] + arr[k] == 0{
					// fmt.Println(arr[i], arr[j], arr[k])
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return[]int{}
}

func soal6plusOneDigit(arr []int)[] int{
	long := len(arr)
	count:= 1
	for i := long - 1; i >= 0; i-- {
		jml := arr[i] + count
		arr[i] = jml % 10
		count = jml / 10
		
	}
	if count > 0 {
		arr = append([]int{count}, arr...)
	}
	return arr
}

func soal7samaDanBeda(arr1, arr2 [] string) {
	// var aray1 []string
	// var aray2 []string
	var same []string
	var different []string
	var gabungan []string

	// for i := 0; i < len(arr1); i++ {
	// 	for j := i+1; j < len(arr1); j++ {
	// 		if arr1[i] != arr1[j] {
	// 			aray1 = append(aray1, arr1[j])
	// 			break
	// 		}
	// 	}
	// }
	// for i := 0; i < len(arr2); i++ {
	// 	for j := i+1; j < len(arr2); j++ {
	// 		if arr2[i] != arr2[j] {
	// 			aray2 = append(aray2, arr2[j])
	// 			break
	// 		}
	// 	}
	// }

	// for i := 0; i < len(aray1); i++ {
	// 	for j := 0; j < len(aray2); j++ {
	// 		if aray1[i] == aray2[j]{
	// 			same = append(same, aray1[i])
	// 		}
	// 	}
	// }

	// gabungan = append(gabungan, aray1...)
	// gabungan = append(gabungan, aray2...)

	gabungan = append(gabungan, arr1...)
	gabungan = append(gabungan, arr2...)

	// for _, val1 := range gabungan {
	// 	found := false
	// 	for _, val2 := range same {
	// 		if val1 == val2 {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	if !found {
	// 		different = append(different, val1)
	// 	}
	// }

	mapfruit := map[string]int{}

	// declare mapfruit to store key, value
	for key := range gabungan {
		mapfruit[gabungan[key]]++
	}
	fmt.Println(mapfruit)


	for key, v := range mapfruit {
		if v > 1{
			same = append(same, key)
		}else {
			different = append(different, key)
		}
	}
	sort.Strings(same)
	sort.Strings(different)
	
	fmt.Println("Same = ", same)
	fmt.Println("Different = ", different)
}

func soal8isPalindrome(arr []string) bool {
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}

	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}

	return true
}

func soal9minMaxArray(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}


func soal10insertElement(arr []int, target int) []int {
	// Menambahkan elemen target ke dalam array
	arr = append(arr, target)

	// Mengurutkan array secara ascending
	sort.Ints(arr)

	return arr
}


func main(){
//  soal2looping()
	// soal3FindTarget([]int {1, 2, 4}, 4)
	// soal3FindTarget([]int{-1, 2, 5, 6, 7}, 6)
	// soal4AddSum([]int{3, 2, 3}, 6)
	// fmt.Println(soal5SumZero([]int {-1, 0, 1, 2, -1, 4}))
	//fmt.Print(soal6plusOneDigit([]int{1, 2, 3, 9}))

	// 7
	soal7samaDanBeda([]string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}, 
					 []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"})

	//8
	// arr := []string{"asep", "budi", "-", "budi", "asep"}
	// result := soal8isPalindrome(arr)
	// fmt.Println(result)


	// 9
	// array1 := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	// min1, max1 := soal9minMaxArray(array1)
	// fmt.Println("min =", min1)
	// fmt.Println("max =", max1)

	// 10
	// array := []int{4, 7, 1, 20}
	// target := 9
	// result := soal10insertElement(array, target)
	// fmt.Println(result)

}