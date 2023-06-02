package main

import (
	"fmt"
	"strconv"
)

func quiz1SumDiagonal(){
	var matrix [5][5]int
	sumdiagoal:=0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = row + 1
				sumdiagoal += matrix[row][col]
			} else if row > col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d ", matrix[row][Colom])
		}
		fmt.Println()
	}
	fmt.Print("   Total Sum : ", sumdiagoal)
}

func quiz2JumlahDiagnal2(){
	var matrix [5][5]int
	ndiagonal := make([]int, len(matrix))
	ndiagonal[0] = 2
	sumdiagoal:=0

	for i := 1; i < len(matrix); i++ {
		ndiagonal[i] = ndiagonal[i-1]*2
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = ndiagonal[row]
				sumdiagoal += matrix[row][col]
			} else if row > col {
				matrix[row][col] = 22
			} else {
				matrix[row][col] = 88
			}
		}
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d ", matrix[row][Colom])
		}
		fmt.Println()
	}
	fmt.Print("   Total Sum : ", sumdiagoal)
}

func quiz3PointSiswa(){
	matrix := [9][11] string{
		{"student-0", "A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-1", "D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"student-2", "E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"student-3", "C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"student-4", "A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"student-5", "B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"student-6", "B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-7", "E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"kunci jawaban", "D", "B", "D", "C", "C", "D", "A", "E", "A", "D"},
	}

	nilai := 0
	compare := len(matrix)-1

	for row := 0; row < compare; row++ {
		for col := 1; col < len(matrix[row]); col++ {
			if  matrix[row][col] == matrix[compare][col]{
				nilai += 10
			}
		}
		fmt.Printf(matrix[row][0] + " : %d Point", nilai)
		fmt.Println()
		nilai = 0
	}
	
}

func quiz4MatrixDeretBilangan(){
	var  matrix [7][7] int
	m := 3
	vary := m*3
	n := len(matrix)-1
	p := 17
	hitung := p + ((len(matrix)+1) * 3)
	matrix[6][0] = 2
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {

			if row == n/2 && col == n/2{
				matrix[row][col] = m
			}else if col + row == n && col < n/2 && col > 0{
				matrix[row][col] = (3 * (col + 1)) -1
			}else if col + row == n && col > n/2{
				matrix[row][col] = (3 * col) -1
			 }else if (col == n && row > 1) && row != len(matrix)-2 && row < len(matrix)-2 {
			 	matrix[row][col] = (3 * (col + row -1)) - 1
			 }else if col == n && row > len(matrix)-2{
				matrix[row][col] = (3 * (col + row - 1)) - 4
			 }else if col == n && row == 1 {
				matrix[row][col] = vary
			 }else if  row == len(matrix)-2 && col == n  {
				matrix[row][col] = vary * m
			 }else if row == n && col > 0 && col != n/2 && col != n {
				matrix[row][col] = hitung
				hitung -= m
			 }else if row == n && col == n/2{
				matrix[row][col] = vary * 9
			 }
		}
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d  ", matrix[row][Colom])
		}
		fmt.Println()
	}
}



func quiz5AngkaTepi(n, o, m int) {

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	rowStart, rowEnd := 0, n-1
	colStart, colEnd := 0, n-1
	//direction := 0
	for o <= n*n {
		// Menambahkan elemen pada baris atas
		for i := colStart; i <= colEnd; i++ {

			matrix[rowStart][i] = o

			if (i+1)%(m+1) == 0 {
				matrix[rowStart][i] = m
				o -= m

			}

			o += m
		}
		sisa := (len(matrix)) % 4

		rowStart++

		// Menambahkan elemen pada kolom kanan
		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = o
			if (i+sisa)%(m+1) == 0 {
				matrix[i][colEnd] = m
				o -= m
			}

			o += m
		}
		colEnd--
		var sisa2 int
		if sisa == 0 {
			sisa2 = 3

		} else if sisa == 1 {
			sisa2 = 1

		} else if sisa == 2 {
			sisa2 = 3

		} else if sisa == 3 {
			sisa2 = 1

		}

		// Menambahkan elemen pada baris bawah
		for i := colEnd; i >= colStart; i-- {
			matrix[rowEnd][i] = o

			if (i+sisa)%(m+1) == 2 {

				matrix[rowEnd][i] = m
				o -= m
			}
			o += m

		}
		rowEnd--
		var sisa3 int
		if sisa2 == 0 {
			sisa3 = 3

		} else if sisa2 == 1 {
			sisa3 = 1

		} else if sisa2 == 2 {
			sisa3 = 3

		} else if sisa2 == 3 {
			sisa3 = 1

		}

		// Menambahkan elemen pada kolom kiri
		for i := rowEnd; i >= rowStart; i-- {
			matrix[i][colStart] = o
			if (i+sisa3)%(m+1) == 2 {

				matrix[i][colStart] = m
				o -= m
			}
			o += m
		}
		colStart++
	}

	//print matrix
	stringMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		stringMatrix[i] = make([]string, len(matrix[i]))
		for j := 0; j < len(matrix); j++ {
			stringMatrix[i][j] = strconv.Itoa(matrix[i][j])
			if stringMatrix[i][j] == "0" {

				stringMatrix[i][j] = " "

			}
		}
	}

	for i := 0; i < len(stringMatrix); i++ {
		for j := 0; j < len(stringMatrix); j++ {
			fmt.Print(stringMatrix[i][j], "	")
		}
		fmt.Println()
	}
}

func quiz6BarisanAngka(){
	var  matrix [7][7] int
	cout := 0
	cout2 := 0
	long := len(matrix)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == 0 {
				matrix[row][col] = cout
				cout += 1
			}else if col == len(matrix)-1 && row != 0{
				matrix[row][col] = long
				long += 1
			}else if col == 0 && row != 0{
				matrix[row][col] = cout2 + 1
				cout2 += 1
			}else if (col > 0 && col < len(matrix)-1) && row == len(matrix)-1{
				matrix[row][col] = cout2 + 1
				cout2 += 1
			}
		}		
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d  ", matrix[row][Colom])
		}
		fmt.Println()
	}
}

func quiz7SumElementMatrix(){
	var matrix [7][8]int
	count := 0
	sumdiagoal:= 0
	jmlrow := 0
	sumrow := make([]int, len(matrix))
	// panjang := len(matrix)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix)+1; col++ {
			matrix[row][col] = count
			count++
			if col == len(matrix){
				matrix[row][col] = jmlrow
				sumrow[row] = jmlrow
			}
			jmlrow += matrix[row][col]
			if row == col  {
				sumdiagoal += matrix[row][col]	
			}
		}
		count = row + 1
		jmlrow = 0
	}

	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix)+1; Colom++ {
			fmt.Printf("%d ", matrix[row][Colom])
		}
		fmt.Println()
	}
	fmt.Println(sumrow, sumdiagoal)
}

func displayMatrix1(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d ", matrix[row][Colom])
		}
		fmt.Println()
	}
}

func main() {
	quiz1SumDiagonal()
	quiz2JumlahDiagnal2()
	quiz3PointSiswa()
	quiz4MatrixDeretBilangan()
	quiz5AngkaTepi(7, 2, 3)
	quiz6BarisanAngka()
	quiz7SumElementMatrix()
}