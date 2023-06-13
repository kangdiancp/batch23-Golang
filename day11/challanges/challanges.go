package main

import "fmt"

//struct
var nextEmpID = 120

type mainStruct struct {
	empId                     int
	firstName, lastName, role string
	basicSalary               int
	tunjangan                 *Tunjangan
	totalSalary               int
}
type Tunjangan struct {
	tunjanganSales      *TunjanganSales
	tunjanganProgrammer *TunjanganProgrammer
	tunjanganQA         *TunjanganQA
}

type TunjanganSales struct {
	bonus  int
	komisi int
}
type TunjanganProgrammer struct {
	lembur int
}
type TunjanganQA struct {
	makan int
}

type SalaryCalculator struct {
	employees []mainStruct
}

func newEmployee(firstName, lastName, role string, bonus, komisi int) *mainStruct {
	e := &mainStruct{
		empId:     nextEmpID,
		firstName: firstName,
		lastName:  lastName,
		role:      role,
	}
	nextEmpID++

	if role == "Programmer" {
		e.basicSalary = 6_000_000
		e.tunjangan = &Tunjangan{
			tunjanganProgrammer: &TunjanganProgrammer{lembur: 500_000},
		}
		e.totalSalary = e.basicSalary + e.tunjangan.tunjanganProgrammer.lembur

	} else if role == "QA" {
		e.basicSalary = 4_500_000
		e.tunjangan = &Tunjangan{
			tunjanganQA: &TunjanganQA{makan: 10_000},
		}
		e.totalSalary = e.basicSalary + e.tunjangan.tunjanganQA.makan
	} else if role == "Sales" {
		e.basicSalary = 3_000_000
		e.tunjangan = &Tunjangan{
			tunjanganSales: &TunjanganSales{bonus: bonus, komisi: komisi},
		}
		e.totalSalary = e.basicSalary + e.tunjangan.tunjanganSales.bonus + e.tunjangan.tunjanganSales.komisi

	}

	return e
}

// Method untuk menghitung total basic salary
func (calc *SalaryCalculator) SumBasicSalary() int {
	total := 0
	for _, emp := range calc.employees {
		total += emp.basicSalary
	}
	return total
}

//methode untuk menghitung  total salary
func (calc *SalaryCalculator) SumTotalSalary() int {
	total := 0
	for _, emp := range calc.employees {
		total += emp.totalSalary
	}
	return total
}

//methode untuk menghitung total salary berdasarkan role
func (calc *SalaryCalculator) TotalSalaryByRole(role string) int {
	total := 0
	for _, emp := range calc.employees {
		if emp.role == role {
			total += emp.totalSalary
		}
	}
	return total
}

//methode untuk menghitung total employee berdasarkan role
func (calc *SalaryCalculator) TotalEmployeeByRole(role string) int {
	total := 0
	for _, emp := range calc.employees {
		if emp.role == role {
			total++
		}
	}
	return total
}

func main() {
	// tunjanganSales := &TunjanganSales{500_000, 200_000}
	// tunjanganProgrammer := &TunjanganProgrammer{0}
	// tunjanganQA := &TunjanganQA{0}
	e1 := newEmployee("Anton", "Pratama", "Programmer", 0, 0)
	e2 := newEmployee("Budi", "Junaedi", "Programmer", 0, 0)
	e3 := newEmployee("Charlie", "Van Dick", "Programmer", 0, 0)
	e4 := newEmployee("Dian", "Permana", "Sales", 500_000, 200_000)
	e5 := newEmployee("Fatur", "Rohman", "Sales", 350_000, 250_000)
	e6 := newEmployee("Ellise", "Toon", "QA", 0, 0)
	e7 := newEmployee("Gita", "Gutawa", "QA", 0, 0)

	//print employee menggunakan run and debug
	employees := []mainStruct{*e1, *e2, *e3, *e4, *e5, *e6, *e7}
	for _, value := range employees {

		fmt.Println(value)
	}
	fmt.Println()

	//methode SumBasicSalary
	calculatorBasic := SalaryCalculator{employees: employees}
	totalBasicSalary := calculatorBasic.SumBasicSalary()
	fmt.Println("Total Basic Salary: ", totalBasicSalary)

	//metode sumTotalSalary
	calculatorTotal := SalaryCalculator{employees: employees}
	totalSalary := calculatorTotal.SumTotalSalary()
	fmt.Println("Total Salary All: ", totalSalary)

	//metode sum total salary berdasarkan role
	calTotalByRole := SalaryCalculator{employees: employees}

	totalProgrammerSalary := calTotalByRole.TotalSalaryByRole("Programmer")
	fmt.Println("Total Programmer Salary:", totalProgrammerSalary)

	totalSalesSalary := calTotalByRole.TotalSalaryByRole("Sales")
	fmt.Println("Total Sales Salary:", totalSalesSalary)

	totalQASalary := calTotalByRole.TotalSalaryByRole("QA")
	fmt.Println("Total QA Salary:", totalQASalary)

	//metode sum total salary berdasarkan role
	calEmployeeByRole := SalaryCalculator{employees: employees}

	totalEmployeeProgrammer := calEmployeeByRole.TotalEmployeeByRole("Programmer")
	fmt.Println("Total Programmer Salary:", totalEmployeeProgrammer)

	totalEmployeeSales := calEmployeeByRole.TotalEmployeeByRole("Sales")
	fmt.Println("Total Sales Salary:", totalEmployeeSales)

	totalEmployeeQA := calEmployeeByRole.TotalEmployeeByRole("QA")
	fmt.Println("Total QA Salary:", totalEmployeeQA)

	//total employee
	totalEmployee := totalEmployeeProgrammer + totalEmployeeSales + totalEmployeeQA
	fmt.Println("Total Epmloyee", totalEmployee)

}
