package service

import (
	"employee-manager/backend/models"
)

type DbService struct {
}

func (db DbService) saveEmployee(employee models.Employee) models.Employee {

	return models.Employee{}
}

func (db DbService) getEmployee(empId string) models.Employee {
	return models.Employee{}
}

func (db DbService) getEmployees() models.Employee {
	return models.Employee{}
}

func (db DbService) deleteEmployee(employee models.Employee) models.Employee {
	return models.Employee{}
}

func (db DbService) updateEmployee(employee models.Employee) models.Employee {
	return models.Employee{}
}

func (db DbService) updateEmployeeAddress(address models.Address, empid string) models.Employee {
	return models.Employee{}
}
