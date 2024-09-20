package models

type Employee struct {
	empId string,
	name string,
	email string,
	address Address,
	contact string
	
}

type Address {
	street string,
	location string,
	city string,
	state string,
	country string,
	zipCode string
}