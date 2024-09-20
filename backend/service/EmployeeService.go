package service

type EmployeeService interface {
	saveEmployee(employee Employee)
	getEmployee(empId string)
	getEmployees()
	deleteEmployee(employee Employee)
	updateEmployee(employee Employee)
	updateEmployeeAddress(address Address, empid string)
}
