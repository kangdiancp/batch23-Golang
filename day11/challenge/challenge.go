package main

import "fmt"

type Employees struct {
	EmpId                                                     int
	FirstName, LastName, Role                                 string
	basicSalary, bonus, commision, lembur, makan, totalSalary int
}

type Employee struct {
	employee []Employees
}

func NewEmplooye(EmpId int, FirstName, LastName, Role string, basicSalary, bonus, commision, lembur, makan, totalSalary int) *Employees {
	if Role == "Programmer" {
		basicSalary = 6000000
		bonus = 0
		commision = 0
		lembur = 500000
		makan = 0
	} else if Role == "Sales" {
		basicSalary = 3000000
	} else {
		basicSalary = 4500000
		bonus = 0
		commision = 0
		lembur = 0
		makan = 10000
	}

	totalSalary = basicSalary + bonus + commision + lembur + makan
	return &Employees{EmpId, FirstName, LastName, Role, basicSalary, bonus, commision, lembur, makan, totalSalary}
}

func (employ *Employee) totalBasicSalary() int {
	totalBasic := 0
	for _, e := range employ.employee {
		totalBasic += e.basicSalary
	}
	return totalBasic
}

func (employ *Employee) totalAllSalary() int {
	totalAll := 0
	for _, e := range employ.employee {
		totalAll += e.totalSalary
	}
	return totalAll
}

func (employ *Employee) totalSalaryByRole(Role string) int {
	total := 0
	for _, e := range employ.employee {
		if e.Role == Role {
			total += e.totalSalary
		}
	}
	return total
}

func (employ *Employee) totalEmployeeByRole(Role string) int {
	total := 0
	for _, e := range employ.employee {
		if e.Role == Role {
			total++
		}
	}
	return total
}

func main() {
	emp1 := NewEmplooye(120, "Anton", "Pratama", "Programmer", 0, 0, 0, 0, 0, 0)
	emp2 := NewEmplooye(121, "Budi", "Junaedi", "Programmer", 0, 0, 0, 0, 0, 0)
	emp3 := NewEmplooye(122, "Charlie", "Van Dijk", "Programmer", 0, 0, 0, 0, 0, 0)
	emp4 := NewEmplooye(123, "Dian", "Permana", "Sales", 0, 500000, 200000, 0, 0, 0)
	emp5 := NewEmplooye(125, "Fatur", "Rohman", "Sales", 0, 350000, 250000, 0, 0, 0)
	emp6 := NewEmplooye(124, "Ellise", "Toon", "QA", 0, 0, 0, 0, 0, 0)
	emp7 := NewEmplooye(126, "Gita", "Gutawa", "QA", 0, 0, 0, 0, 0, 0)

	employees := []Employees{*emp1, *emp2, *emp3, *emp4, *emp5, *emp6, *emp7}
	for _, v := range employees {
		fmt.Println(v)
	}
	// fmt.Println(employees)

	total1 := &Employee{employee: employees}
	totalBasicSalary := total1.totalBasicSalary()
	fmt.Println("\nTotal Basic Salary	: ", totalBasicSalary)

	total2 := &Employee{employee: employees}
	totalAllSalary := total2.totalAllSalary()
	fmt.Println("Total All Salary	: ", totalAllSalary, "\n")

	total3 := &Employee{employee: employees}
	totalSalaryProgrammer := total3.totalSalaryByRole("Programmer")
	totalSalarySales := total3.totalSalaryByRole("Sales")
	totalSalaryQA := total3.totalSalaryByRole("QA")
	fmt.Println("Total Salary Programmer	: ", totalSalaryProgrammer)
	fmt.Println("Total Salary Sales	: ", totalSalarySales)
	fmt.Println("Total Salary QA		: ", totalSalaryQA, "\n")

	total4 := &Employee{employee: employees}
	totalProgrammer := total4.totalEmployeeByRole("Programmer")
	totalSales := total4.totalEmployeeByRole("Sales")
	totalQA := total4.totalEmployeeByRole("QA")
	fmt.Println("Total Programmer	: ", totalProgrammer)
	fmt.Println("Total Sales		: ", totalSales)
	fmt.Println("Total QA		: ", totalQA)

}
