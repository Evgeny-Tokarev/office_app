package office_repository

import (
	"context"
	"database/sql"
	"github.com/evgeny-tokarev/office_app/backend/util"
)

var TestOfficeQueries *Queries

func CreateRandomOffice(testDb *sql.DB) (CreateOfficeParams, Office, error) {
	arg := CreateOfficeParams{
		Name:    util.RandomString(10),
		Address: util.RandomString(20),
	}
	TestOfficeQueries = New(testDb)
	office, err := TestOfficeQueries.CreateOffice(context.Background(), arg)
	return arg, office, err
}
