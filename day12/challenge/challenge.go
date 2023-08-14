package main

import (
	"fmt"
	"math/rand"
)

const (
	lembur = 500_000
	makan  = 10_000
)

// 1. menyimpan data employees
type Employees struct {
	EmpId                     int
	FirstName, LastName, Role string
	BasicSalary               float64
	tunjangan                 Allowances
	totalSalary               float64
}

type Allowances struct {
	sales      TunjanganSales
	programmer TunjanganProgrammer
	qa         TunjanganQA
}

type TunjanganSales struct {
	bonus, Commission float64
}

type TunjanganProgrammer struct {
	lembur float64
}

type TunjanganQA struct {
	makan float64
}

// hold data employee in slice
type EmployeeList []Employees

// 2. create constructor
func NewEmployee(firstName, lastName, role string, basicSalary float64, allowances Allowances) *Employees {
	return &Employees{
		rand.Intn(10),
		firstName, lastName, role, basicSalary, allowances, setTotalSalary(basicSalary, allowances)}
}

// 3. Create InitEmployees
func initEmployee() EmployeeList {
	emp1 := NewEmployee("Anton", "Pratama", "sales", 6_000_000,
		caclTunjangan("sales", Allowances{
			sales:      TunjanganSales{bonus: 350_000, Commission: 50_000},
			programmer: TunjanganProgrammer{},
			qa:         TunjanganQA{},
		}))

	emp2 := NewEmployee("Budi", "Doremi", "Programmer", 6_000_000,
		caclTunjangan("programmer", Allowances{
			sales:      TunjanganSales{},
			programmer: TunjanganProgrammer{lembur},
			qa:         TunjanganQA{},
		}))

	return EmployeeList{*emp1, *emp2}
}

// function caclTunjangan
func caclTunjangan(role string, allowc Allowances) Allowances {
	tunjangan := Allowances{}
	switch role {
	case "sales":
		tunjangan.sales.bonus = allowc.sales.bonus
		tunjangan.sales.Commission = allowc.sales.bonus
	case "programmer":
		tunjangan.programmer.lembur = lembur
	case "QA":
		tunjangan.qa.makan = makan
	}
	return tunjangan
}

func setTotalSalary(basicSalary float64, allowc Allowances) float64 {
	return basicSalary + allowc.programmer.lembur + allowc.qa.makan + allowc.sales.bonus + allowc.sales.Commission
}

// 4. create interface
type summary interface {
	setBasicSalary(salary float64) float64
}

func (emp *Employees) setBasicSalary(salary float64) {
	emp.BasicSalary = salary
	emp.totalSalary = emp.BasicSalary + emp.tunjangan.programmer.lembur + emp.tunjangan.sales.bonus + emp.tunjangan.sales.Commission + emp.tunjangan.qa.makan
}

func main() {
	EmployeeList := initEmployee()

	EmployeeList[1].setBasicSalary(7_000_000)

	for _, v := range EmployeeList {
		fmt.Println(v)
	}

}
