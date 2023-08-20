package employeesqlx

import (
	"context"
	"github.com/jmoiron/sqlx"
	"lesson4/internal/repositories/employeerepository"
)

type EmployeeSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *EmployeeSqlx {
	return &EmployeeSqlx{
		db: db,
	}
}

func (r *EmployeeSqlx) Create(ctx context.Context, e *employeerepository.Employee) error {
	const q = `
		insert into employees (name, age, office_id) 
			values (:name, :age, :office_id)
	`
	_, err := r.db.NamedExecContext(ctx, q, e)
	return err
}

func (r *EmployeeSqlx) Get(ctx context.Context, id int64) (*employeerepository.Employee, error) {
	const q = `
		select id, name, age, office_id from employees where id = $1
	`
	e := new(employeerepository.Employee)
	err := r.db.GetContext(ctx, e, q, id)
	return e, err
}

func (r *EmployeeSqlx) List(ctx context.Context, officeID int64) ([]*employeerepository.Employee, error) {
	const q = `
		select id, name, age, office_id from employees where office_id = $1
	`
	var list []*employeerepository.Employee
	err := r.db.SelectContext(ctx, &list, q, officeID)
	return list, err
}

func (r *EmployeeSqlx) Update(ctx context.Context, e *employeerepository.Employee) error {
	const q = `
		update employees set name=:name, age=:age, office_id=:office_id 
			where id = :id
	`
	_, err := r.db.NamedExecContext(ctx, q, e)
	return err
}

func (r *EmployeeSqlx) Delete(ctx context.Context, id int64) error {
	const q = `
		delete from employees where id = $1
	`
	_, err := r.db.ExecContext(ctx, q, id)
	return err
}
