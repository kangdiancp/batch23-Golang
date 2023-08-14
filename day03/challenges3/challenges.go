package main

import (
	"fmt"
)

func c1(count int) {
	for col := 0; col < count; col++ {
		for row := 0; row < count; row++ {
			if row == count-1-col {
				fmt.Print("* ")
			} else {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}

}

func c2(count int) {

	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col >= count-row-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func c3(count int) {
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

func c4(count int) {
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

func c5(count int) {
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

func q1(count int) {
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

func q2(count int) {
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

func q3(count int) {
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
	c1(5)
	// c2(5)
	//c3(5)
	//c4(5)
	//c5(5)
	//q1(5)
	//q2(5)
	//q3(5)

}
