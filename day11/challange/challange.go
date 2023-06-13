package main

import "fmt"

type employees struct {
	empid                                                      int
	firstName, LastName, role                                  string
	BasicSalary, Bonus, commission, lembur, makan, totalsalary float64
}

func newemp(empid int, firstName, LastName, role string, BasicSalary, Bonus, commission, lembur, makan, totalsalary float64) *employees {
	return &employees{
		empid:       empid,
		firstName:   firstName,
		LastName:    LastName,
		role:        role,
		BasicSalary: BasicSalary,
		Bonus:       Bonus,
		commission:  commission,
		lembur:      lembur,
		makan:       makan,
	}

}

func (employees *employees) calculation() {
	switch employees.role {
	case "programmer":
		employees.totalsalary = employees.BasicSalary + employees.lembur
	case "sales":
		employees.totalsalary = employees.BasicSalary + employees.Bonus + employees.commission
	case "QA":
		employees.totalsalary = employees.BasicSalary + employees.makan
	default:
		employees.totalsalary = employees.BasicSalary
	}
}

func (e *employees) Display() {
	fmt.Printf("empid %d\n", e.empid)
	fmt.Printf("firstName %s\n", e.firstName)
	fmt.Printf("LastName %s\n", e.LastName)
	fmt.Printf("role %s\n", e.role)
	fmt.Printf("BasicSalary %.2f\n", e.BasicSalary)
	fmt.Printf("Bonus %.2f\n", e.Bonus)
	fmt.Printf("commission %.2f\n", e.commission)
	fmt.Printf("lembur %.2f\n", e.lembur)
	fmt.Printf("makan %.2f\n", e.makan)
	fmt.Printf("totalsalary %.2f\n\n", e.totalsalary)
}

func main() {
	emp := []*employees{}
	emp1 := newemp(120, "Anu", "Pebriyanto", "programmer", 6_000_000, 0, 0, 5_000_00, 0, 0)
	emp2 := newemp(121, "Didi", "Andriansyah", "programmer", 6_000_000, 0, 0, 5_000_00, 0, 0)
	emp3 := newemp(122, "Dodo", "Syarif", "programmer", 6_000_000, 0, 0, 5_000_00, 0, 0)
	emp4 := newemp(123, "Ani", "Pebriyanti", "Sales", 3_000_000, 5_00_000, 2_00_000, 0, 0, 0)
	emp5 := newemp(124, "Petet", "Muhamad", "Sales", 3_000_000, 3_50_00_000, 2_50_000, 0, 0, 0)
	emp6 := newemp(125, "Van djik", "Muhamad", "QA", 4_500_000, 0, 0, 0, 10_000, 0)
	emp7 := newemp(125, "Annisa", "Fadhila", "QA", 4_500_000, 0, 0, 0, 10_000, 0)
	emp = append(emp, emp1, emp2, emp3, emp4, emp5, emp6, emp7)
	for _, emp := range emp {
		emp.calculation()
		emp.Display()
	}
	fmt.Printf("Total employees: %d\n", len(emp))
	totalbasicsalary := 0.0
	for _, emp := range emp {
		totalbasicsalary += emp.BasicSalary
	}
	fmt.Printf("total basic salary: %.2f\n", totalbasicsalary)
	totalsalaryall := 0.0
	for _, emp := range emp {
		totalsalaryall += emp.totalsalary
	}
	fmt.Printf("total salary: %.2f\n", totalsalaryall)

	fmt.Printf("Total Employees by Role : \n")
	roles := make(map[string]int)
	for _, emp := range emp {
		roles[emp.role]++
	}
	for role, count := range roles {
		fmt.Printf("%s : %d\n", role, count)
	}

	fmt.Printf("-------------------------------------\n")

	fmt.Printf("Total Employees Salary by Role : \n")
	salaryByRole := make(map[string]float64)
	for _, emp := range emp {
		salaryByRole[emp.role] += emp.totalsalary
	}
	for role, count := range salaryByRole {
		fmt.Printf("%s : %.2f\n", role, count)
	}

}
