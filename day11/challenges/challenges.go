package main

import "fmt"

type Employee struct {
	EmpId                                                                                   int
	FirstName, LastName, Role                                                               string
	BasicSalary, SalesBonus, SalesCommission, ProgrammerAllowance, QAAllowance, TotalSalary float64
}

func NewEmployee(empID int, firstName, lastName, role string, basicSalary, salesBonus, salesCommission, programmerAllowance, qaAllowance float64) *Employee {
	return &Employee{
		EmpId:               empID,
		FirstName:           firstName,
		LastName:            lastName,
		Role:                role,
		BasicSalary:         basicSalary,
		SalesBonus:          salesBonus,
		SalesCommission:     salesCommission,
		ProgrammerAllowance: programmerAllowance,
		QAAllowance:         qaAllowance,
	}
}

func (e *Employee) CalculateTotalSalary() {
	switch e.Role {
	case "Programmer":
		e.TotalSalary = e.BasicSalary + e.ProgrammerAllowance
	case "Sales":
		e.TotalSalary = e.BasicSalary + e.SalesBonus + e.SalesCommission
	case "QA":
		e.TotalSalary = e.BasicSalary + e.QAAllowance
	default:
		e.TotalSalary = e.BasicSalary
	}
}

func (e *Employee) Display() {
	fmt.Printf("EmpID: %d\n", e.EmpId)
	fmt.Printf("First Name: %s\n", e.FirstName)
	fmt.Printf("Last Name: %s\n", e.LastName)
	fmt.Printf("Role: %s\n", e.Role)
	fmt.Printf("Basic Salary: %.2f\n", e.BasicSalary)
	fmt.Printf("Sales Bonus: %.2f\n", e.SalesBonus)
	fmt.Printf("Sales Commission: %.2f\n", e.SalesCommission)
	fmt.Printf("Programmer Allowance: %.2f\n", e.ProgrammerAllowance)
	fmt.Printf("QA Allowance: %.2f\n", e.QAAllowance)
	fmt.Printf("Total Salary: %.2f\n\n", e.TotalSalary)
}

func main() {
	employees := []*Employee{}

	emp1 := NewEmployee(1, "Fatahillah", ".", "Programmer", 5_555_555, 0, 0, 2_345_678, 0)
	emp2 := NewEmployee(2, "Adjik", "LT", "Sales", 4_567_890, 1_234_567, 1_111_111, 0, 0)
	emp3 := NewEmployee(3, "Ralif", "CA", "QA", 5_000_000, 0, 0, 0, 2_222_222)

	employees = append(employees, emp1, emp2, emp3)

	for _, emp := range employees {
		emp.CalculateTotalSalary()
		emp.Display()
	}

	fmt.Printf("-------------- Summary --------------\n")

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

	fmt.Printf("-------------------------------------\n")

	fmt.Printf("Total Employees by Role : \n")
	roles := make(map[string]int)
	for _, emp := range employees {
		roles[emp.Role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}

	fmt.Printf("-------------------------------------\n")

	fmt.Printf("Total Employees Salary by Role : \n")
	salaryByRole := make(map[string]float64)
	for _, emp := range employees {
		salaryByRole[emp.Role] += emp.TotalSalary
	}
	for role, count := range salaryByRole {
		fmt.Printf("%s : %.2f\n", role, count)
	}
}
