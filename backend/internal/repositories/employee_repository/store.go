package employee_repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbError)
		}
		return err
	}
	return tx.Commit()
}

type EmployeeTransferTxParams struct {
	ID           int64 `json:"id"`
	FromOfficeId int64 `json:"from_office_id"`
	ToOfficeId   int64 `json:"to_office_id"`
}

type EmployeeTransferTxResult struct {
	OriginalEmployeeId int64 `json:"original_employee_id"`
	ResultEmployeeId   int64 `json:"result_employee_id"`
	FromOfficeId       int64 `json:"from_office_id"`
	ToOfficeId         int64 `json:"to_office_id"`
}

func (store *Store) TransferEmployeeTx(ctx context.Context, arg EmployeeTransferTxParams) (EmployeeTransferTxResult, error) {
	var resultEmployee Employee
	err := store.execTx(ctx, func(q *Queries) error {
		var employee Employee
		var err error
		employee, err = q.GetEmployee(ctx, arg.ID)
		if err != nil {
			return err
		}
		resultEmployee, err = q.CreateEmployee(ctx, CreateEmployeeParams{
			Name:     employee.Name,
			Age:      employee.Age,
			OfficeID: arg.ToOfficeId,
		})
		err = q.DeleteEmployee(ctx, arg.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return EmployeeTransferTxResult{
		OriginalEmployeeId: arg.ID,
		ResultEmployeeId:   resultEmployee.ID,
		FromOfficeId:       arg.FromOfficeId,
		ToOfficeId:         arg.ToOfficeId,
	}, err
}
