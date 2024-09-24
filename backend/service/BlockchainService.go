package service

import (
	"employee-manager/backend/models"
)

type BlockchainService struct {
}

func (bService BlockchainService) saveEmployee(employee models.Employee) models.Employee {
	return models.Employee{}
}

func (bService BlockchainService) getEmployee(empId string) models.Employee {
	return models.Employee{}
}

func (bService BlockchainService) getEmployees() models.Employee {
	return models.Employee{}
}

func (bService BlockchainService) deleteEmployee(employee models.Employee) models.Employee {
	return models.Employee{}
}

func (dbServiceb BlockchainService) updateEmployee(employee models.Employee) models.Employee {
	return models.Employee{}
}

func (bService BlockchainService) updateEmployeeAddress(address models.Address, empid string) models.Employee {
	return models.Employee{}
}
