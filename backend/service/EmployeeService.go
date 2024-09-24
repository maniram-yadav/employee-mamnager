package service

import (
	"employee-manager/backend/models"
)

type EmployeeService interface {
	saveEmployee(employee models.Employee) models.Employee
	getEmployee(empId string)
	getEmployees()
	deleteEmployee(employee models.Employee) models.Employee
	updateEmployee(employee models.Employee) models.Employee
	updateEmployeeAddress(address models.Address, empid string) models.Employee
}
