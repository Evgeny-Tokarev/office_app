package employee_repository

import (
	"context"
	"database/sql"
	"github.com/evgeny-tokarev/office_app/backend/util"
)

var TestEmployeeQueries *Queries

func createRandomEmployee(testDB *sql.DB, id int64) (CreateEmployeeParams, Employee, error) {
	arg := CreateEmployeeParams{
		Name:     util.RandomString(10),
		Age:      int32(util.RandomInt(20, 50)),
		OfficeID: id,
	}
	TestEmployeeQueries = New(testDB)
	employee, err := TestEmployeeQueries.CreateEmployee(context.Background(), arg)
	return arg, employee, err
}
