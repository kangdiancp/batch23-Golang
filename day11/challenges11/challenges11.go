package main

import (
	"fmt"
)

type Employees struct {
	EmpId                                                                                   int
	FirstName, LastName, Role                                                               string
	BasicSalary, SalesBonus, SalesCommission, tunjanganProgrammer, tunjanganQA, TotalSalary float64
}

// func constructor
func NewEmployee(empID int, firstName, lastName, role string, basicSalary, salesBonus, salesCommission, tunjanganProgrammer, tunjanganQA float64) *Employees {
	return &Employees{
		EmpId:               empID,
		FirstName:           firstName,
		LastName:            lastName,
		Role:                role,
		BasicSalary:         basicSalary,
		SalesBonus:          salesBonus,
		SalesCommission:     salesCommission,
		tunjanganProgrammer: tunjanganProgrammer,
		tunjanganQA:         tunjanganQA,
	}
}

// method function
func (e *Employees) CalcTotalSalary() {
	switch e.Role {
	case "Programmer":
		e.TotalSalary = e.BasicSalary + e.tunjanganProgrammer
	case "Sales":
		e.TotalSalary = e.BasicSalary + e.SalesBonus + e.SalesCommission
	case "QA":
		e.TotalSalary = e.BasicSalary + e.tunjanganQA
	default:
		e.TotalSalary = e.BasicSalary
	}
}

func (e *Employees) Display() {
	fmt.Printf("%-12d%-12s%-12s%-12sRp. %-14.0fRp. %-12.0fRp. %-10.0fRp. %-16.0fRp. %-8.0fRp. %.0f\n", e.EmpId, e.FirstName, e.LastName, e.Role, e.BasicSalary, e.SalesBonus, e.SalesCommission, e.tunjanganProgrammer, e.tunjanganQA, e.TotalSalary)
}

func main() {
	employees := []*Employees{}

	fmt.Printf("-------------------------------------------------------------------------------------------------------------------------------------------\n")

	emp1 := NewEmployee(1, "Anton", "Pratama", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp2 := NewEmployee(2, "Budi", "Junaedi", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp3 := NewEmployee(3, "Charlie", "Van Dijk", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp4 := NewEmployee(4, "Dian", "Permana", "Sales", 3_000_000, 500_000, 200_000, 0, 0)
	emp5 := NewEmployee(5, "Fatur", "Rohman", "Sales", 3_000_000, 350_000, 250_000, 0, 0)
	emp6 := NewEmployee(6, "Ellise", "Toon", "QA", 4_500_000, 0, 0, 0, 10_000)
	emp7 := NewEmployee(7, "Gita", "Gutawa", "QA", 4_500_000, 0, 0, 0, 10_000)

	employees = append(employees, emp1, emp2, emp3, emp4, emp5, emp6, emp7)

	for _, emp := range employees {
		emp.CalcTotalSalary()
		emp.Display()
	}

	fmt.Printf("-------------------------------------------------------------------------------------------------------------------------------------------\n")

	fmt.Println()

	fmt.Printf("Total Employees : %d\n", len(employees))

	var totalBasicSalary float64
	for _, emp := range employees {
		totalBasicSalary += emp.BasicSalary
	}
	fmt.Printf("Total Basic Salary : %.2f\n", totalBasicSalary)

	var totalSalaryAll float64
	for _, emp := range employees {
		totalSalaryAll += emp.TotalSalary
	}
	fmt.Printf("Total Salary All : %.2f\n", totalSalaryAll)

	fmt.Println()

	fmt.Printf("Total Employees by Role : \n")
	roles := make(map[string]int)
	for _, emp := range employees {
		roles[emp.Role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}

	fmt.Println()

	fmt.Printf("Total Employees Salary by Role : \n")
	salaryByRole := make(map[string]float64)
	for _, emp := range employees {
		salaryByRole[emp.Role] += emp.TotalSalary
	}
	for role, count := range salaryByRole {
		fmt.Printf("%s : %.2f\n", role, count)
	}
}
