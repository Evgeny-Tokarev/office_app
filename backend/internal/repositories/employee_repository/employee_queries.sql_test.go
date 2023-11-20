package employee_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/office_repository"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var testDb *sql.DB
var office office_repository.Office

func TestMain(m *testing.M) {
	var err error
	testDb, err = util.InitTestDB()
	if err != nil {
		fmt.Println("Unable to create testing Database")
		return
	}
	_, office, err = office_repository.CreateRandomOffice(testDb)
	if err != nil {
		fmt.Println("Unable to create testing office", err.Error())
		return
	}
	os.Exit(m.Run())
}

func TestCreateEmployee(t *testing.T) {
	arg, employee, err := createRandomEmployee(testDb, office.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Name, employee.Name)
	require.Equal(t, arg.Age, employee.Age)
	require.NotZero(t, employee.ID)
	require.NotZero(t, employee.CreatedAt)
	require.NotZero(t, employee.UpdatedAt)
}

func TestGetEmployee(t *testing.T) {
	_, employee1, _ := createRandomEmployee(testDb, office.ID)
	employee2, err := TestEmployeeQueries.GetEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, employee2)
	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, employee1.Name, employee2.Name)
	require.WithinDuration(t, employee1.CreatedAt, employee2.CreatedAt, time.Second)
	require.WithinDuration(t, employee1.UpdatedAt, employee2.UpdatedAt, time.Second)
}

func TestDeleteEmployee(t *testing.T) {
	_, employee1, _ := createRandomEmployee(testDb, office.ID)
	err := TestEmployeeQueries.DeleteEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)
	employee2, err := TestEmployeeQueries.GetEmployee(context.Background(), employee1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, employee2)
}

func TestListEmployee(t *testing.T) {
	var employees []Employee
	for i := 0; i < 5; i++ {
		_, currentEmployee, curEr := createRandomEmployee(testDb, office.ID)
		require.NoError(t, curEr)
		employees = append(employees, currentEmployee)
	}
	employee2, err := TestEmployeeQueries.ListEmployees(context.Background(), office.ID)
	require.NoError(t, err)
	require.True(t, len(employee2) >= 5)
	for i, employee1 := range employees {
		employee2 := employee2[len(employee2)-5+i]
		require.NotEmpty(t, employee1)
		require.Equal(t, employee1.ID, employee2.ID)
		require.Equal(t, employee1.Name, employee2.Name)
		require.WithinDuration(t, employee1.CreatedAt, employee2.CreatedAt, time.Second)
		require.WithinDuration(t, employee1.UpdatedAt, employee2.UpdatedAt, time.Second)
	}
}

func TestAttachePhoto(t *testing.T) {
	_, employee1, _ := createRandomEmployee(testDb, office.ID)
	testParams := AttachePhotoParams{
		ImgFile: sql.NullString{String: util.RandomString(10), Valid: true},
		ID:      employee1.ID,
	}

	err := TestEmployeeQueries.AttachePhoto(context.Background(), testParams)
	require.NoError(t, err)
	employee2, err := TestEmployeeQueries.GetEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)
	require.Equal(t, testParams.ImgFile, employee2.ImgFile)
}

func TestTransferEmployee(t *testing.T) {
	n := 5
	var res EmployeeTransferTxResult
	var employee Employee
	fromEmployees := make(map[int64]Employee)
	errs := make(chan error)
	results := make(chan EmployeeTransferTxResult)
	_, office2, err := office_repository.CreateRandomOffice(testDb)
	testStore := NewStore(testDb)
	for i := 0; i < n; i++ {
		go func() {
			_, employee, _ = createRandomEmployee(testDb, office.ID)
			arg := EmployeeTransferTxParams{
				ID:           employee.ID,
				FromOfficeId: office.ID,
				ToOfficeId:   office2.ID,
			}
			fromEmployees[employee.ID] = employee
			res, err = testStore.TransferEmployeeTx(context.Background(), arg)
			errs <- err
			results <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		res := <-results
		employee2, err2 := TestEmployeeQueries.GetEmployee(context.Background(), res.ResultEmployeeId)
		require.NoError(t, err2)
		require.NotEmpty(t, employee2)
		fmt.Println()
		require.Equal(t, fromEmployees[res.OriginalEmployeeId].Name, employee2.Name)
		require.Equal(t, fromEmployees[res.OriginalEmployeeId].Age, employee2.Age)
		require.WithinDuration(t, fromEmployees[res.OriginalEmployeeId].CreatedAt, employee2.CreatedAt, time.Second)
		require.WithinDuration(t, fromEmployees[res.OriginalEmployeeId].UpdatedAt, employee2.UpdatedAt, time.Second)
	}

}
