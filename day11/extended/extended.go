package main

import "fmt"

type person struct {
	name string
	age  int
}

// immutable : tidak bisa diubah data originnya
func immutable(i int, s string, p person) {
	i = i * 5
	s = "bootcamp go"
	p.name = "dendy"
}

//mutable, copy by pointer
func mutable(i *int, s *string, p *person) {
	*i = *i * 5
	*s = "bootcamp go"
	p.name = "dendy"
}

func monMap(m map[int]string) {
	m[2] = "second"
	m[3] = "third"
	delete(m, 1)
}

func main() {
	//sample immutable, copy value
	p := person{}
	i := 10
	s := "java"
	immutable(i, s, p)
	fmt.Println(i, s, p)

	//call mutable
	mutable(&i, &s, &p)
	fmt.Println(i, s, p)

	//modifikasi map
	m := map[int]string{
		1: "first",
		2: "dua",
	}
	monMap(m)
	fmt.Println(m)

}
