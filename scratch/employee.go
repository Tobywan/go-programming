package scratch

// Employee is an employee
type Employee struct {
	ID        int
	Name      string
	ManagerID int
}

func main() {
	e := EmployeeByID2(0)
	e.ManagerID = 10

	EmployeeByID(0).Name = "Wendy"
}

//EmployeeByID is a method
func EmployeeByID(id int) *Employee {
	return &Employee{1, "Toby", 0}

}

//EmployeeByID2 is a method
func EmployeeByID2(id int) Employee {
	return Employee{1, "Toby", 0}

}
