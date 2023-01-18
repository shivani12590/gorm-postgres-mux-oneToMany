package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"onetomany/controller"
)

func Handler() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", controller.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{employeeId}", controller.GetOneEmployee).Methods("GET")
	r.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{employeeId}", controller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{employeeId}", controller.DeleteEmployee).Methods("DELETE")
	http.ListenAndServe(":2020", r)

}
