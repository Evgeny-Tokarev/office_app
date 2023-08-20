package employeerepository

import "context"

type Employee struct {
	ID       int64  `db:"id"        gorm:"id;primaryKey"`
	Name     string `db:"name"      gorm:"name"`
	Age      int    `db:"age"       gorm:"age"`
	OfficeID int64  `db:"office_id" gorm:"office_id"`
}

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) error
	Get(ctx context.Context, id int64) (*Employee, error)
	List(ctx context.Context, officeID int64) ([]*Employee, error)
	Update(ctx context.Context, e *Employee) error
	Delete(ctx context.Context, id int64) error
}
