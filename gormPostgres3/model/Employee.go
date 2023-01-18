package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EmployeeName string       `json:"employeeName"`
	EmpSalary    uint         `json:"empSalary"`
	Departments  []Department `json:"departments"gorm:"foreignKey:EmployeeId"`
}
