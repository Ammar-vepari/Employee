package employee

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/services/employee/pkg/entity/model"
	errors "github.com/services/employee/pkg/error"
)

type (
	Service struct {
	}
)

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateEmployee(ctx context.Context, employee model.Detail) error {
	fmt.Println("inside create employee function")
	err := CreateEmployeeData(ctx, employee)
	if err != nil {
		return errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "error occured while creating data"})
	}
	return nil
}

func (s *Service) GetAllEmployee(ctx context.Context, pageSize string, pageNumber string) ([]model.Detail, error) {
	fmt.Println("inside get all employee function")

	if pageSize == "" {
		pageSize = "0"
	}
	size, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "unexpected error"})
	}

	if pageNumber == "" {
		pageNumber = "1"
	}
	offset, err := strconv.Atoi(pageNumber)
	if err != nil {
		return nil, errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "unexpected error"})
	}

	employeeList, err := GetAllEmployeeData(ctx, size, offset)
	if err != nil {
		return nil, errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "error occured while getting data"})
	}
	return employeeList, nil
}

func (s *Service) GetEmployeeById(ctx context.Context, employeeId uuid.UUID) (model.Detail, error) {
	fmt.Println("inside get by employee id function")

	employeeList, err := GetEmployeeDataById(ctx, employeeId)
	if err != nil {
		return model.Detail{}, errors.NewApplicationError(http.StatusBadRequest, errors.ErrorResponse{ErrorType: "EmployeeDetilsNotFound", Message: "Employee details doesnt exist"})
	}
	return employeeList, nil
}

func (s *Service) UpdateEmployee(ctx context.Context, employee model.Detail) (model.Detail, error) {
	fmt.Println("inside update employee function")

	//validate employee exist
	_, err := GetEmployeeDataById(ctx, employee.Id)
	if err != nil {
		return model.Detail{}, errors.NewApplicationError(http.StatusBadRequest, errors.ErrorResponse{ErrorType: "EmployeeDetilsNotFound", Message: "Employee details doesnt exist"})
	}
	employeeList, err := UpdateEmployeeData(ctx, employee)
	if err != nil {

		return model.Detail{}, errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "error occured while updating data"})
	}
	return employeeList, nil
}

func (s *Service) DeleteEmployee(ctx context.Context, employeeId uuid.UUID) error {
	fmt.Println("inside delete employee function")

	//check if employee exist
	_, err := GetEmployeeDataById(ctx, employeeId)
	if err != nil {
		return nil
	}

	err = DeleteEmployeeData(ctx, employeeId)
	if err != nil {
		return errors.NewApplicationError(http.StatusInternalServerError, errors.ErrorResponse{ErrorType: "UnexpectedError", Message: "error occured while deleting"})
	}
	return nil
}
