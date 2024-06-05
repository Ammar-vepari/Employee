package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/services/employee/pkg/entity/employee"
)

func loadRoutes() *mux.Router {
	serviceRouter := mux.NewRouter().PathPrefix("/employee-service").Subrouter()
	apiRouter := serviceRouter.PathPrefix("/api").Subrouter()

	apiV1Router := apiRouter.PathPrefix("/v1").Subrouter()

	setEmployeeDetailsAPIRoutes(apiV1Router)

	return serviceRouter
}

func setEmployeeDetailsAPIRoutes(router *mux.Router) {
	handler := NewEmployeeHandler(&employee.Service{})
	router.HandleFunc("/employees/create", handler.Create).Methods(http.MethodPost)
	router.HandleFunc("/employees", handler.Get).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", handler.GetById).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", handler.Update).Methods(http.MethodPut)
	router.HandleFunc("/employees/{id}", handler.Delete).Methods(http.MethodDelete)
}

// BuildPath get elements and return path
func BuildPath(elements ...string) string {
	return "/" + strings.Join(elements, "/")
}
