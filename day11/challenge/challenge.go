package main

import "fmt"

type ListEmployees struct {
	EmpId       int
	FirstName   string
	LastName    string
	Role        string
	BasicSalary int
	Bonus, Commission, Lembur, Makan int
	TotalSalary int
}

type Employee struct{
	employee  []ListEmployees
}

var lastEmpId int = 119

//constructor
func dataEmployee(EmpId int, FirstName, LastName, Role string, BasicSalary int, Bonus, Commission, Lembur, Makan int, TotalSalary int) *ListEmployees {
	return &ListEmployees{EmpId, FirstName, LastName, Role, BasicSalary, Bonus, Commission, Lembur, Makan, TotalSalary}
}

//method masukan data 
func(masuk *ListEmployees) IsiData(){
	lastEmpId++
	masuk.EmpId = lastEmpId
	if masuk.Role == "Programmer" {
		masuk.BasicSalary = 6000000
		masuk.Bonus = 0
		masuk.Commission = 0
		masuk.Lembur = 500000
		masuk.Makan = 0
		masuk.TotalSalary = masuk.BasicSalary+masuk.Lembur
	} else if masuk.Role == "Sales" {
		masuk.BasicSalary = 3000000
		masuk.TotalSalary = masuk.BasicSalary + masuk.Bonus + masuk.Commission
	} else {
		masuk.BasicSalary = 4500000
		masuk.Bonus = 0
		masuk.Commission = 0
		masuk.Lembur = 0
		masuk.Makan = 10000
		masuk.TotalSalary = masuk.BasicSalary + masuk.Makan
	}
}

//method sum basic salary
func (employees *Employee) SumBasicSalary()int{
	totalBasic := 0
	for _, employee := range employees.employee {
		totalBasic += employee.BasicSalary
	}
	return totalBasic
}

//method sum all salary
func (employees *Employee) SumAllSalary()int{
	totalAll := 0
	for _, employee := range employees.employee {
		totalAll += employee.TotalSalary
	}
	return totalAll
}

//Method Total Employee By Role
func (employees *Employee) SumEmployee()int{
	total := 0
	for i := 0; i < len(employees.employee); i++ {
		total++
	}
	return total
}

// method total Programmer
func (employees *Employee) SumProg()int{
	programmer := 0
	for _, employee := range employees.employee {
		if employee.Role == "Programmer"{
			programmer ++
		}
	}
	return programmer
}

// method total Sales
func (employees *Employee) SumSales()int{
	sales := 0
	for _, employee := range employees.employee {
		if employee.Role == "Sales"{
			sales ++
		}
	}
	return sales
}

//method total  QA
func (employees *Employee) SumQa()int{
	qa := 0
	for _, employee := range employees.employee {
		if employee.Role == "QA"{
			qa ++
		}
	}
	return qa
}

// method total salary programmer 
func (employees *Employee) SumSalaryProg()int{
	totalsalary:= 0
	for _, employee := range employees.employee {
		if employee.Role == "Programmer"{
			totalsalary += employee.TotalSalary
		}
	}
	return totalsalary
}

func (employees *Employee) SumSalarySales()int{
	sales := 0
	for _, employee := range employees.employee {
		if employee.Role == "Sales"{
			sales += employee.TotalSalary
		}
	}
	return sales
}

func (employees *Employee) SumSalaryQa()int{
	qa := 0
	for _, employee := range employees.employee {
		if employee.Role == "QA"{
			qa += employee.TotalSalary
		}
	}
	return qa
}

func main() {
	p1 := dataEmployee(0,"Anton", "Pratama", "Programmer", 0, 0, 0, 0, 0, 0)
	p2 := dataEmployee(0, "Budi", "Junaedi", "Programmer", 0, 0, 0, 0, 0, 0)
	p3 := dataEmployee(0, "Charlie", "Van Dijk", "Programmer", 0, 0, 0, 0, 0, 0)
	s1 := dataEmployee(0, "Dian", "Permana", "Sales",0, 500000, 200000, 0, 0, 0)
	s2 := dataEmployee(0, "Fatur", "Rohman", "Sales",0, 350000, 250000, 0, 0, 0)
	q1 := dataEmployee(0, "Ellise", "Toon", "QA", 0, 0, 0, 0, 0, 0)
	q2 := dataEmployee(0, "Gita", "Gutawa", "QA", 0, 0, 0, 0, 0, 0)
	p1.IsiData()
	p2.IsiData()
	p3.IsiData()
	s1.IsiData()
	s2.IsiData()
	q1.IsiData()
	q2.IsiData()
	

	// Print Data employee
	employee := []ListEmployees{*p1, *p2, *p3, *s1, *s2, *q1, *q2}
	for _, v := range employee {
		fmt.Println(v)
	}
	
	// Print Total Basic Salary
	jumlah := Employee{employee: employee}
	totalBasicSalary := jumlah.SumBasicSalary()
	fmt.Println("Total Basic Salary	: ", totalBasicSalary)
	fmt.Println()

	//Print Total All Salary
	totalAllSalary := jumlah.SumAllSalary()
	fmt.Println("Total All Salarys	: ", totalAllSalary)
	fmt.Println()

	//Print total employee
	totalEmployee := jumlah.SumEmployee()
	fmt.Println("Total Employee		: ", totalEmployee)
	fmt.Println()

	// Print total employee by role
	totalProgrammer := jumlah.SumProg()
	totalSales := jumlah.SumSales()
	totalQa := jumlah.SumQa()
	fmt.Println("Total Programmer	: ", totalProgrammer)
	fmt.Println("Total Sales		: ", totalSales)
	fmt.Println("Total QA		: ", totalQa)
	fmt.Println()

	// Print Total Salary By Role
	totalSalaryProg := jumlah.SumSalaryProg()
	totalSalarySales := jumlah.SumSalarySales()
	totalSalaryQa := jumlah.SumSalaryQa()
	fmt.Println("Total Salary Programmer	: ", totalSalaryProg)
	fmt.Println("Total Salary Sales	: ", totalSalarySales)
	fmt.Println("Total Salary QA		: ", totalSalaryQa)
	fmt.Println()

}