package main

import "fmt"

func main() {

	main1(2005)
	main2()

}

func main1(tahunlahir int) {

	var currentYear = 2021

	if age := currentYear - tahunlahir; age < 17 {
		fmt.Println("kamu belum boleh membuat sim")
	} else {
		fmt.Println("kamu sudah boleh membuat sim")
	}
}

func main2() {

}
