package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound  = gorm.ErrRecordNotFound
	DbConnection GormDbAPI
	SqlDb        SqlDbAPI
)

type (
	GormDbAPI interface {
		DB() (*sql.DB, error)
		Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
		Where(query interface{}, args ...interface{}) (tx *gorm.DB)
		Create(value interface{}) (tx *gorm.DB)
		Updates(values interface{}) (tx *gorm.DB)
		Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
		Save(value interface{}) (tx *gorm.DB)
		Model(value interface{}) (tx *gorm.DB)
		Limit(limit int) (tx *gorm.DB)
		Offset(offset int) (tx *gorm.DB)
	}

	SqlDbAPI interface {
		Close() error
	}
)

func Load() (err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"postgres",
		"applicationuser",
		"applicationpassword",
		"employeesvc",
		"5432",
		"disable",
	)

	DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		PrepareStmt:            true,
		NowFunc: func() time.Time {
			utc, _ := time.LoadLocation("")
			return time.Now().In(utc)
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "employee" + ".",
			SingularTable: false,
		},
	})
	return
}

func CloseDatabase() error {
	return SqlDb.Close()
}

func ScanAll(target interface{}) error {
	return DbConnection.Find(target).Error
}

func GetPaginatedData(pageSize int, pageNumber int, target interface{}) error {
	return DbConnection.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(target).Error
}

func Scan(target interface{}, query string, queryArgs ...interface{}) error {
	return DbConnection.Where(query, queryArgs...).Find(target).Error
}

func ScanOne(target interface{}, query string, queryArgs ...interface{}) error {
	res := DbConnection.Where(query, queryArgs...).First(target)
	return HandleOneError(res)
}

func Insert(target interface{}) (err error) {
	return DbConnection.Create(target).Error
}

func Update(target interface{}) (err error) {
	return DbConnection.Save(target).Error
}

func UpdateWhere(target interface{}, query string, queryArgs ...interface{}) (err error) {
	res := DbConnection.Where(query, queryArgs...).Updates(target)
	return HandleOneError(res)
}

func Delete(target interface{}) (err error) {
	return DbConnection.Delete(target).Error
}

func DeleteWhere(target interface{}, query string, queryArgs ...interface{}) (err error) {
	res := DbConnection.Where(query, queryArgs...).Delete(target)
	return HandleOneError(res)
}

func HandleOneError(res *gorm.DB) error {
	if err := HandleError(res); err != nil {
		return err
	}

	if res.RowsAffected != 1 {
		return ErrNotFound
	}

	return nil
}

func HandleError(res *gorm.DB) error {
	if res.Error != nil && !errors.Is(res.Error, ErrNotFound) {
		return fmt.Errorf("error: %w", res.Error)
	}

	return nil
}
