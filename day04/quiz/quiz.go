package main

import (
	"fmt"
	"strconv"
)

func quiz1evenOrOddNumber(n int) {
	OddNum := []int{}

	for i := 0; i < n; i++ {
		count := (i * 2) + 1
		OddNum = append(OddNum, count)
	}
	fmt.Println(OddNum)
}

func quiz2PowNumber(n int){
	Pownum := make([]int, n)
	Pownum[0] = 1

	for i := 1; i < n; i++ {
		Pownum[i] = Pownum[i-1] * 2
	}
	fmt.Print(Pownum)
}

func quiz3Fibonnaci(n int){
	fibo := make([]int, n)
	fibo[0] = 2
	fibo[1] = 3
	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-2] + fibo[i-1]
	}
	fmt.Println(fibo)
}

func quiz4ReverseNumber(arr []int)[]int{
	for i := 1; i < len(arr)/2; i++ {
		rever := arr[len(arr)-i]
		arr[len(arr)-i] = arr[i-1]
		arr[i-1] = rever
	}
	return arr
}

func IsPrime(n int)bool{
	for i := 2; i <= n/2; i++ {
		if n%i == 0{
			return false
		}
	}
	return true
}

func quiz5IsPrimeNumber(n int)[]string{
	prime := make([]string, n+1) 
	for i := 1; i <= n; i++ {  
		if IsPrime(i){
			prime[i] = strconv.Itoa(i)+" = true"
		}else{
			prime[i] = strconv.Itoa(i)+" = false"
		}
	}
	return prime
}


func main() {
	//quiz1evenOrOddNumber(5)
	//quiz2PowNumber(6)
	//quiz3Fibonnaci(6)
	//fmt.Println(quiz4ReverseNumber([]int {4, 3, 6, 7, 8, 1}))
	fmt.Println(quiz5IsPrimeNumber(4))
	//fmt.Println(IsPrime(7))
}