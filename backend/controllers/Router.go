package controllers

import (
	"github.com/gorilla/mux"
)

// var controller EmployeeController

func Init(controller EmployeeController) *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/employee", controller.saveEmployee).Methods("POST")
	route.HandleFunc("/employee", controller.getEmployees)
	route.HandleFunc("/employee/{empid}", controller.deleteEmployee).Methods("DELETE")
	route.HandleFunc("/employee/{empid}", controller.updateEmployee).Methods("PUT")
	route.HandleFunc("/employee/address/{empid}", controller.updateEmployeeAddress).Methods("PUT")
	route.HandleFunc("/employee/{empid}", controller.getEmployee)
	return route
}
