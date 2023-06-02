package main

import (
	"fmt"
)

func main() {
	println("Nomor 1 -------------")
	arrayDiagonalNo1()
	fmt.Println()
	println("Nomor 2 -------------")
	arrayDiagonalNo2()
	fmt.Println()
	println("Nomor 3 -------------")
	nomor03()
	fmt.Println()
	println("Nomor 4 -------------")
	nomor04()
	fmt.Println()
	println("Nomor 5 -------------")
	nomor05()
	fmt.Println()
	println("Nomor 6 -------------")
	nomor06()
	fmt.Println()
	println("Nomor 7 -------------")
	nomor07()
}

func arrayDiagonalNo1() {
	var matrix [5][5]int
	simpanData := make([]int, 5)
	count := 0
	for i := 0; i < len(matrix); i++ {
		count += i + 1
		simpanData[i] = i + 1
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = simpanData[col]
			} else if col > row {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 10
			}
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Printf("Total Sum : %d", count)
	println()
}

func arrayDiagonalNo2() {
	var matrix [5][5]int
	simpanData := make([]int, len(matrix))
	count := 0
	simpanData[0] = 2
	//untuk membuat nilai diagonal 2,4,8,16,32
	for i := 1; i < len(matrix); i++ {
		simpanData[i] = simpanData[i-1] * 2
	}
	//untuk menghitung jumlah array simpan data
	for _, value := range simpanData {
		count += value
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = simpanData[col]
			} else if col > row {
				matrix[row][col] = 88
			} else {
				matrix[row][col] = 22
			}
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Printf("Total Sum : %d", count)
	println()
}

func nomor03() {

	var data = [9][11]string{
		{"Student-0", "A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-1", "D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"Student-2", "E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"Student-3", "C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"Student-4", "A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-5", "B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-6", "B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-7", "E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Kunci Jawaban", "D", "B", "D", "C", "C", "D", "A", "E", "A", "D"},
	}
	perbandingan := len(data) - 1
	count := 0
	var simpanDataStudent = make([]string, perbandingan)
	//for untuk student
	for i := 0; i < perbandingan; i++ {
		simpanDataStudent[i] = data[i][0]
	}
	// fmt.Println(simpanDataStudent)
	for row := 0; row < perbandingan; row++ {
		for col := 1; col < len(data[row]); col++ {
			if data[row][col] == data[perbandingan][col] {
				count += 10
			}
		}

		fmt.Printf(simpanDataStudent[row]+": %d point", count)
		fmt.Println()
		count = 0
	}

	// fmt.Printf(student)
	// for row := 0; row < perbandingan; row++ {
	// 	for col := 1; col < len(data); col++ {
	// 		if data[row][col] == data[perbandingan][col] {
	// 			count = count + 10
	// 		}else{
	// 			continue
	// 		}
	// 	}
	// 	fmt.Println()
	// }

}

func nomor04() {
	var matrix [7][7]int
	simpanDiagonal := make([]int, len(matrix))
	simpanRight := make([]int, len(matrix))
	simpanBottom := make([]int, len(matrix))
	varInti := len(matrix) - 1

	simpanDiagonal[0] = 2
	increment := 3

	//Nilai array untuk diagonal
	for i := 1; i < len(matrix); i++ {
		if i == varInti/2 {
			simpanDiagonal[i] = increment
			simpanDiagonal[i+1] = simpanDiagonal[i-1] + simpanDiagonal[i]
			i++
		} else if i%3 == 0 || i%3 != 0 {
			simpanDiagonal[i] = simpanDiagonal[i-1] + increment
		}
	}
	// fmt.Printf("%d ", simpanDiagonal)
	//Nilai Array untuk right
	simpanRight[0] = simpanDiagonal[varInti]
	simpanRight[1] = simpanDiagonal[3] * increment
	for i := 2; i < len(matrix); i++ {
		if simpanRight[0] != 0 {
			if i == 2 {
				simpanRight[i] = simpanDiagonal[varInti] + increment
			} else if i == 5 {
				simpanRight[i] = simpanRight[1] * increment
				simpanRight[i+1] = simpanRight[i-1] + increment
				i++
			} else {
				simpanRight[i] = simpanRight[i-1] + increment
			}
		}
	}
	// fmt.Printf("%d ", simpanRight)
	//Nilai Array untuk Bottom
	simpanBottom[0] = 2
	simpanBottom[1] = 41
	simpanBottom[varInti] = simpanRight[varInti]
	for i := 2; i < len(matrix)-1; i++ {
		if simpanBottom[0] != 0 {
			if i == 3 {
				simpanBottom[i] = simpanRight[varInti-1] * increment
				simpanBottom[i+1] = simpanBottom[i-1] - increment
				i++
			} else {
				simpanBottom[i] = simpanBottom[i-1] - increment
			}
		}
	}
	// fmt.Printf("%d ", simpanBottom)

	// Berikan Metrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row+col == varInti {
				matrix[row][col] = simpanDiagonal[col]
			} else if col == varInti {
				matrix[row][col] = simpanRight[row]
			} else if row == varInti {
				matrix[row][col] = simpanBottom[col]
			} else {
				matrix[row][col] = 0
			}
			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}
	// fmt.Println(simpanBottom)
}

