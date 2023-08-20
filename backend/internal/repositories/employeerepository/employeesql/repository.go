package employeesql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"lesson4/internal/repositories/employeerepository"
)

type EmployeeSql struct {
	db *sql.DB
}

func New(db *sql.DB) *EmployeeSql {
	return &EmployeeSql{
		db: db,
	}
}

func (r *EmployeeSql) Create(ctx context.Context, e *employeerepository.Employee) error {
	const q = `
		insert into employees (name, age, office_id) 
			values ($1, $2, $3)
		returning id
	`
	err := r.db.QueryRowContext(ctx, q, e.Name, e.Age, e.OfficeID).Scan(&e.ID)
	return err
}

func (r *EmployeeSql) Get(ctx context.Context, id int64) (*employeerepository.Employee, error) {
	const q = `
		select id, name, age, office_id from employees where id = $1
	`
	e := new(employeerepository.Employee)
	err := r.db.QueryRowContext(ctx, q, id).Scan(&e.ID, &e.Name, &e.Age, &e.OfficeID)
	return e, err
}

func (r *EmployeeSql) List(ctx context.Context, officeID int64) ([]*employeerepository.Employee, error) {
	const q = `
		select id, name, age, office_id from employees where office_id = $1
	`
	var list []*employeerepository.Employee
	rows, err := r.db.QueryContext(ctx, q, officeID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("no data found for the office %d", officeID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		e := new(employeerepository.Employee)
		err := rows.Scan(&e.ID, &e.Name, &e.Age, &e.OfficeID)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}

func (r *EmployeeSql) Update(ctx context.Context, e *employeerepository.Employee) error {
	const q = `
		update employees set name=$1, age=$2, office_id=$3 
			where id = $4
	`
	_, err := r.db.ExecContext(ctx, q, e.Name, e.Age, e.OfficeID, e.ID)
	return err
}

func (r *EmployeeSql) Delete(ctx context.Context, id int64) error {
	const q = `
		delete from employees where id = $1
	`
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM employees WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("no such ID found")
	}

	_, err = r.db.ExecContext(ctx, q, id)
	return err
}
