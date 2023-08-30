package employeegorm

import (
	"context"
	"github.com/evgeny-tokarev/office_app/backend/internal/repositories/employeerepository"
	"gorm.io/gorm"
)

type EmployeeGorm struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EmployeeGorm {
	return &EmployeeGorm{
		db: db,
	}
}

func (r *EmployeeGorm) Create(ctx context.Context, e *employeerepository.Employee) error {
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *EmployeeGorm) Get(ctx context.Context, id int64) (*employeerepository.Employee, error) {
	e := new(employeerepository.Employee)
	err := r.db.WithContext(ctx).First(e, id).Error
	return e, err
}

func (r *EmployeeGorm) List(ctx context.Context, officeID int64) ([]*employeerepository.Employee, error) {
	var list []*employeerepository.Employee
	err := r.db.WithContext(ctx).Find(&list).Where("office_id = ?", officeID).Error
	return list, err
}

func (r *EmployeeGorm) Update(ctx context.Context, e *employeerepository.Employee) error {
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *EmployeeGorm) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&employeerepository.Employee{ID: id}).Error
}
