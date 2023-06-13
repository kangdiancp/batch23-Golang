package main

import (
	"fmt"
)

// Struct Employees
type Employees struct {
	EmployeesID     int
	FirstName       string
	LastName        string
	Role            string
	BasicSalary     float64
	ProgrammerTunj  float64
	SalesBonus      float64
	SalesCommission float64
	QATunj          float64
	TotalSalary     float64
}

// Constructor untuk membuat instance Employees
func NewEmployees(employeesID int, firstName, lastName, role string) *Employees {
	return &Employees{
		EmployeesID: employeesID,
		FirstName:   firstName,
		LastName:    lastName,
		Role:        role,
	}
}

// Method untuk menghitung gaji
func (emp *Employees) CalculateSalary() {
	switch emp.Role {
	case "Programmer":
		emp.BasicSalary = 6_000_000
		emp.ProgrammerTunj = 500_000
	case "Sales":
		emp.BasicSalary = 3_000_000
		if emp.EmployeesID == 123 {
			emp.SalesBonus = 500_000
			emp.SalesCommission = 200_000
		} else if emp.EmployeesID == 124 {
			emp.SalesBonus = 350_000
			emp.SalesCommission = 250_000
		}
	case "QA":
		emp.BasicSalary = 4_500_000
		emp.QATunj = 10_000
	}
	emp.TotalSalary = emp.BasicSalary + emp.ProgrammerTunj + emp.SalesBonus + emp.SalesCommission + emp.QATunj
}

func listEmployees(employees []*Employees) {
	fmt.Println("EmployeesID  FirstName  LastName    Role       BasicSalary    ProgrammerTunj    SalesBonus   SalesCommission       QATunj       TotalSalary")
	TotalBasicSalary := 0.0
	TotalSalary := 0.0
	TotalEmployeeByRole := make(map[string]int)
	TotalEmployeeSalaryByRole := make(map[string]float64)

	for _, emp := range employees {
		emp.CalculateSalary() // Panggil method CalculateSalary()
		fmt.Printf("%-12d%-12s%-12s%-12sRp. %-14.0fRp. %-12.0fRp. %-10.0fRp. %-16.0fRp. %-8.0fRp. %.0f\n", emp.EmployeesID, emp.FirstName, emp.LastName, emp.Role, emp.BasicSalary, emp.ProgrammerTunj, emp.SalesBonus, emp.SalesCommission, emp.QATunj, emp.TotalSalary)
		// Update Total Basic Salary
		TotalBasicSalary += emp.BasicSalary
		// Update Total Salary
		TotalSalary += emp.TotalSalary

		// Update Total Employee By Role
		TotalEmployeeByRole[emp.Role]++

		// Update Total Employee Salary By Role
		TotalEmployeeSalaryByRole[emp.Role] += emp.TotalSalary
	}
	// Update Total Employee
	totalEmployee := len(employees)

	fmt.Printf("Total Employee: %d\n", totalEmployee)
	fmt.Printf("Total Basic Salary: Rp. %.f\n", TotalBasicSalary)
	fmt.Printf("Total Salary: Rp. %.f\n\n", TotalSalary)

	fmt.Println("Total Employee By Role and Employee Salary By Role:")
	for role, count := range TotalEmployeeByRole {
		fmt.Printf("%s:\n", role)
		fmt.Printf("  Count: %d\n", count)
		fmt.Printf("  Total Salary: Rp. %.f\n\n", TotalEmployeeSalaryByRole[role])
	}
}

func main() {
	employees := []*Employees{
		NewEmployees(120, "Anton", "Pratama", "Programmer"), // Gunakan constructor untuk membuat instance Employee
		NewEmployees(121, "Budi", "Junaedi", "Programmer"),
		NewEmployees(122, "Charlie", "Van Dijk", "Programmer"),
		NewEmployees(123, "Dian", "Permana", "Sales"),
		NewEmployees(124, "Fatur", "Rohman", "Sales"),
		NewEmployees(125, "Elise", "Toon", "QA"),
		NewEmployees(126, "Gita", "Gutawa", "QA"),
	}

	listEmployees(employees)
}