func nomor05() {
	var matrix [7][7]int
	var simpanDataTop = make([]int, len(matrix))
	var simpanDataLeft = make([]int, len(matrix))
	var simpanDataRight = make([]int, len(matrix))
	var simpanDataBottom = make([]int, len(matrix))

	valueIndeksAkhir := len(matrix) - 1
	valueIndeksAwal := 0
	increment := 3

	//nilai array atas
	simpanDataTop[0] = 2
	for i := 1; i < len(matrix); i++ {
		if i == valueIndeksAkhir/2 {
			simpanDataTop[i] = increment
			simpanDataTop[i+1] = simpanDataTop[i-1] + simpanDataTop[i]
			i++
		} else {
			simpanDataTop[i] = simpanDataTop[i-1] + increment
		}
	}
	// fmt.Println(simpanDataTop)

	//nilai array kanan terusan nilai array atas
	simpanDataRight[0] = simpanDataTop[valueIndeksAkhir]

	for i := 1; i < len(matrix); i++ {
		if i == 1 || i == valueIndeksAkhir-1 {
			simpanDataRight[i] = increment
			simpanDataRight[i+1] = simpanDataRight[i-1] + simpanDataRight[i]
			i++
		} else {
			simpanDataRight[i] = simpanDataRight[i-1] + increment
		}
	}
	// fmt.Println(simpanDataRight)

	//Nilai array bawah terusan nilai array kanan
	simpanDataBottom[valueIndeksAkhir] = simpanDataRight[valueIndeksAkhir]
	for i := valueIndeksAkhir - 1; i >= 0; i-- {
		if i == valueIndeksAkhir/2 {
			simpanDataBottom[i] = increment
			simpanDataBottom[i-1] = simpanDataBottom[i+1] + simpanDataBottom[i]
			i--
		} else {
			simpanDataBottom[i] = simpanDataBottom[i+1] + increment
		}
	}
	// fmt.Println(simpanDataBottom)

	//nilai array kiri
	simpanDataLeft[valueIndeksAwal] = simpanDataTop[valueIndeksAwal]
	simpanDataLeft[valueIndeksAkhir] = simpanDataBottom[valueIndeksAwal]
	for i := valueIndeksAkhir - 1; i > 0; i-- {
		if i == 1 || i == valueIndeksAkhir-1 {
			simpanDataLeft[i] = increment
			if i == valueIndeksAkhir-1 {
				simpanDataLeft[i-1] = simpanDataLeft[i+1] + simpanDataLeft[i]
				i--
			}
		} else {
			simpanDataLeft[i] = simpanDataLeft[i+1] + increment
		}
	}
	// fmt.Println(simpanDataLeft)

	//Masukkan ke dalam matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == valueIndeksAwal {
				matrix[row][col] = simpanDataTop[col]
			} else if col == valueIndeksAkhir {
				matrix[row][col] = simpanDataRight[row]
			} else if col == valueIndeksAwal {
				matrix[row][col] = simpanDataLeft[row]
			} else if row == valueIndeksAkhir {
				matrix[row][col] = simpanDataBottom[col]
			}
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}

}

func nomor06() {
	var matrix [7][7]int
	var simpanDataPertama = make([]int, len(matrix))
	var simpanDataKedua = make([]int, len(matrix))
	valueIndeksAkhir := len(matrix) - 1

	simpanValueAwal := 0

	//nilai array pertama
	for i := 0; i <= valueIndeksAkhir; i++ {
		simpanDataPertama[i] = i
	}
	// fmt.Println(simpanDataPertama)

	//nilai array kedua
	for i := 0; i <= valueIndeksAkhir; i++ {
		simpanDataKedua[i] = i + 6
	}
	// fmt.Println(simpanDataKedua)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == simpanValueAwal {
				matrix[row][col] = simpanDataPertama[col]
			} else if col == simpanValueAwal {
				matrix[row][col] = simpanDataPertama[row]
			} else if col == valueIndeksAkhir {
				matrix[row][col] = simpanDataKedua[row]
			} else if row == valueIndeksAkhir {
				matrix[row][col] = simpanDataKedua[col]
			} else {
				matrix[row][col] = 0
			}
			fmt.Printf("%d ", matrix[row][col])
		}
		println()
	}
}

func nomor07() {
	// matrix := make([][]int, n+1)

	var matrix [7][8]int
	var simpanValue = make([]int, len(matrix))
	simpanValueAwal := 0
	simpanValueDiagon := 0
	hitungValue := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)+1; col++ {
			matrix[row][col] = simpanValueAwal
			simpanValueAwal++
			if col == len(matrix) {
				matrix[row][col] = hitungValue
				simpanValue[row] = hitungValue
			}

			hitungValue += matrix[row][col]
			if row == col {
				simpanValueDiagon += matrix[row][col]
			}
		}
		simpanValueAwal = row + 1
		hitungValue = 0
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)+1; col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println(simpanValue, simpanValueDiagon)
}
