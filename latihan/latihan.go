package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// 1
func countNegatives(grid [][]int) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				count = count + 1
			}
		}
	}
	return count
}

// 2 Two Sum
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 9 {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 3 Palindrome Number
func isPalindrome(num int) bool {
	strNum := strconv.Itoa(num) // Mengubah angka menjadi string

	// Memeriksa apakah angka negatif
	if num < 0 {
		strNum = strNum[1:] // Menghilangkan tanda negatif dari string
	}

	arr := make([]int, len(strNum))

	for i, char := range strNum {
		digit, _ := strconv.Atoi(string(char)) // Mengubah karakter menjadi integer
		arr[i] = digit
	}

	// Menyesuaikan dengan tanda negatif
	if num < 0 {
		arr[0] *= -1
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}

	}
	if arr[0] < 0 {
		return false
	}

	return true
}

// 4 PlusOne
func plusOneDigit(digits []int) []int {
	n := len(digits)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

// 5 Search Insert Position
func searchInsert(nums []int, target int) int {

	nums = append(nums, target)
	sort.Ints(nums)
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return index
}

// 6  Length of Last Word
func lengthOfLastWord(s string) int {
	var lenght int
	WordSplit := strings.Fields(s)
	lenght = len(WordSplit)
	panjang := WordSplit[lenght-1]
	panjanga := len(panjang)

	return panjanga
}

// 7 Roman to Integer
func romanToInt(s string) int {

	s = strings.ToUpper(s)

	values := make([]int, len(s))

	for i, c := range s {
		switch c {
		case 'I':
			values[i] = 1
		case 'V':
			values[i] = 5
		case 'X':
			values[i] = 10
		case 'L':
			values[i] = 50
		case 'C':
			values[i] = 100
		case 'D':
			values[i] = 500
		case 'M':
			values[i] = 1000
		}
	}

	// Menghitung nilai desimal dengan menjumlahkan atau mengurangi nilai setiap karakter Romawi
	result := values[len(values)-1]
	for i := len(values) - 2; i >= 0; i-- {
		if values[i] < values[i+1] {
			result -= values[i]
		} else {
			result += values[i]
		}
	}

	return result
}

// 8 Valid Palindrome
func isPalindrome2(s string) bool {
	var sb strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			sb.WriteRune(unicode.ToLower(c))
		}
	}

	// Check if the string is a palindrome
	runes := []rune(sb.String())
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true

}

// 9 Longest Common Prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for !startsWith(strs[i], prefix) {

			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
func startsWith(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}

//10 Valid Parentheses

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			// Push opening brackets onto the stack
			stack = append(stack, c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) == 0 {
				// If the stack is empty and we encounter a closing bracket, it's invalid
				return false
			}

			// Get the top element from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if the current closing bracket matches the top opening bracket
			if (c == ')' && top != '(') || (c == '}' && top != '{') || (c == ']' && top != '[') {
				return false
			}
		}
	}

	// If there are still opening brackets left in the stack, it's invalid
	return len(stack) == 0
}

// 11 remove duplicate
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Variable to keep track of the number of unique elements

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			// Found a new unique element
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// 12 remove element
func removeElement(nums []int, val int) int {
	k := 0 // Variable to keep track of the number of elements not equal to val

	for _, num := range nums {
		if num != val {
			// Found an element not equal to val
			nums[k] = num
			k++
		}
	}

	return k
}

// 13 Find Smallest Letter Greater Than Target
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left, right := 0, n-1

	// If the target is greater than or equal to the last element,
	// the answer is the first element
	if target >= letters[right] {
		return letters[0]
	}

	// Binary search to find the smallest character greater than target
	for left <= right {
		mid := left + (right-left)/2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return letters[left]
}

//14 Merge Sorted Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1     // Indeks terakhir elemen aktif dalam nums1
	j := n - 1     // Indeks terakhir elemen dalam nums2
	k := m + n - 1 // Indeks terakhir dalam hasil penggabungan

	// Mulai dari belakang dan pindahkan elemen terbesar dari kedua array ke nums1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// Jika masih ada elemen tersisa di nums2, salin ke nums1
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 15 Pascal's Triangle
func generatePascalsTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

// 16 Pascal's Triangle II
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}

	return row
}

// 17 SIngle Number
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 18 Climbing Stairs
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// 19 Excel Sheet Column Title
func convertToTitle(columnNumber int) string {
	var result []byte

	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		result = append([]byte{byte('A' + remainder)}, result...)
		columnNumber /= 26
	}

	return string(result)
}

// 20  Excel Sheet Column Number
func titleToNumber(columnTitle string) int {
	result := 0
	for _, ch := range columnTitle {
		result = result*26 + int(ch-'A'+1)
	}
	return result
}

// 21 Power of Two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// 22 Constain Duplicat
func isDuplicat(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

// 23 Intersection of Two Arrays
func twoSame(nums1 []int, nums2 []int) []int {
	var arr []int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2)-1; j++ {
			if nums1[i] == nums2[j] {
				arr[i] = nums1[i]
			}
		}
	}
	return arr
}

//24 Robot Return to Origin

func isOrigin(moves string) bool {
	lenght := len(moves)

	var n int
	for i := 0; i < lenght; i++ {
		switch {
		case moves[i] == 'U':
			n++
		case moves[i] == 'D':
			n--
		case moves[i] == 'R':
			n++
		case moves[i] == 'L':
			n--
		}
	}
	if n == 0 {
		return true

	}
	return false
}

//23

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(twoSame(nums1, nums2))
}
