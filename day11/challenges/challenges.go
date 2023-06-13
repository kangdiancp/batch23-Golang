package main

import "fmt"

type Employee struct {
	empid                     int
	firstname, lastname, role string
	basicSalary, totalSalary  int
}

//func constructor
func listEmployees(empid int, firstname, lastname, role string, basicSalary, totalSalary int) *Employee {
	return &Employee{empid, firstname, lastname, role, basicSalary, totalSalary}
}

var emp int = 119

//method
func (tunjangan *Employee) calcTunjangan(bonus, commision int) int {
	emp++
	tunjangan.empid = emp

	if tunjangan.role == "Programmer" {
		tunjangan.basicSalary = 6000000
		tunjangan.totalSalary = tunjangan.basicSalary + 500000
	} else if tunjangan.role == "Sales" {
		tunjangan.basicSalary = 3000000
		tunjangan.totalSalary = tunjangan.basicSalary + bonus + commision
	} else if tunjangan.role == "QA" {
		tunjangan.basicSalary = 4500000
		tunjangan.totalSalary = tunjangan.basicSalary + 10000
	}
	return tunjangan.totalSalary
}

func main() {
	//initial constructor
	emp1 := listEmployees(0, "Anton", "Pratama", "Programmer", 0, 0)
	emp2 := listEmployees(0, "Budi", "Junaidi", "Programmer", 0, 0)
	emp3 := listEmployees(0, "Charlie", "Van Dijk", "Programmer", 0, 0)
	emp4 := listEmployees(0, "Dian", "Permana", "Sales", 0, 0)
	emp5 := listEmployees(0, "Fatur", "Rohman", "Sales", 0, 0)
	emp6 := listEmployees(0, "Ellise", "Toon", "QA", 0, 0)
	emp7 := listEmployees(0, "Gita", "Gutawa", "QA", 0, 0)

	emp1.totalSalary = emp1.calcTunjangan(0, 0)
	emp2.totalSalary = emp2.calcTunjangan(0, 0)
	emp3.totalSalary = emp3.calcTunjangan(0, 0)
	emp4.totalSalary = emp4.calcTunjangan(500000, 200000)
	emp5.totalSalary = emp5.calcTunjangan(350000, 250000)
	emp6.totalSalary = emp6.calcTunjangan(0, 0)
	emp7.totalSalary = emp7.calcTunjangan(0, 0)
	employee := []Employee{*emp1, *emp2, *emp3, *emp4, *emp5, *emp6, *emp7}
	count := 0
	basic := 0
	basicAll := 0
	totEmpProg := 0
	totEmpSales := 0
	totEmpQA := 0
	totSalProg := 0
	totSalSales := 0
	totSalQA := 0
	for _, v := range employee {
		fmt.Println(v)
		count++
		basic += v.basicSalary
		basicAll += v.totalSalary
		if v.role == "Programmer" {
			totEmpProg++
			totSalProg += v.totalSalary
		} else if v.role == "Sales" {
			totEmpSales++
			totSalSales += v.totalSalary
		} else {
			totEmpQA++
			totSalQA += v.totalSalary
		}
	}
	fmt.Println("---Summary---")
	fmt.Println("Total Employee :", count)
	fmt.Println("Total Basic Salary :", basic)
	fmt.Println("Total Salary All : ", basicAll)
	fmt.Println("---Total Employee by Role---")
	fmt.Println("Programmer : ", totEmpProg)
	fmt.Println("Sales : ", totEmpSales)
	fmt.Println("QA : ", totEmpQA)
	fmt.Println("---Total Employee Salary by Role---")
	fmt.Println("Programmer : ", totSalProg)
	fmt.Println("Sales : ", totSalSales)
	fmt.Println("QA : ", totSalQA)
}
