package main

import "fmt"

func soalLooping1(count int) {
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col + row == count - 1 {
				fmt.Printf(" * ")
			}else{
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
}

func soalLooping2(count int){
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col + row < count - 1 {
				fmt.Printf("   ")
			}else{
				fmt.Printf(" * ")
			}
		}
		fmt.Println()
	}
}

func soalLooping3(count int){
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col + row < count - 1 {
				fmt.Printf("   ")
			}else if col + row == count - 1{
				fmt.Printf(" 10")
			}else{
				fmt.Printf(" 20")
			}
		}
		fmt.Println()
	}
}

func soalLooping4(count int){
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col + row > count - 1 {
				fmt.Printf("   ")
			}else if col + row == count - 1{
				fmt.Printf(" 10")
			}else{
				fmt.Printf(" 21")
			}
		}
		fmt.Println()
	}
}

func soalLooping5(count int){
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col + row > count - 1 {
				fmt.Printf(" 20")
			}else if col + row == count - 1{
				fmt.Printf(" 10")
			}else{
				fmt.Printf(" 21")
			}
		}
		fmt.Println()
	}
}

func soalLooping6(count int){
	for row := 1; row <= count; row++ {
		if row % 2 == 0{
			for col := 1; col <= count; col++ {
				fmt.Print(col)
			}
		}else {
			for col := count; col >= 1; col-- {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func soalLooping7(count int){
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col%2 == 0 {
				fmt.Print(row + 1)
			} else {
				fmt.Print(count - row)
			}
		}
		fmt.Println()
	}

}

func soalLooping8(count int){
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row%2 == 1 && col%2 == 1{
				fmt.Print("-")
			}else if row%2 == 0 && col%2 == 0{
				fmt.Print("-")
			}else{
				fmt.Print(col + 1)
			}
		}
		fmt.Println()
	}
}

func main() {
	//soalLooping1(5)
	//soalLooping2(5)
	//soalLooping3(5)
	//soalLooping4(5)
	//soalLooping5(5)
	//soalLooping6(9)
	//soalLooping7(9)
	soalLooping8(9)
}