package main

import (
	"fmt"
)

type Employee struct {
	EmpId               int
	FirstName           string
	LastName            string
	Role                string
	BasicSalary         float64
	SalesBonus          float64
	SalesComission      float64
	ProgrammerAllowance float64
	QaAllowance         float64
}

// construct
func listEmployees(empid int, firstname, lastname, role string, basicSalary, salesBonus, salesComission, programmerAllowance, qaAllowance float64) *Employee {
	return &Employee{
		EmpId:               empid,
		FirstName:           firstname,
		LastName:            lastname,
		Role:                role,
		BasicSalary:         basicSalary,
		SalesBonus:          salesBonus,
		SalesComission:      salesComission,
		ProgrammerAllowance: programmerAllowance,
		QaAllowance:         qaAllowance,
	}
}

// mencetak metod detail data employes
func (e *Employee) PrintEmployeeDetails() {
	fmt.Printf("Employee : %d\n", e.EmpId)
	fmt.Printf("Name : %s\n", e.FirstName)
	fmt.Printf("Name : %s\n", e.LastName)
	fmt.Printf("Role : %s\n", e.Role)
	fmt.Printf("Basic Salary : %.2f\n", e.BasicSalary)
	fmt.Printf("Sales Bonus : %.2f\n", e.SalesBonus)
	fmt.Printf("Sales Comission : %.2f\n", e.SalesComission)
	fmt.Printf("Programmer Allowance : %.2f\n", e.ProgrammerAllowance)
	fmt.Printf("Qa Allowance : %.2f\n", e.QaAllowance)
	fmt.Printf("Total Salary : %.2f\n\n", e.BasicSalary+e.SalesBonus+e.SalesComission+e.ProgrammerAllowance+e.QaAllowance)
}

// mencetak semua detail data employes
func PrintAllEmployeeDetails(employees []*Employee) {
	for _, emp := range employees {
		emp.PrintEmployeeDetails()
	}
}

// mengkalkulasi dan mencetak nilai summary
func PrintSummary(employees []*Employee) {
	var totalBasicSalary float64
	var totalSalary float64
	employeeByRole := make(map[string]int)
	employeeSalaryByRole := make(map[string]float64)

	for _, emp := range employees {
		totalBasicSalary += emp.BasicSalary

		totalSalary += emp.BasicSalary + emp.SalesBonus + emp.SalesComission + emp.ProgrammerAllowance + emp.QaAllowance

		employeeByRole[emp.Role]++

		employeeSalaryByRole[emp.Role] += emp.BasicSalary + emp.SalesBonus + emp.SalesComission + emp.ProgrammerAllowance + emp.QaAllowance
	}

	fmt.Printf("Total Basic Salary : Rp.%.2f\n", totalBasicSalary)
	fmt.Printf("Total Salary : Rp.%.2f\n\n", totalSalary)

	fmt.Println("Total Employee By Role")
	for role, count := range employeeByRole {
		fmt.Printf("%s: %d\n", role, count)
	}
	fmt.Println()

	fmt.Println("Total Employee Salary By Role")
	for role, salary := range employeeSalaryByRole {
		fmt.Printf("%s: Rp.%.2f\n", role, salary)
	}

}

func main() {
	employee := []*Employee{
		// listEmployees(22, "Kus", "Suma", "Programmer", 5_000_000, 0, 0, 200_000, 0),
		listEmployees(120, "Anton", "Pratama", "Programmer", 6_000_000, 0, 0, 500_000, 0),
		listEmployees(121, "Budi", "Junaedi", "Programmer", 6_000_000, 0, 0, 500_000, 0),
		listEmployees(122, "Charlie", "Van Dijk", "Programmer", 6_000_000, 0, 0, 500_000, 0),
		listEmployees(123, "Dian", "Permana", "Sales", 3_000_000, 500_000, 200_000, 0, 0),
		listEmployees(124, "Fatur", "Rohman", "Sales", 3_000_000, 350_000, 250_000, 0, 0),
		listEmployees(125, "Ellise", "Toon", "QA", 4500_000, 0, 0, 0, 1_0000),
		listEmployees(126, "Gita", "Gutawa", "QA", 4500_000, 0, 0, 0, 0),
	}
	//menampilkan hasil print detail employe
	PrintAllEmployeeDetails(employee)

	//menampilkan hasil print summary
	PrintSummary(employee)
}
