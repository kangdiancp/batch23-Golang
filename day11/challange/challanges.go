package main

import "fmt"

// Struct Employee merepresentasikan data employee
type Employee struct {
	empId               int
	firstName           string
	lastName            string
	role                string
	basicSalary         float64
	tunjanganLembur     float64
	tunjanganBonus      float64
	tunjanganCommission float64
	tunjanganMakan      float64
}

// Constructor NewEmployee untuk membuat objek Employee baru
func NewEmployee(empId int, firstName, lastName, role string, basicSalary, tunjanganLembur, tunjanganBonus, tunjanganCommission, tunjanganMakan float64) *Employee {
	// Mengembalikan pointer ke objek Employee yang baru dibuat
	return &Employee{empId, firstName, lastName, role, basicSalary, tunjanganLembur, tunjanganBonus, tunjanganCommission, tunjanganMakan}
}

// Method PrintEmployeeDetails untuk mencetak detail employee
func (e *Employee) PrintEmployeeDetails() {
	fmt.Printf("Employeed: %d\n", e.empId)
	fmt.Printf("Name: %s %s\n", e.firstName, e.lastName)
	fmt.Printf("Role: %s\n", e.role)
	fmt.Printf("Basic Salary: Rp.%.2f\n", e.basicSalary)
	fmt.Printf("Tunjangan Lembur: Rp.%.2f\n", e.tunjanganLembur)
	fmt.Printf("Tunjangan Bonus: Rp.%.2f\n", e.tunjanganBonus)
	fmt.Printf("Tunjangan Commission: Rp.%.2f\n", e.tunjanganCommission)
	fmt.Printf("Tunjangan Makan: Rp.%.2f\n", e.tunjanganMakan)
	fmt.Printf("Total Salary: Rp.%.2f\n\n", e.basicSalary+e.tunjanganLembur+e.tunjanganBonus+e.tunjanganCommission+e.tunjanganMakan)
}

// Function untuk mencetak detail semua employee
func PrintAllEmployeeDetails(employees []*Employee) {
	// Melakukan iterasi melalui array employees menggunakan range
	for _, emp := range employees {
		// Memanggil method PrintEmployeeDetails untuk mencetak detail employee
		emp.PrintEmployeeDetails()
	}
}

// Function untuk menghitung dan mencetak summary data employee
func PrintSummary(employees []*Employee) {
	var totalBasicSalary float64
	var totalSalary float64
	employeeByRole := make(map[string]int)
	employeeSalaryByRole := make(map[string]float64)

	// Melakukan iterasi melalui array employees menggunakan range
	for _, emp := range employees {
		// Menghitung totalBasicSalary dan totalSalary
		totalBasicSalary += emp.basicSalary
		totalSalary += emp.basicSalary + emp.tunjanganLembur + emp.tunjanganBonus + emp.tunjanganCommission + emp.tunjanganMakan

		// Menghitung jumlah employee berdasarkan role
		employeeByRole[emp.role]++

		// Menghitung total gaji employee berdasarkan role
		employeeSalaryByRole[emp.role] += emp.basicSalary + emp.tunjanganLembur + emp.tunjanganBonus + emp.tunjanganCommission + emp.tunjanganMakan
	}

	// Mencetak totalBasicSalary dan totalSalary
	fmt.Printf("Total Basic Salary: Rp.%.2f\n", totalBasicSalary)
	fmt.Printf("Total Salary: Rp.%.2f\n\n", totalSalary)

	// Mencetak jumlah employee berdasarkan role
	fmt.Println("Total Employee By Role")
	for role, count := range employeeByRole {
		fmt.Printf("%s: %d\n", role, count)
	}
	fmt.Println()

	// Mencetak total salary employee berdasarkan role
	fmt.Println("Total Employee Salary By Role")
	for role, salary := range employeeSalaryByRole {
		fmt.Printf("%s: Rp.%.2f\n", role, salary)
	}
}

func main() {
	// membuat list employees
	employees := []*Employee{
		NewEmployee(120, "Anton", "Pratama", "Programmer", 6000000, 500000, 0, 0, 0),
		NewEmployee(121, "Budi", "Junaedi", "Programmer", 6000000, 500000, 0, 0, 0),
		NewEmployee(122, "Charlie", "Van Dijk", "Programmer", 6000000, 500000, 0, 0, 0),
		NewEmployee(123, "Dian", "Permana", "Sales", 3000000, 0, 500000, 200000, 0),
		NewEmployee(124, "Fatur", "Rohman", "Sales", 3000000, 0, 350000, 250000, 0),
		NewEmployee(125, "Ellise", "Toon", "QA", 4500000, 0, 0, 0, 10000),
		NewEmployee(126, "Gita", "Gutawa", "QA", 4500000, 0, 0, 0, 10000),
	}

	// Memanggil fungsi untuk mencetak detail semua employee
	PrintAllEmployeeDetails(employees)

	// memanggil function summary
	PrintSummary(employees)
}
