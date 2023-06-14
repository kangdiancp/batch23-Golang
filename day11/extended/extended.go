package main

import "fmt"

type Person struct {
	name string
	age  int
}

// immutable : tidak bisa diubah data origin-nya
func immutable(i int, s string, p Person) {
	i = i * 5
	s = "Bootcamp Go"
	p.name = "Dendy"
}

//mutable, copy by pointer
func mutable(i *int, s *string, p *Person) {
	*i = *i * 5
	*s = "Bootcamp Go"
	p.name = "Dendy"
}

func modMap(m map[int]string) {
	m[2] = "Second"
	m[3] = "Third"
	delete(m, 1)
}

func main() {
	//sample immutable, copy value
	p := Person{}
	i := 10
	s := "Java"
	immutable(i, s, p)
	fmt.Println(i, s, p)

	//call mutable
	mutable(&i, &s, &p)
	fmt.Println(i, s, p)

	// modifikasi map
	m := map[int]string{
		1: "First",
		2: "Dua",
	}
	modMap(m)
	fmt.Println(m)
}
