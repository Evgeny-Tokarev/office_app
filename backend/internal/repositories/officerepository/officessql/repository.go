package officessql

import (
	"context"
	"database/sql"
	"errors"
	"lesson4/internal/repositories/officerepository"
)

type OfficeSql struct {
	db *sql.DB
}

func New(db *sql.DB) *OfficeSql {
	return &OfficeSql{
		db: db,
	}
}

func (r *OfficeSql) Create(ctx context.Context, o *officerepository.Office) error {
	const q = `
		insert into offices (name, address, created_at, updated_at) 
			values ($1, $2, $3, $4)
		returning id
	`

	err := r.db.QueryRowContext(ctx, q, o.Name, o.Address, o.CreatedAt, o.UpdatedAt).Scan(&o.ID)
	return err
}

func (r *OfficeSql) Get(ctx context.Context, id int64) (*officerepository.Office, error) {
	const q = `
		select id, name, address, created_at, updated_at from offices where id = $1
	`
	o := new(officerepository.Office)
	err := r.db.QueryRowContext(ctx, q, id).Scan(&o.ID, &o.Name, &o.Address, &o.CreatedAt, &o.UpdatedAt)
	return o, err
}

func (r *OfficeSql) List(ctx context.Context) ([]*officerepository.Office, error) {
	const q = `
		select id, name, address, created_at, updated_at from offices
	`
	var list []*officerepository.Office
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		o := new(officerepository.Office)
		err := rows.Scan(&o.ID, &o.Name, &o.Address, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *OfficeSql) Update(ctx context.Context, o *officerepository.Office) error {
	const q = `
		update offices set name=$1, address=$2, updated_at=$3 
			where id = $4
	`
	_, err := r.db.ExecContext(ctx, q, o.Name, o.Address, o.UpdatedAt, o.ID)
	return err
}

func (r *OfficeSql) Delete(ctx context.Context, id int64) error {
	const q = `
		delete from offices where id = $1
	`
	var exists bool
	err := r.db.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM offices WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("no such ID found")
	}

	_, err = r.db.ExecContext(ctx, q, id)
	return err
}
