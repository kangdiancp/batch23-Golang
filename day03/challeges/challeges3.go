package main

import "fmt"

func s1(count int) {
	for row := 0; row < count; row++ {
		for col := 0; row < count; col++ {
			if row == count-1-col {
				fmt.Print("* ")
			} else {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}

}

func s2(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if j >= count-i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func s3(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print("20 ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()

	}
}

func s4(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print("	")
			} else {
				fmt.Print("21 ")
			}
		}
		fmt.Println()
	}
}

func s5(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col+row == count-1 {
				fmt.Printf("10")
			} else if col+row > count-1 {
				fmt.Printf("20")
			} else {
				fmt.Printf("21")
			}

		}
		fmt.Println()
	}
}

func s6(count int) {
	for row := 1; row <= count; row++ {
		if row%2 == 0 {
			for col := 1; col <= count; col++ {
				fmt.Print(col)
			}
		} else {
			for col := count; col >= 1; col-- {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func s7(count int) {
	for row := 1; row <= count; row++ {
		for ulang := 0; ulang < count/2; ulang++ {

			for col := 1; col <= count; col++ {
				if col == count {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}

			for col := count; col >= 1; col-- {
				if col == 1 {
					fmt.Print(count - row + 1)
				} else {
					fmt.Print("")
				}
			}

		}

		if count%2 == 1 {
			for col := 1; col <= count; col++ {
				if col == count {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
		}
		fmt.Println()
	}
}

func s8(count int) {
	for row := 0; row < count; row++ {
		if row%2 == 0 {
			for col := 1; col <= count; col++ {
				if col%2 == 0 {
					print(col)
				} else {
					print("-")
				}
			}
		} else {
			for col := 1; col <= count; col++ {
				if col%2 == 0 {
					print("-")
				} else {
					print(col)
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	//s1(5)
	//s2(5)
	//s3(5)
	//s4(5)
	//s5(5)
	//s6(5)
	//s7()
	s8(5)

}
