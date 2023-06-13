package main

import "fmt"

//main struct
type WhiteBoard struct {
	empId                     int
	firstName, lastName, role string
	basicSalary               int
	tunjangan                 *Tunjangan
	totalsalary               int
}
type Tunjangan struct {
	bonus, commision, lembur, makan int
}

//constructor
func whiteOfBoard(empId int, firstName, lastName, role string, basicSalary int, tunjangan *Tunjangan, totalsalary int) *WhiteBoard {

	return &WhiteBoard{empId, firstName, lastName, role, basicSalary, tunjangan, totalsalary}

}

var seqEmpId int = 119

func (board *WhiteBoard) calcTotal() int {
	//untuk memberikan employee id secara automatis
	seqEmpId++
	board.empId = seqEmpId

	if board.role == "Programmer" {
		board.basicSalary = 6_000_000
		board.totalsalary = board.basicSalary + board.tunjangan.lembur
	} else if board.role == "Sales" {
		board.basicSalary = 3_000_000
		board.totalsalary = board.basicSalary + board.tunjangan.commision + board.tunjangan.bonus
	} else {
		board.basicSalary = 4_500_000
		board.totalsalary = board.basicSalary + board.tunjangan.makan
	}
	return board.totalsalary
}

func main() {
	//Memberikan data tunjangan baik programmer, sales, maupun qa secara manual
	tunjanganProgrammer := &Tunjangan{0, 0, 500_000, 0}
	tunjanganSales1 := &Tunjangan{500_000, 200_000, 0, 0}
	tunjanganSales2 := &Tunjangan{350_000, 250_000, 0, 0}
	tunjanganQA := &Tunjangan{0, 0, 0, 10_000}

	//init constructor
	p1 := whiteOfBoard(0, "Anton", "Pratama", "Programmer", 0, tunjanganProgrammer, 0)
	p2 := whiteOfBoard(0, "Budi", "Junaedi", "Programmer", 0, tunjanganProgrammer, 0)
	p3 := whiteOfBoard(0, "Charlie", "Van Dijk", "Programmer", 0, tunjanganProgrammer, 0)
	s1 := whiteOfBoard(0, "Dian", "Permana", "Sales", 0, tunjanganSales1, 0)
	s2 := whiteOfBoard(0, "Fatur", "Rohman", "Sales", 0, tunjanganSales2, 0)
	q1 := whiteOfBoard(0, "Ellise", "Toon", "QA", 0, tunjanganQA, 0)
	q2 := whiteOfBoard(0, "Gita", "Gutawa", "QA", 0, tunjanganQA, 0)

	p1.totalsalary = p1.calcTotal()
	p2.totalsalary = p2.calcTotal()
	p3.totalsalary = p3.calcTotal()
	s1.totalsalary = s1.calcTotal()
	s2.totalsalary = s2.calcTotal()
	q1.totalsalary = q1.calcTotal()
	q2.totalsalary = q2.calcTotal()

	employees := []WhiteBoard{*p1, *p2, *p3, *s1, *s2, *q1, *q2}

	//untuk menampilkan output Rapi

	//untuk mencari nilai total employee
	count := 0
	//untuk mencari nilai total basic salary
	basicSalary := 0
	//untuk mencari nilai total salary
	TotalSalary := 0

	//untuk mencari total employee
	employeeProgrammer := 0
	employeeSales := 0
	employeeQa := 0

	//untuk mencari total salary by role
	TotalSalaryProgrammer := 0
	TotalSalarySales := 0
	TotalSalaryQa := 0

	for _, value := range employees {
		//total employee
		count++
		//total basic salary
		basicSalary += value.basicSalary
		TotalSalary += value.totalsalary

		if value.role == "Programmer" {
			employeeProgrammer++
			TotalSalaryProgrammer += value.totalsalary
			fmt.Printf("Employee Id : %d\nEmployee Name : %s %s\nEmployee Role : %s\nEmployee Basic Salary : %d\nEmployee Tunjangan : %d\nEmployee Total Salary : %d", value.empId, value.firstName, value.lastName, value.role, value.basicSalary, value.tunjangan.lembur, value.totalsalary)
			println()
			println()
		} else if value.role == "Sales" {
			employeeSales++
			TotalSalarySales += value.totalsalary
			totalTunjangan := value.tunjangan.bonus + value.tunjangan.commision
			fmt.Printf("Employee Id : %d\nEmployee Name : %s %s\nEmployee Role : %s\nEmployee Basic Salary : %d\nEmployee Tunjangan : %d\nEmployee Total Salary : %d", value.empId, value.firstName, value.lastName, value.role, value.basicSalary, totalTunjangan, value.totalsalary)
			println()
			println()
		} else {
			employeeQa++
			TotalSalaryQa += value.totalsalary
			fmt.Printf("Employee Id : %d\nEmployee Name : %s %s\nEmployee Role : %s\nEmployee Basic Salary : %d\nEmployee Tunjangan : %d\nEmployee Total Salary : %d", value.empId, value.firstName, value.lastName, value.role, value.basicSalary, value.tunjangan.makan, value.totalsalary)
			println()
			println()
		}
	}
	println("----------Summary-----------")
	fmt.Printf("Total Employee : %d", count)
	println()
	fmt.Printf("Total Basic Salary : %d", basicSalary)
	println()
	fmt.Printf("Total Salary All : %d\n", TotalSalary)
	println()
	println("----------Total Employee By Role-----------")
	fmt.Printf("Total Programmer : %d", employeeProgrammer)
	println()
	fmt.Printf("Total Sales : %d", employeeSales)
	println()
	fmt.Printf("Total QA : %d\n", employeeQa)
	println()
	println("------Total Employee Salary By Role------")
	fmt.Printf("Programmer : %d", TotalSalaryProgrammer)
	println()
	fmt.Printf("Sales : %d", TotalSalarySales)
	println()
	fmt.Printf("QA : %d", TotalSalaryQa)

}
