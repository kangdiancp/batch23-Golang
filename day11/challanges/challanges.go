package main

import "fmt"

type Employees struct {
	empId                                                      int
	firstName, lastName, role                                  string
	basicSalary, bonus, commission, lembur, makan, totalSalary int
}

func newEmployees(empId int, firstName, lastName, role string, basicSalary, bonus, commission, lembur, makan int) *Employees {
	totalSalary := bonus + commission + lembur + makan
	return &Employees{empId, firstName, lastName, role, basicSalary, bonus, commission, lembur, makan, totalSalary}
}

func totalBasicSalary(employees ...*Employees) int {
	totalBasicSalary := 0

	for _, emp := range employees {
		totalBasicSalary += emp.basicSalary
	}

	return totalBasicSalary
}

func totalSalaryAll(employees ...*Employees) int {
	totalSalaryAll := 0

	for _, emp := range employees {
		totalSalaryAll += emp.totalSalary
	}

	return totalSalaryAll
}

func totalSalaryByRole(employees []*Employees, role string) int {
	totalSalaryByRole := 0

	for _, emp := range employees {
		if emp.role == role {
			totalSalaryByRole += emp.totalSalary
		}
	}

	return totalSalaryByRole
}

func totalEmployeesByRole(employees []*Employees, role string) int {
	totalEmployeesByRole := 0

	for _, emp := range employees {
		if emp.role == role {
			totalEmployeesByRole++
		}
	}

	return totalEmployeesByRole
}

func (employees *Employees) calcTotalSalary() {
	totalSalary := employees.basicSalary + employees.bonus + employees.commission + employees.lembur + employees.makan
	employees.totalSalary = totalSalary
}

func main() {
	emp1 := newEmployees(120, "Anton", "Pratama", "Programmer", 6000_000, 0, 0, 500_000, 0)
	emp2 := newEmployees(121, "Budi", "Junaedi", "Programmer", 6000_000, 0, 0, 500_000, 0)
	emp3 := newEmployees(122, "Charlie", "Van Dijk", "Programmer", 6000_000, 0, 0, 500_000, 0)
	emp4 := newEmployees(123, "Dian", "Permana", "Sales", 3000_000, 500_000, 200_000, 0, 0)
	emp5 := newEmployees(124, "Fatur", "Rohman", "Sales", 3000_000, 350_000, 250_000, 0, 0)
	emp6 := newEmployees(125, "Ellise", "Toon", "QA", 4500_000, 0, 0, 0, 10_000)
	emp7 := newEmployees(126, "Gita", "Gutawa", "QA", 4500_000, 0, 0, 0, 10_000)

	// emp1.totalSalary = emp1.calctotalSalary()
	// emp2.totalSalary = emp2.calctotalSalary()
	// emp3.totalSalary = emp3.calctotalSalary()
	// emp4.totalSalary = emp4.calctotalSalary()
	// emp5.totalSalary = emp5.calctotalSalary()
	// emp6.totalSalary = emp6.calctotalSalary()
	// emp7.totalSalary = emp7.calctotalSalary()

	employees := []*Employees{emp1, emp2, emp3, emp4, emp5, emp6, emp7}

	for _, emp := range employees {
		emp.calcTotalSalary()
	}

	for _, emp := range employees {
		if emp.role == "Programmer" {
			fmt.Printf("Employee ID: %d\n", emp.empId)
			fmt.Printf("Name: %s %s\n", emp.firstName, emp.lastName)
			fmt.Printf("Role: %s\n", emp.role)
			fmt.Printf("Basic Salary: %d\n", emp.basicSalary)
			fmt.Printf("Tunjangan Programmer Lembur: %d\n", emp.lembur)
			fmt.Printf("Total Salary: %d\n", emp.totalSalary)
			fmt.Println("-------------------")
		} else if emp.role == "Sales" {
			fmt.Printf("Employee ID: %d\n", emp.empId)
			fmt.Printf("Name: %s %s\n", emp.firstName, emp.lastName)
			fmt.Printf("Role: %s\n", emp.role)
			fmt.Printf("Basic Salary: %d\n", emp.basicSalary)
			fmt.Printf("Tunjangan Sales Bonus: %d\n", emp.bonus)
			fmt.Printf("Tunjangan Sales Commission: %d\n", emp.commission)
			fmt.Printf("Total Salary: %d\n", emp.totalSalary)
			fmt.Println("-------------------")
		} else {
			fmt.Printf("Employee ID: %d\n", emp.empId)
			fmt.Printf("Name: %s %s\n", emp.firstName, emp.lastName)
			fmt.Printf("Role: %s\n", emp.role)
			fmt.Printf("Basic Salary: %d\n", emp.basicSalary)
			fmt.Printf("Tunjangan QA Makan: %d\n", emp.makan)
			fmt.Printf("Total Salary: %d\n", emp.totalSalary)
			fmt.Println("-------------------")
		}
	}

	fmt.Println("Total Basic Salary : Rp.", totalBasicSalary(employees...))
	fmt.Println("Total Salary All : Rp.", totalSalaryAll(employees...))

	fmt.Println("\nTotal Employees Programmer : Rp.", totalEmployeesByRole(employees, "Programmer"))
	fmt.Println("Total Employees Sales : Rp.", totalEmployeesByRole(employees, "Sales"))
	fmt.Println("Total Employees QA : Rp.", totalEmployeesByRole(employees, "QA"))

	fmt.Println("\nTotal Salary Programmer : Rp.", totalSalaryByRole(employees, "Programmer"))
	fmt.Println("Total Salary Sales : Rp.", totalSalaryByRole(employees, "Sales"))
	fmt.Println("Total Salary QA : Rp.", totalSalaryByRole(employees, "QA"))
}
