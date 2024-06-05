package employee

import (
	"context"

	"github.com/google/uuid"
	"github.com/services/employee/pkg/database"
	"github.com/services/employee/pkg/entity/model"
)

var (
	createEmployee   = database.Insert
	getAllEmployee   = database.ScanAll
	getPaginatedData = database.GetPaginatedData
	getEmployee      = database.ScanOne
	updateEmployee   = database.Update
	deleteEmployee   = database.Delete
)

func CreateEmployeeData(ctx context.Context, employee model.Detail) error {
	return createEmployee(&employee)
}

func GetAllEmployeeData(ctx context.Context, size int, offset int) ([]model.Detail, error) {
	employees := []model.Detail{}
	var err error
	if size > 0 && offset > 0 {
		err = getPaginatedData(size, offset, &employees)
	} else {
		err = getAllEmployee(&employees)
	}
	if err != nil {
		return nil, err
	}
	return employees, err
}

func GetEmployeeDataById(ctx context.Context, employeeId uuid.UUID) (model.Detail, error) {
	employee := model.Detail{}
	err := getEmployee(&employee, "id = ?", employeeId)
	if err != nil {
		return model.Detail{}, err
	}
	return employee, err
}

func UpdateEmployeeData(ctx context.Context, employee model.Detail) (model.Detail, error) {

	err := updateEmployee(employee)
	if err != nil {
		return model.Detail{}, err
	}
	return employee, err
}

func DeleteEmployeeData(ctx context.Context, employeeId uuid.UUID) error {
	return deleteEmployee(model.Detail{Id: employeeId})
}
