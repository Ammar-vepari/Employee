package server

import (
	"context"
	"fmt"

	"github.com/services/employee/pkg/database"
)

var (
	loadDB = database.Load
)

func LoadDependencies(ctx context.Context) error {
	if err := loadDB(); err != nil {
		return err
	}
	return nil
}

func closeDependencies() {
	if database.SqlDb != nil {
		if err := database.CloseDatabase(); err != nil {
			fmt.Printf("Failed to gracefully close database connection with error as %v", err)
		}
	}

}
