package main

import "fmt"

func main() {

	//no 2
	//main1ver2(7, 1)

	//no 3
	//fmt.Println(findTarget([]int{1, 2, 4}, 4))
	//fmt.Println(findTarget([]int{-1, 2, 5, 6, 7}, 6))

	//no 4
	//fmt.Println(addSum([]int{2, 7, 11, 15}, 9))

	//no 5
	//fmt.Println(sumZero([]int{-1, 0, 1, 2, -1, 4}, 0))

	//no 6
	fmt.Println(plusOneDigit([]int{1, 2, 3}))

	//no 7
	// arr1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	// arr2 := []string{"Bayam", "Apel", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}
	// same, different := samaDanBeda(arr1, arr2)
	// fmt.Println("Same =", same)
	// fmt.Println("Different =", different)

	//no 8
	// fmt.Println(isPalindrome([]string{"asep", "budi", "-", "budi", "asep"}))
	// fmt.Println(isPalindrome([]string{"Tom", "Tim", "tim", "tom"}))
	// fmt.Println(isPalindrome([]string{"tik", "tok", "toko", "tik"}))

	//no9
	// fmt.Println(minMaxArray([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}))

	//no10
	// fmt.Println(sisipAngkaArr([]int{4, 7, 1, 20}, 9))
}

//func main1ver2(baris int, number int) {
//	for row := 1; row <= baris; row++ {
//		for cols := 1; cols <= row; cols++ {
//			fmt.Print(" ", number)
//			number++
//		}
//		number = number - (row - 1)
//		fmt.Println()
//	}
//}

//func findTarget(arr []int, target int) int {
//	for i, value := range arr {
//		if value == target {
//			return i
//		}
//	}
//	return -1
//}

//func addSum(arr []int, target int) []int {
//	for i := 0; i < len(arr); i++ {
//		for j := 1; j < len(arr); j++ {
//			if arr[i]+arr[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

//func sumZero(arr []int, target int) []int {
//	for i := 0; i < len(arr); i++ {
//		for j := 0; j < len(arr); j++ {
//			for k := 0; k < len(arr); k++ {
//				if arr[i]+arr[j]+arr[k] == target {
//					return []int{arr[i], arr[j], arr[k]}
//				}
//			}
//		}
//	}
//	return []int{}
//}

// soal no 6
func plusOneDigit(digit []int) []int {
	number := len(digit)
	core := 1

	for i := number - 1; i >= 0; i-- {
		sum := digit[i] + core
		digit[i] = sum % 10
		core = sum / 10
	}

	if core > 0 {
		digit = append([]int{core}, digit...)
	}
	return digit
}

// soal no 7
//func sameDifferent(arr1, arr2 []string) ([]string, []string) {
//same := []string{}
//different := []string{}

// Membuat slice kosong untuk menyimpan string yang sudah ditemukan
//seen := []string{}

// Mengecek setiap string pada arr1
//for _, str1 := range arr1 {
// Mengecek apakah string sudah pernah ditemukan sebelumnya
//found := false
//for _, seenStr := range seen {
//	if seenStr == str1 {
//		found = true
//		break
//	}
//}

// Jika string sudah pernah ditemukan, lanjut ke string berikutnya
//if found {
//	continue
//}

// Mengecek apakah string ada di arr2
//isSame := false
//for _, str2 := range arr2 {
//	if str1 == str2 {
//		isSame = true
//		break
//	}
//}

// Memasukkan string ke dalam slice yang sesuai
//if isSame {
//	same = append(same, str1)
//} else {
//	different = append(different, str1)
//}

// Menambahkan string ke dalam slice seen
//	seen = append(seen, str1)
//}

// Mengecek setiap string pada arr2
//for _, str2 := range arr2 {
// Mengecek apakah string sudah pernah ditemukan sebelumnya
//found := false
//for _, seenStr := range seen {
//	if seenStr == str2 {
//		found = true
//		break
//	}
//}

// Jika string sudah pernah ditemukan, lanjut ke string berikutnya
//if found {
//	continue
//}

// Mengecek apakah string ada di arr1
//isSame := false
//for _, str1 := range arr1 {
//	if str2 == str1 {
//		isSame = true
//		break
//	}
//}

// Memasukkan string ke dalam slice yang sesuai
//if isSame {
//	same = append(same, str2)
//} else {
//	different = append(different, str2)
//}

// Menambahkan string ke dalam slice seen
//		seen = append(seen, str2)
//	}
//
//	return same, different
//}

// soal no 8
//func isPalindrome(kata []string) bool {
//	length := len(kata)
//	for i := 0; i < len(kata); i++ {
//		kata[i] = strings.ToLower(kata[i])
//	}
//	for i := 0; i < length/2; i++ {
//		if kata[i] != kata[length-1-i] {
//			return false
//		}
//	}
//	return true
//}

// soal no 9
//func minMaxArray(arr []int) (int, int) {
//	min := arr[0]
//	max := arr[0]
//
//	for _, nums := range arr {
//		if nums < min {
//			min = nums
//		}
//		if nums > max {
//			max = nums
//		}
//	}
//	return min, max
//}

// soal no 10
//func sisipAngkaArr(arr []int, target int) []int {
//	// Menambahkan elemen target ke dalam array
//	arr = append(arr, target)
//
//	// Mengurutkan array secara ascending
//	sort.Ints(arr)
//
//	return arr
//}
