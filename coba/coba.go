package main

import (
	"fmt"
	"sort"
)

// Nomor 2 :
// Fungsi ini mencetak pola segitiga dengan angka yang bertambah di setiap barisnya
func printPattern() {
	rows := 7  // Mengatur jumlah baris
	start := 1 // Angka awal

	// Melakukan iterasi sebanyak jumlah baris
	for i := 0; i < rows; i++ {
		count := start + i // Nilai awal untuk setiap baris
		// Mencetak angka di setiap baris
		for j := 0; j <= i; j++ {
			fmt.Printf("%-3d", count) // Mencetak angka
			count++                   // Menambahkan angka setelah mencetak
		}
		fmt.Println() // Membuat baris baru setelah setiap baris
	}
}

// Nomor 3 :
// Fungsi ini mencari target dalam array dan mengembalikan indeksnya
func findTarget(arr []int, target int) int {
	index := -1 // Nilai default jika target tidak ditemukan
	for i := 0; i < len(arr); i++ {
		if arr[i] == target { // Jika target ditemukan, mengembalikan indeksnya
			return i
		}
	}
	return index // Mengembalikan nilai default jika target tidak ditemukan
}

// Nomor 4 :
// Fungsi ini mencari dua angka dalam array yang jumlahnya sama dengan target dan mengembalikan indeksnya
func addSum(nums []int, target int) []int {
	slice := []int{} // Inisialisasi slice kosong
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target { // Jika dua angka yang jumlahnya sama dengan target ditemukan, mengembalikan indeksnya
				return []int{i, j}
			}
		}
	}
	return slice // Mengembalikan slice kosong jika tidak ada dua angka yang jumlahnya sama dengan target
}

// Nomor 5 :
// Fungsi ini mencari tiga angka dalam array yang jumlahnya sama dengan 0 dan mengembalikan angka-angka tersebut
func sumZero(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 { // Jika tiga angka yang jumlahnya sama dengan 0 ditemukan, mengembalikan angka-angka tersebut
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return []int{} // Mengembalikan slice kosong jika tidak ada tiga angka yang jumlahnya sama dengan 0
}

// Nomor 6 :
// Fungsi ini menambahkan satu ke array dari digit terakhir dan mengembalikan array baru
func plusOne(digits []int) []int {
	return increment(digits) // Memanggil fungsi increment dan mengembalikan hasilnya
}

// Fungsi ini menambahkan satu ke array dari digit terakhir dan mengembalikan array baru
func increment(digits []int) []int {
	if len(digits) == 0 { // Jika array kosong, mengembalikan array dengan satu elemen [1]
		return []int{1}
	}

	last_digit := digits[len(digits)-1] // Mendapatkan digit terakhir

	last_digit += 1     // Menambahkan satu ke digit terakhir
	if last_digit > 9 { // Jika hasilnya lebih besar dari 9 (perlu dilakukan carry)
		last_digit = 0                              // Membuat digit menjadi 0
		result := increment(digits[:len(digits)-1]) // Melakukan rekursi ke elemen sebelumnya
		return append(result, last_digit)           // Menggabungkan hasil rekursi dengan digit terakhir yang telah ditambahkan
	} else { // Jika tidak perlu dilakukan carry
		digits[len(digits)-1] = last_digit // Mengubah digit terakhir
		return digits                      // Mengembalikan array
	}
}

// Nomor 7 :
// Fungsi ini membandingkan dua array dan mengembalikan elemen yang sama dan berbeda
func compareArrays(arr1, arr2 []string) ([]string, []string) {
	same := []string{}      // Inisialisasi slice untuk elemen yang sama
	different := []string{} // Inisialisasi slice untuk elemen yang berbeda

	// Membuat map untuk menyimpan elemen-elemen dalam arr1
	arr1Map := make(map[string]bool)
	for _, elem := range arr1 {
		arr1Map[elem] = true // Menambahkan elemen ke map
	}

	// Membandingkan elemen-elemen dalam arr2 dengan arr1Map
	for _, elem := range arr2 {
		if arr1Map[elem] { // Jika elemen ada dalam arr1Map, menambahkan ke slice 'same'
			same = append(same, elem)
		} else { // Jika elemen tidak ada dalam arr1Map, menambahkan ke slice 'different'
			different = append(different, elem)
		}
	}

	return same, different // Mengembalikan slice 'same' dan 'different'
}

// Nomor 8 :
// Fungsi ini mengecek apakah array adalah palindrome atau tidak
func isPalindrome(arr []string) bool {
	length := len(arr) // Mendapatkan panjang array

	// Iterasi melalui setengah panjang array
	for i := 0; i < length/2; i++ {
		// Membandingkan elemen-elemen pada posisi yang berlawanan
		if arr[i] != arr[length-1-i] { // Jika elemen tidak sama, mengembalikan false
			return false
		}
	}

	return true // Jika semua elemen cocok dengan elemen pada posisi yang berlawanan, mengembalikan true
}

// Nomor 9 :
// Fungsi ini mencari elemen minimum dan maksimum dalam array dan mengembalikannya
func minMaxArray(arr []int) (int, int) {
	if len(arr) == 0 { // Jika array kosong, mengembalikan 0 untuk nilai minimum dan maksimum
		return 0, 0
	}

	min := arr[0] // Mendeklarasikan dan menginisialisasi min dengan elemen pertama dari arr
	max := arr[0] // Mendeklarasikan dan menginisialisasi max dengan elemen pertama dari arr

	// Iterasi melalui setiap elemen dalam array
	for _, num := range arr {
		if num < min { // Jika elemen lebih kecil dari min, mengubah min menjadi elemen tersebut
			min = num
		}
		if num > max { // Jika elemen lebih besar dari max, mengubah max menjadi elemen tersebut
			max = num
		}
	}
	return min, max // Mengembalikan nilai min dan max
}

// Nomor 10 :
// Fungsi ini menambahkan elemen target ke array dan mengurutkannya
func insertElement(arr []int, target int) []int {
	arr = append(arr, target) // Menambahkan elemen target ke array

	sort.Ints(arr) // Mengurutkan array

	return arr // Mengembalikan array yang sudah diurutkan
}

func main() {
	// Di sini Anda dapat memanggil fungsi-fungsi yang telah Anda definisikan di atas
	// Misalnya, Anda bisa memanggil fungsi printPattern() seperti ini:
	// printPattern()
}
