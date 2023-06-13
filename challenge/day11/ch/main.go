package main

import "fmt"

func main() {

	employees := []*Employee{}

	fmt.Println("")
	emp1 := NewEmployee(1, "Anton", "Pratama", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp2 := NewEmployee(2, "Budi", "Junaedi", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp3 := NewEmployee(3, "Charlie", "Van Dijk", "Programmer", 6_000_000, 0, 0, 500_000, 0)
	emp4 := NewEmployee(4, "Dian", "Permana", "Sales", 3_000_000, 500_000, 200_000, 0, 0)
	emp5 := NewEmployee(5, "Fatur", "Rohman", "Sales", 3_000_000, 350_000, 250_000, 0, 0)
	emp6 := NewEmployee(6, "Ellise", "Toon", "QA", 4_500_000, 0, 0, 0, 10_000)
	emp7 := NewEmployee(7, "Gita", "Gutawa", "QA", 4_500_000, 0, 0, 0, 10_000)
	employees = append(employees, emp1, emp2, emp3, emp4, emp5, emp6, emp7)

	for _, emp := range employees {
		emp.CalculateTotalSalary()
		emp.Display()
	}

	fmt.Println("")

	fmt.Printf("Total Employees : %d\n", len(employees))

	var totalBasicSalary float64
	for _, emp := range employees {
		totalBasicSalary += emp.basicSalary
	}
	fmt.Printf("Total Basic Salary : %.2f\n", totalBasicSalary)

	var totalSalaryAll float64
	for _, emp := range employees {
		totalSalaryAll += emp.totalGaji
	}
	fmt.Printf("Total Salary All : %.2f\n", totalSalaryAll)

	fmt.Println("")

	fmt.Printf("Total Employees by Role : \n")
	roles := make(map[string]int)
	for _, emp := range employees {
		roles[emp.role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}

	fmt.Println("")

	fmt.Printf("Total Employees Salary by Role : \n")
	salaryByRole := make(map[string]float64)
	for _, emp := range employees {
		salaryByRole[emp.role] += emp.totalGaji
	}
	for role, count := range salaryByRole {
		fmt.Printf("%s : %.2f\n", role, count)
	}
}

type Employee struct {
	empId           int
	firstName       string
	lastName        string
	role            string
	basicSalary     float64
	salesBonus      float64
	salesCommission float64
	gajiProgrammer  float64
	gajiQA          float64
	totalGaji       float64
}

//constructor

func NewEmployee(empId int, firstName, lastName, role string, basicSalary,
	salesBonus, salesCommission, gajiProgrammer, gajiQA float64) *Employee {
	return &Employee{
		empId:           empId,
		firstName:       firstName,
		lastName:        lastName,
		role:            role,
		basicSalary:     basicSalary,
		salesBonus:      salesBonus,
		salesCommission: salesCommission,
		gajiProgrammer:  gajiProgrammer,
		gajiQA:          gajiQA,
	}
}

func (e *Employee) CalculateTotalSalary() {
	switch e.role {
	case "Programmer":
		e.totalGaji = e.basicSalary + e.gajiProgrammer
	case "Sales":
		e.totalGaji = e.basicSalary + e.salesBonus + e.salesCommission
	case "QA":
		e.totalGaji = e.basicSalary + e.gajiQAgit
	default:
		e.totalGaji = e.basicSalary
	}
}

func (e *Employee) Display() {
	fmt.Printf("EmpID: %d\n", e.empId)
	fmt.Printf("First Name: %s\n", e.firstName)
	fmt.Printf("Last Name: %s\n", e.lastName)
	fmt.Printf("Role: %s\n", e.role)
	fmt.Printf("Basic Salary: %.2f\n", e.basicSalary)
	fmt.Printf("Sales Bonus: %.2f\n", e.salesBonus)
	fmt.Printf("Sales Commission: %.2f\n", e.salesCommission)
	fmt.Printf("Gaji Programmer: %.2f\n", e.gajiProgrammer)
	fmt.Printf("Gaji QA: %.2f\n", e.gajiQA)
	fmt.Printf("Total Gaji: %.2f\n\n", e.totalGaji)
}
