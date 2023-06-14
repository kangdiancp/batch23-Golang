package main

import (
	"fmt"
	"math/rand"
)

const (
	LEMBUR = 500_000
	MAKAN  = 10_000
)

// 1. create struct
type Employee struct {
	empId                     int
	firstName, lastName, role string
	basicSalary               float64
	tunjangan                 Allowances
	totalSalary               float64
	records                   []int
}

type Allowances struct {
	sales      TunjSales
	programmer TunjProgrammer
	qa         TunjQA
}

type TunjSales struct {
	bonus, commision float64
}

type TunjProgrammer struct {
	lembur float64
}

type TunjQA struct {
	makan float64
}

// hold data employee in slice
type EmployeeList []Employee

// 2. create constuctor
func newEmployee(firstName, lastName, role string,
	basicSalary float64, allowences Allowances) *Employee {

	return &Employee{
		rand.Intn(10),
		firstName, lastName, role,
		basicSalary, allowences,
		setTotalSalary(basicSalary, allowences),
		[]int{1, 2, 3, 4}}
}

// 3. create initEmployee
func initEmployee() EmployeeList {

	return EmployeeList{
		*newEmployee("budi", "juna", "Sales", 4_500_000,
			calcTunjangan("Sales", Allowances{
				sales:      TunjSales{bonus: 350_000, commision: 50_000},
				programmer: TunjProgrammer{},
				qa:         TunjQA{},
			})),

		*newEmployee("dwi", "runa", "Programmer", 6_000_000,
			calcTunjangan("Programmer", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgrammer{},
				qa:         TunjQA{},
			})),

		*newEmployee("dwi", "runa", "Programmer", 6_000_000,
			calcTunjangan("Programmer", Allowances{
				sales:      TunjSales{},
				programmer: TunjProgrammer{},
				qa:         TunjQA{},
			})),
	}

	//return EmployeeList{*emp1, *emp2}
}

// function calcTunjangan
func calcTunjangan(role string, allowc Allowances) Allowances {
	tunjangan := Allowances{}
	switch role {
	case "Sales":
		tunjangan.sales.bonus = allowc.sales.bonus
		tunjangan.sales.commision = allowc.sales.commision
	case "Programmer":
		tunjangan.programmer.lembur = LEMBUR
	case "QA":
		tunjangan.qa.makan = MAKAN
	}
	return tunjangan
}

func setTotalSalary(basicSalary float64, allowc Allowances) float64 {
	return basicSalary + allowc.programmer.lembur + allowc.sales.bonus +
		allowc.sales.commision + allowc.qa.makan
}

// 4. create interface
type summary interface {
	setBasicSalary(salary float64)
}

func (emp *Employee) setBasicSalary(salary float64) {
	emp.basicSalary = salary
	emp.totalSalary = emp.basicSalary + emp.tunjangan.programmer.lembur +
		emp.tunjangan.sales.bonus + emp.tunjangan.sales.commision +
		emp.tunjangan.qa.makan

}

func main() {
	employeeList := initEmployee()

	employeeList[1].setBasicSalary(7_000_000)

	for _, v := range employeeList {
		fmt.Println(v)
	}
}
