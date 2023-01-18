package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	DepartmentName string `json:"departmentName"`
	EmployeeId     uint   `gorm:"column:employee_id"`
	Employee       Employee
}
