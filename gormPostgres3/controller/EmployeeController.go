package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"onetomany/databaseconnection"
	"onetomany/model"
)

func CreateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("creating employee")
	db := databaseconnection.Connection()
	if db == nil {
		log.Println("database connection failed")
	}
	var Employee model.Employee
	var departments []model.Department = Employee.Departments
	json.NewDecoder(request.Body).Decode(&Employee)
	emp := db.Create(&Employee)
	if emp != nil {
		log.Println("one employee created ")
	}
	emp1 := db.Create(&departments)
	if emp1 != nil {
		log.Println("department created")
	}
	json.NewEncoder(response).Encode(Employee.ID)
	return

}
func GetOneEmployee(response http.ResponseWriter, request *http.Request) {
	log.Println("getting one employee")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.Connection()
	params := mux.Vars(request)
	var Employee model.Employee
	db.Model(model.Employee{}).Preload("Departments").First(&Employee, params["employeeId"])
	json.NewEncoder(response).Encode(Employee)
	return

}
func GetAllEmployees(response http.ResponseWriter, request *http.Request) {
	log.Println("getting all employees")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.Connection()
	var employees []model.Employee
	db.Model(model.Employee{}).Preload("Departments").Find(&employees)
	json.NewEncoder(response).Encode(&employees)
	return

}
func UpdateEmployee(response http.ResponseWriter, request *http.Request) {
	log.Println("updating employee")
	response.Header().Set("content-type", "application/json")
	db := databaseconnection.Connection()
	if db != nil {
		log.Println("database connection successfully")
	}
	params := mux.Vars(request)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Departments").First(&employee, params["employeeId"])
	json.NewDecoder(request.Body).Decode(&employee)
	var departments []model.Department = employee.Departments
	db.Save(&employee)
	db.Save(departments)
	json.NewEncoder(response).Encode(employee.ID)
	return

}
func DeleteEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	log.Println("deleting one employee")
	db := databaseconnection.Connection()
	if db != nil {
		log.Println("connection successful")
	}
	params := mux.Vars(request)
	log.Println(params)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Departments").First(&employee, params["employeeId"])
	var department = employee.Departments
	db.Delete(&department)
	db.Delete(&employee)
	json.NewEncoder(response).Encode("data of employee and department successfully deleted")
	return

}
