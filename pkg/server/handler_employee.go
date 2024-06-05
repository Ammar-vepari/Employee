package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/services/employee/pkg/entity/employee"
	"github.com/services/employee/pkg/entity/model"
	errors "github.com/services/employee/pkg/error"
)

type (
	EmployeeService interface {
		CreateEmployee(context.Context, model.Detail) error
		GetAllEmployee(ctx context.Context, pageSize string, pageNumber string) ([]model.Detail, error)
		GetEmployeeById(ctx context.Context, employeeId uuid.UUID) (model.Detail, error)
		UpdateEmployee(ctx context.Context, employee model.Detail) (model.Detail, error)
		DeleteEmployee(ctx context.Context, employeeId uuid.UUID) error
	}

	EmployeeHandler struct {
		service EmployeeService
	}
)

func NewEmployeeHandler(employeeService *employee.Service) *EmployeeHandler {
	return &EmployeeHandler{
		service: employeeService,
	}
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("inisde create handler")

	employeeDetails := model.Detail{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&employeeDetails)
	if err != nil {
		fmt.Println("failed to decode body")
		errors.SendError(w, err)
		return
	}
	id := uuid.New()
	employeeDetails.Id = id
	err = h.service.CreateEmployee(ctx, employeeDetails)
	if err != nil {
		fmt.Println("error occured inside service")
		errors.SendError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employeeDetails)

}

func (h *EmployeeHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("inisde Get handler")

	pageSize := strings.TrimSpace(r.URL.Query().Get("PageSize"))
	pageNumber := strings.TrimSpace(r.URL.Query().Get("PageNumber"))
	fmt.Println("pagesize is: ", pageSize)
	fmt.Println("PageNumber is: ", pageNumber)
	employeeList, err := h.service.GetAllEmployee(ctx, pageSize, pageNumber)
	if err != nil {
		fmt.Println("error occured inside service")
		errors.SendError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeList)

}

func (h *EmployeeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("inisde Get by id handler")

	params := mux.Vars(r)
	employeeId := params["id"]
	employeeUUID, err := uuid.Parse(employeeId)
	if err != nil {
		fmt.Println("error occured while parsing uuid")
		errors.SendError(w, err)
		return
	}
	employeeData, err := h.service.GetEmployeeById(ctx, employeeUUID)
	if err != nil {
		fmt.Println("error occured inside service")
		errors.SendError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employeeData)

}

func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("inisde update handler")

	params := mux.Vars(r)
	employeeId := params["id"]
	employeeUUID, err := uuid.Parse(employeeId)
	if err != nil {
		fmt.Println("error occured while parsing uuid")
		errors.SendError(w, err)
		return
	}
	employeeDetails := model.Detail{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&employeeDetails)
	if err != nil {
		fmt.Println("failed to decode body")
		errors.SendError(w, err)
		return
	}
	employeeDetails.Id = employeeUUID
	employeeData, err := h.service.UpdateEmployee(ctx, employeeDetails)
	if err != nil {
		fmt.Println("error occured inside service")
		errors.SendError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employeeData)

}

func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("inisde delete handler")

	params := mux.Vars(r)
	employeeId := params["id"]
	employeeUUID, err := uuid.Parse(employeeId)
	if err != nil {
		fmt.Println("error occured while parsing uuid")
		errors.SendError(w, err)
		return
	}
	err = h.service.DeleteEmployee(ctx, employeeUUID)
	if err != nil {
		fmt.Println("error occured inside service")
		errors.SendError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
